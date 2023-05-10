package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//ID          uint   `gorm:"PrimaryKey"`
	Items       []Item `gorm:"foreignKey:UserID; references:ID"`
	Username    string
	Email       string
	Password    string
	SocialMedia string
	Cash        int
	Comments    []Comment `gorm:"foreignKey:UserID;references:ID"`
	Orders      []Order   `gorm:"foreignKey:UserID;references:ID"`
	Ratings     []Rating  `gorm:"foreignKey:UserID;references:ID"`
}
type Admin struct {
	User
}
type Client struct {
	User
}

type Item struct {
	gorm.Model
	//ID          uint `gorm:"PrimaryKey"`
	UserID      User `gorm:"foreignKey: ID"`
	Name        string
	Description string
	Price       int
	Rating      float64
	SellerInfo  string
	Ratings     []Rating `gorm:"foreignKey:ItemID;references:ID"`
}
type Rating struct {
	gorm.Model
	UserID User `gorm:"foreignKey:ID"`
	ItemID Item `gorm:"foreignKey: ID"`
	Rating float64
}

type Comment struct {
	//ID uint `gorm:"primary_key"`
	gorm.Model
	//ID      uint `gorm:"PrimaryKey"`
	UserID  User `gorm:"foreignKey: ID"`
	ItemID  Item `gorm:"foreignKey: ID"`
	Comment string
}
type Roles struct {
}

type Order struct {
	//ID uint `gorm:"primary_key"`
	gorm.Model
	//ID          uint `gorm:"PrimaryKey"`
	UserID      User `gorm:"foreignKey: ID"`
	ItemID      Item `gorm:"foreignKey: ID"`
	OrderStatus string
}
