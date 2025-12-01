package controller

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"milestone3/be/internal/dto"
	"milestone3/be/internal/repository"
	"milestone3/be/internal/service"
	"milestone3/be/internal/utils"

	"github.com/labstack/echo/v4"
)

type DonationController struct {
	svc     service.DonationService
	storage repository.GCSStorageRepo
}

func NewDonationController(s service.DonationService, storage repository.GCSStorageRepo) *DonationController {
	return &DonationController{svc: s, storage: storage}
}

// donor/user: POST /donations
func (h *DonationController) CreateDonation(c echo.Context) error {
	var payload dto.DonationDTO
	contentType := c.Request().Header.Get("Content-Type")

	// support multipart/form-data (with files) OR JSON body
	if strings.HasPrefix(contentType, "multipart/form-data") {
		// parse form values
		if err := c.Request().ParseMultipartForm(32 << 20); err != nil {
			return utils.BadRequestResponse(c, "invalid multipart form")
		}
		form := c.Request().MultipartForm

		// simple mapping: use same form field names as DTO json tags
		payload.Title = form.Value["title"][0]
		payload.Description = form.Value["description"][0]
		payload.Category = form.Value["category"][0]
		payload.Condition = form.Value["condition"][0]
		if v, ok := form.Value["status"]; ok && len(v) > 0 {
			payload.Status = v[0]
		}

		// handle files under form field name "photos"
		if form.File != nil {
			if fhs, ok := form.File["photos"]; ok {
				for _, fh := range fhs {
					f, err := fh.Open()
					if err != nil {
						// skip file on error or return bad request
						_ = f
						return utils.BadRequestResponse(c, "failed open uploaded file")
					}
					// create object name (unique)
					objName := fmt.Sprintf("donations/%d_%s", time.Now().UnixNano(), fh.Filename)
					// upload
					url, err := h.storage.UploadFile(c.Request().Context(), f, objName)
					// close file immediately after upload to avoid FD leaks
					_ = f.Close()
					if err != nil {
						// return internal error on upload failure
						return utils.InternalServerErrorResponse(c, "failed upload file")
					}
					payload.Photos = append(payload.Photos, url)
				}
			}
		}
	} else {
		// JSON binding path
		if err := c.Bind(&payload); err != nil {
			return utils.BadRequestResponse(c, "invalid payload")
		}
	}

	// set user id from context (auth middleware must set it)
	userID, ok := utils.GetUserID(c)
	if !ok || userID == 0 {
		return utils.UnauthorizedResponse(c, "unauthenticated")
	}
	payload.UserID = userID

	// call service
	if err := h.svc.CreateDonation(payload); err != nil {
		// map known service errors if you want; fallback internal error
		return utils.InternalServerErrorResponse(c, "failed creating donation")
	}
	return utils.CreatedResponse(c, "donation created successfully", nil)
}

// user/admin: GET /donations
// - admin: returns all donations
// - user: returns only own donations
func (h *DonationController) GetAllDonations(c echo.Context) error {
	userID, _ := utils.GetUserID(c) // unauthenticated => 0,false
	isAdm := utils.IsAdmin(c)

	// require auth for user-level listing; admin may call even without user_id set by middleware
	if !isAdm {
		if userID == 0 {
			return utils.UnauthorizedResponse(c, "unauthenticated")
		}
	}

	donations, err := h.svc.GetAllDonations(userID, isAdm)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "failed fetching donations")
	}
	return utils.SuccessResponse(c, "donations fetched", donations)
}

// user/admin: GET /donations/:id (owner or admin)
func (h *DonationController) GetDonationByID(c echo.Context) error {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.BadRequestResponse(c, "invalid id")
	}

	d, err := h.svc.GetDonationByID(uint(id64))
	if err != nil {
		if errors.Is(err, service.ErrDonationNotFound) {
			return utils.NotFoundResponse(c, "donation not found")
		}
		return utils.InternalServerErrorResponse(c, "failed fetching donation")
	}

	// permission check: owner or admin
	if !utils.IsAdmin(c) {
		userID, ok := utils.GetUserID(c)
		if !ok {
			return utils.UnauthorizedResponse(c, "unauthenticated")
		}
		if d.UserID != userID {
			return utils.ForbiddenResponse(c, "forbidden")
		}
	}

	return utils.SuccessResponse(c, "donation fetched", d)
}

// user/admin: PUT /donations/:id (only owner or admin can update)
func (h *DonationController) UpdateDonation(c echo.Context) error {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.BadRequestResponse(c, "invalid id")
	}

	var payload dto.DonationDTO
	if err := c.Bind(&payload); err != nil {
		return utils.BadRequestResponse(c, "invalid payload")
	}
	payload.ID = uint(id64)

	userID, ok := utils.GetUserID(c)
	if !ok {
		return utils.UnauthorizedResponse(c, "unauthenticated")
	}
	isAdm := utils.IsAdmin(c)

	if err := h.svc.UpdateDonation(payload, userID, isAdm); err != nil {
		if errors.Is(err, service.ErrDonationNotFound) {
			return utils.NotFoundResponse(c, "donation not found")
		}
		if errors.Is(err, service.ErrForbidden) {
			return utils.ForbiddenResponse(c, "forbidden")
		}
		return utils.InternalServerErrorResponse(c, "failed updating donation")
	}
	return utils.SuccessResponse(c, "donation updated", nil)
}

// user/admin: DELETE /donations/:id (only owner or admin can delete)
func (h *DonationController) DeleteDonation(c echo.Context) error {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.BadRequestResponse(c, "invalid id")
	}

	userID, ok := utils.GetUserID(c)
	if !ok {
		return utils.UnauthorizedResponse(c, "unauthenticated")
	}
	isAdm := utils.IsAdmin(c)

	if err := h.svc.DeleteDonation(uint(id64), userID, isAdm); err != nil {
		if errors.Is(err, service.ErrDonationNotFound) {
			return utils.NotFoundResponse(c, "donation not found")
		}
		if errors.Is(err, service.ErrForbidden) {
			return utils.ForbiddenResponse(c, "forbidden")
		}
		return utils.InternalServerErrorResponse(c, "failed deleting donation")
	}
	return utils.NoContentResponse(c)
}

func (h *DonationController) PatchDonation(c echo.Context) error {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return utils.BadRequestResponse(c, "invalid id")
	}

	var payload dto.DonationDTO
	if err := c.Bind(&payload); err != nil {
		return utils.BadRequestResponse(c, "invalid payload")
	}
	payload.ID = uint(id64)

	userID, ok := utils.GetUserID(c)
	if !ok {
		return utils.UnauthorizedResponse(c, "unauthenticated")
	}
	isAdm := utils.IsAdmin(c)

	if err := h.svc.PatchDonation(payload, userID, isAdm); err != nil {
		if errors.Is(err, service.ErrDonationNotFound) {
			return utils.NotFoundResponse(c, "donation not found")
		}
		if errors.Is(err, service.ErrForbidden) {
			return utils.ForbiddenResponse(c, "forbidden")
		}
		return utils.InternalServerErrorResponse(c, "failed patching donation")
	}
	return utils.SuccessResponse(c, "donation patched", nil)
}
