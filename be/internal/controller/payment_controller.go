package controller

import (
	"milestone3/be/internal/dto"
	"milestone3/be/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type PaymentService interface {
	CreatePayment(req dto.PaymentRequest) (res dto.PaymentResponse, err error)
}

type PaymentController struct {
	paymentService PaymentService
	validate *validator.Validate
}

func NewPaymentController(validate *validator.Validate, ps PaymentService) *PaymentController {
	return &PaymentController{paymentService:ps, validate: validate}
}

func (pc *PaymentController) CreatePayment(c echo.Context) error {
	req := new(dto.PaymentRequest)

	if err := c.Bind(req); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	if err := pc.validate.Struct(req); err != nil {
		return utils.BadRequestResponse(c, err.Error())
	}

	resp, err := pc.paymentService.CreatePayment(*req)
	if err != nil {
		return utils.InternalServerErrorResponse(c, "internal server error")
	}

	return utils.CreatedResponse(c, "create", resp)
}