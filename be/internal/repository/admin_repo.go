package repository

import (
	"context"
	"milestone3/be/internal/dto"
	"milestone3/be/internal/entity"

	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
	ctx context.Context
}

func NewAdminRepository(db *gorm.DB, ctx context.Context) *AdminRepo {
	return &AdminRepo{db: db, ctx: ctx}
}

var count int64

//count total transaction
func (ar *AdminRepo) CountPayment() (resp dto.TotalPayment, err error) {
var payment entity.Payment
	if err := ar.db.WithContext(ar.ctx).Model(&payment).Count(&count).Error; err != nil {
		return dto.TotalPayment{}, err
	}

	return resp, nil
}
 
// //count total donation
func (ar *AdminRepo) CountDonation() (resp dto.TotalDonation, err error) {
	var donation entity.Donation
	if err := ar.db.WithContext(ar.ctx).Model(&donation).Count(&count).Error; err != nil {
		return dto.TotalDonation{}, err
	}

	return resp, nil
}

// count total auction
// work in progress (WIP)
// func (ar *AdminRepo) CountAuction() (resp dto.TotalAuction, err error) {
// var auction entitiy.Auction
// if err := ar.db.WithContext(ar.ctx).Model(&auction).Count(&count).Error; err != nil {
// 		return dto.TotalAuction{}, err
// }
// 
// return resp, nil
// }

// //count total article
func (ar *AdminRepo) CountArticle() (resp dto.TotalArticle, err error) {
	var article entity.Article
	if err := ar.db.WithContext(ar.ctx).Model(&article).Count(&count).Error; err != nil {
		return dto.TotalArticle{}, err
	}

	return resp, nil
}

// for reporting endpoint //
// work in progress (WIP)