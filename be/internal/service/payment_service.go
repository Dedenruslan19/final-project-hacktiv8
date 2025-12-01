package service

import (
	"fmt"
	"log"
	"milestone3/be/internal/dto"
	"milestone3/be/internal/entity"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) (error)
	CreateMidtrans(payment entity.Payment, orderId string) (res dto.PaymentResponse, err error)
	CheckPaymentStatusMidtrans(transactionId string) (res dto.CheckPaymentStatusResponse, err error)
}

type PaymentServ struct {
	paymentRepo PaymentRepository
}

func NewPaymentService(pr PaymentRepository) *PaymentServ {
	return &PaymentServ{paymentRepo: pr}
}

func (ps *PaymentServ) CreatePayment(req dto.PaymentRequest, userId int) (res dto.PaymentResponse, err error) {
	//random id for order id
	uuid := uuid.New()
	orderId := fmt.Sprintf("YDR-%d", uuid.ID())
	
	payment := entity.Payment{
		Amount: req.Amount,
		UserId: userId,
		//hard code for now 
		AuctionItemId: 1,
	}

	if err := ps.paymentRepo.Create(&payment); err != nil {
		log.Printf("error create payment %s", err)
		return dto.PaymentResponse{}, err
	}

	log.Println("disini nih")
	resp, _ := ps.paymentRepo.CreateMidtrans(payment, orderId)

	return resp, nil
}

func (ps *PaymentServ) CheckPaymentStatusMidtrans(transactionId string) (res dto.CheckPaymentStatusResponse, err error) {
	resp, _:= ps.paymentRepo.CheckPaymentStatusMidtrans(transactionId)

	return resp, nil
}