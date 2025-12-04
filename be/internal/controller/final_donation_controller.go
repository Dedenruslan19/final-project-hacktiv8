package controller

import (
	"milestone3/be/internal/service"
	"milestone3/be/internal/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FinalDonationController struct {
	svc service.FinalDonationService
}

func NewFinalDonationController(finalDonationService service.FinalDonationService) *FinalDonationController {
	return &FinalDonationController{svc: finalDonationService}
}

// GetAllFinalDonations godoc
// @Summary Get all final donations
// @Description Retrieve all items that were directly donated to institutions with pagination (admin only)
// @Tags Your Donate Rise API - Final Donations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10, max: 100)"
// @Success 200 {object} utils.SuccessResponseData "Final donations fetched successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} utils.ErrorResponse "Forbidden - Admin access required"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /donations/final [get]
func (h *FinalDonationController) GetAllFinalDonations(c echo.Context) error {
	if !utils.IsAdmin(c) {
		return utils.ForbiddenResponse(c, "admin only")
	}

	// Parse pagination params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	finalDonations, total, err := h.svc.GetAllFinalDonations(page, limit)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch final donations")
	}

	response := map[string]interface{}{
		"final_donations": finalDonations,
		"page":            page,
		"limit":           limit,
		"total":           total,
	}
	return utils.SuccessResponse(c, "Final donations fetched successfully", response)
}
// GetMyFinalDonations godoc
// @Summary Get my final donations
// @Description Retrieve all final donations made by the authenticated user
// @Tags Your Donate Rise API - Final Donations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.SuccessResponseData "Final donations fetched successfully"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /donations/final/me [get]
func (h *FinalDonationController) GetMyFinalDonations(c echo.Context) error {
	userID, ok := utils.GetUserID(c)
	if !ok || userID == 0 {
		return utils.UnauthorizedResponse(c, "unauthenticated")
	}

	finalDonations, err := h.svc.GetAllFinalDonationsByUserID(int(userID))
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch final donations")
	}
	return utils.SuccessResponse(c, "Final donations fetched successfully", finalDonations)
}

// GetAllFinalDonationsByUserID godoc
// @Summary Get final donations by user ID
// @Description Retrieve all final donations made by a specific user (admin only)
// @Tags Your Donate Rise API - Final Donations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path int true "User ID"
// @Success 200 {object} utils.SuccessResponseData "Final donations fetched successfully"
// @Failure 400 {object} utils.ErrorResponse "Bad request - Invalid user ID"
// @Failure 401 {object} utils.ErrorResponse "Unauthorized - Invalid or missing token"
// @Failure 403 {object} utils.ErrorResponse "Forbidden - Admin access required"
// @Failure 500 {object} utils.ErrorResponse "Internal server error"
// @Router /donations/final/user/{user_id} [get]
func (h *FinalDonationController) GetAllFinalDonationsByUserID(c echo.Context) error {
	if !utils.IsAdmin(c) {
		return utils.ForbiddenResponse(c, "admin only")
	}

	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return utils.BadRequestResponse(c, "Invalid user id")
	}

	finalDonations, err := h.svc.GetAllFinalDonationsByUserID(userID)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "Failed to fetch final donations")
	}
	return utils.SuccessResponse(c, "Final donations fetched successfully", finalDonations)
}
