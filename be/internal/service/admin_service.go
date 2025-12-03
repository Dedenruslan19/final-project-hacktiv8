package service

import (
	"log"
	"milestone3/be/internal/dto"
)

type AdminRepository interface {
	CountPayment() (resp dto.TotalPayment, err error)
	CountDonation() (resp dto.TotalDonation, err error)
	CountArticle() (resp dto.TotalArticle, err error)
	// CountAuction() (resp dto.TotalAuction, err error)
}

type AdminServ struct {
	adminRepo AdminRepository
}

func NewAdminService(ar AdminRepository) *AdminServ {
	return &AdminServ{adminRepo: ar}
}

func (as *AdminServ) AdminDashboard() (resp dto.AdminDashboardResponse, err error) {
	article, err := as.adminRepo.CountArticle()
	if err != nil {
		log.Printf("error count article %s", err)
		return dto.AdminDashboardResponse{}, err
	}

	donation, err := as.adminRepo.CountDonation() 
	if err != nil {
		log.Printf("error count donation %s", err)
		return dto.AdminDashboardResponse{}, err
	}

	payment, err := as.adminRepo.CountPayment()
	if err != nil {
		log.Printf("error count payment %s", err)
		return dto.AdminDashboardResponse{}, err
	}

	// auction, err := as.adminRepo.CountAuction(); 
	// if err != nil {
	// 	log.Printf("error count payment %s", &err)
	// 	return err
	// }
	respon := dto.AdminDashboardResponse{
		TotalArticle: article.Count,
		TotalDonation: donation.Count,
		TotalPayment: payment.Count,
		// TotalAuction: auction.Count,
	}

	return respon, nil
}

// work in progress (WIP)
// func (as *AdminServ) AdminReport() (err error) { }