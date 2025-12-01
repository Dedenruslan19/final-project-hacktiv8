package entity

type Payment struct {
	Id int
	UserId int
	User Users `gorm:"foreignKey:UserId;references:Id"`
	AuctionItemId float64
	Amount int
}