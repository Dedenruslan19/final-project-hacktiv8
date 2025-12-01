package routes

import (
	"milestone3/be/api/middleware"
	"milestone3/be/internal/controller"
)

func (r *EchoRouter) RegisterDonationRoutes(donationCtrl *controller.DonationController) {
	donationRoutes := r.echo.Group("/donations")

	// public
	donationRoutes.GET("", donationCtrl.GetAllDonations)
	donationRoutes.GET("/:id", donationCtrl.GetDonationByID)

	// authenticated actions (owner or admin logic enforced in controller/service)
	donationRoutes.POST("", donationCtrl.CreateDonation, middleware.JWTMiddleware)
	donationRoutes.PUT("/:id", donationCtrl.UpdateDonation, middleware.JWTMiddleware)
	donationRoutes.PATCH("/:id", donationCtrl.PatchDonation, middleware.JWTMiddleware)
	donationRoutes.DELETE("/:id", donationCtrl.DeleteDonation, middleware.JWTMiddleware)
}
