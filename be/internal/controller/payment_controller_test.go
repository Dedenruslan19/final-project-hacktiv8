package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"milestone3/be/internal/dto"
	"milestone3/be/internal/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPaymentController_CreatePayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPaymentService(ctrl)
	validate := validator.New()
	controller := NewPaymentController(validate, mockService)

	tests := []struct {
		name           string
		auctionId      string
		requestBody    dto.PaymentRequest
		setupAuth      func(c echo.Context)
		setupMock      func()
		expectedStatus int
	}{
		{
			name:      "successful payment creation",
			auctionId: "1",
			requestBody: dto.PaymentRequest{
				Amount: 100000,
			},
			setupAuth: func(c echo.Context) {
				token := &jwt.Token{
					Claims: jwt.MapClaims{
						"id": float64(1),
					},
				}
				c.Set("user", token)
			},
			setupMock: func() {
				mockService.EXPECT().CreatePayment(gomock.Any(), 1, 1).Return(dto.PaymentResponse{
					OrderId:        "YDR-123",
					TransactionId:  "TXN-123",
					PaymentLinkUrl: "https://payment.link",
				}, nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:      "invalid auction id",
			auctionId: "invalid",
			requestBody: dto.PaymentRequest{
				Amount: 100000,
			},
			setupAuth: func(c echo.Context) {
				token := &jwt.Token{
					Claims: jwt.MapClaims{
						"id": float64(1),
					},
				}
				c.Set("user", token)
			},
			setupMock:      func() {},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/payments/auction/"+tt.auctionId, bytes.NewReader(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("auctionId")
			c.SetParamValues(tt.auctionId)

			tt.setupAuth(c)
			tt.setupMock()

			// Execute
			err := controller.CreatePayment(c)

			// Assert
			if tt.expectedStatus >= 400 {
				if err != nil {
					assert.Error(t, err)
				} else {
					assert.Equal(t, tt.expectedStatus, rec.Code)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, rec.Code)
			}
		})
	}
}

func TestPaymentController_GetAllPayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockPaymentService(ctrl)
	validate := validator.New()
	controller := NewPaymentController(validate, mockService)

	tests := []struct {
		name           string
		setupMock      func()
		expectedStatus int
		expectError    bool
	}{
		{
			name: "successful get all payments",
			setupMock: func() {
				mockService.EXPECT().GetAllPayment().Return([]dto.PaymentInfoResponse{
					{Id: 1, Amount: 100000},
					{Id: 2, Amount: 200000},
				}, nil)
			},
			expectedStatus: http.StatusOK,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/payments", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			tt.setupMock()

			// Execute
			err := controller.GetAllPayment(c)

			// Assert
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedStatus, rec.Code)
			}
		})
	}
}