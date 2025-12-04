package routes

import (
	"milestone3/be/api/middleware"
	"milestone3/be/internal/controller"
)

func (r *EchoRouter) RegisterFinalDonationRoutes(finalDonationCtrl *controller.FinalDonationController) {
	finalDonationRoutes := r.echo.Group("/final_donations")

	auth := finalDonationRoutes.Group("")
	auth.Use(middleware.JWTMiddleware)

	admin := auth.Group("")
	admin.Use(middleware.RequireAdmin)
	admin.GET("", finalDonationCtrl.GetAllFinalDonations)
	admin.GET("/user/:user_id", finalDonationCtrl.GetAllFinalDonationsByUserID)

	auth.GET("/me", finalDonationCtrl.GetMyFinalDonations)
}
