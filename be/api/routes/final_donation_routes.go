package routes

import (
	"milestone3/be/api/middleware"
	"milestone3/be/internal/controller"
)

func (r *EchoRouter) RegisterFinalDonationRoutes(finalDonationCtrl *controller.FinalDonationController) {
	finalDonationRoutes := r.echo.Group("/donations/final")
	finalDonationRoutes.Use(middleware.JWTMiddleware)

	finalDonationRoutes.GET("", finalDonationCtrl.GetAllFinalDonations)
	finalDonationRoutes.GET("/me", finalDonationCtrl.GetMyFinalDonations)
	finalDonationRoutes.GET("/user/:user_id", finalDonationCtrl.GetAllFinalDonationsByUserID)
	finalDonationRoutes.POST("/notes", finalDonationCtrl.UpdateNotes)
}
