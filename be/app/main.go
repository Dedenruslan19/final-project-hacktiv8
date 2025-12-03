package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"milestone3/be/api/routes"
	"milestone3/be/config"
	"milestone3/be/internal/controller"
	"milestone3/be/internal/repository"
	"milestone3/be/internal/service"
)

func main() {
	ctx := context.Background()
	db := config.ConnectionDb()
	validate := validator.New()

	// GCP PUBLIC BUCKET
	var gcpPublicRepo repository.GCPStorageRepo
	publicBucket := os.Getenv("GCS_PUBLIC_BUCKET")

	if publicBucket != "" {
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("failed to create public gcs client: %v", err)
		}
		gcpPublicRepo = repository.NewGCPStorageRepo(client, publicBucket, true)
	} else {
		log.Println("GCS_PUBLIC_BUCKET NOT SET")
	}

	// GCP PRIVATE BUCKET
	var gcpPrivateRepo repository.GCPStorageRepo
	privateBucket := os.Getenv("GCS_PRIVATE_BUCKET")

	if privateBucket != "" {
		client, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("failed to create private gcs client: %v", err)
		}
		gcpPrivateRepo = repository.NewGCPStorageRepo(client, privateBucket, false)
	} else {
		log.Println("GCS_PRIVATE_BUCKET NOT SET")
	}

	// repositories
	userRepo := repository.NewUserRepo(db, ctx)
	articleRepo := repository.NewArticleRepo(db)
	donationRepo := repository.NewDonationRepo(db)
	finalDonationRepo := repository.NewFinalDonationRepository(db)
	paymentRepo := repository.NewPaymentRepository(db, ctx)

	// services
	userSvc := service.NewUserService(userRepo)
	articleSvc := service.NewArticleService(articleRepo)
	donationSvc := service.NewDonationService(donationRepo, gcpPrivateRepo)
	finalDonationSvc := service.NewFinalDonationService(finalDonationRepo)
	paymentSvc := service.NewPaymentService(paymentRepo)

	// controllers
	userCtrl := controller.NewUserController(validate, userSvc)
	articleCtrl := controller.NewArticleController(articleSvc, gcpPublicRepo)

	var donationCtrl *controller.DonationController
	if gcpPrivateRepo != nil {
		donationCtrl = controller.NewDonationController(donationSvc, gcpPrivateRepo)
	} else {
		donationCtrl = controller.NewDonationController(donationSvc, nil)
	}
	finalDonationCtrl := controller.NewFinalDonationController(finalDonationSvc)
	paymentCtrl := controller.NewPaymentController(validate, paymentSvc)

	// echo + router
	e := echo.New()
	router := routes.NewRouter(e)

	router.RegisterUserRoutes(userCtrl)
	router.RegisterArticleRoutes(articleCtrl)
	router.RegisterDonationRoutes(donationCtrl)
	router.RegisterFinalDonationRoutes(finalDonationCtrl)
	router.RegisterPaymentRoutes(paymentCtrl)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
