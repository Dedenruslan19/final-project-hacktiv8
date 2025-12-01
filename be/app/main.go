package main

import (
	"context"
	"log"
	"milestone3/be/api/routes"
	"milestone3/be/config"
	"milestone3/be/internal/controller"
	"milestone3/be/internal/repository"
	"milestone3/be/internal/service"
	"os"

	"cloud.google.com/go/storage"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	db := config.ConnectionDb()
	validator := validator.New()
	ctx := context.Background()

	//dependency injection
	// create GCS client
	var gcsRepo repository.GCSStorageRepo
	if bucket := os.Getenv("GCS_BUCKET"); bucket != "" {
		gcsClient, err := storage.NewClient(ctx)
		if err != nil {
			log.Fatalf("failed to create gcs client: %v", err)
		}
		gcsRepo = repository.NewGCSStorageRepo(gcsClient, bucket)
	} else {
		log.Println("GCS_BUCKET not set — file uploads to GCS will fail if used")
	}
	//repository
	userRepo := repository.NewUserRepo(db, ctx)
	articleRepo := repository.NewArticleRepo(db)
	donationRepo := repository.NewDonationRepo(db)
	finalDonationRepo := repository.NewFinalDonationRepository(db)

	//service
	userServ := service.NewUserService(userRepo)
	articleSvc := service.NewArticleService(articleRepo)
	donationSvc := service.NewDonationService(donationRepo)
	finalDonationSvc := service.NewFinalDonationService(finalDonationRepo)

	//controller
	userControl := controller.NewUserController(validator, userServ)
	articleCtrl := controller.NewArticleController(articleSvc)
	// donation controller needs storage repo; pass nil-safe repo if not configured
	var donationCtrl *controller.DonationController
	if gcsRepo != nil {
		donationCtrl = controller.NewDonationController(donationSvc, gcsRepo)
	} else {
		// repository package defines interface; create a no-op implementation if needed,
		// but here we pass nil — ensure controller handles nil storage or set a stub.
		donationCtrl = controller.NewDonationController(donationSvc, nil)
	}
	finalDonationCtrl := controller.NewFinalDonationController(finalDonationSvc)
	//echo
	e := echo.New()
	//router
	router := routes.NewRouter(e)
	router.RegisterUserRoutes(userControl)
	router.RegisterArticleRoutes(articleCtrl)
	router.RegisterDonationRoutes(donationCtrl)
	router.RegisterFinalDonationRoutes(finalDonationCtrl)

	address := os.Getenv("PORT")
	if err := e.Start(":" + address); err != nil {
		log.Printf("faile to start server %s", err)
	}
}
