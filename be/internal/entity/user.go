package entity

type Users struct {
	Id int
	Name string
	Email string
	Password string `json:"-"`
	RoleId int
	Role Role `gorm:"foreignKey:RoleId;references:Id"`
}

type Role struct {
	Id int
	Name string
}