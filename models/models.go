package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Items       []Item `gorm:"foreignKey:UserID"`
	Username    string
	Email       string
	Password    string
	SocialMedia string
	Cash        int
	Comments    []Comment `gorm:"foreignKey:UserID"`
	Orders      []Order   `gorm:"foreignKey:UserID"`
	Ratings     []Rating  `gorm:"foreignKey:UserID"`
}

type Item struct {
	gorm.Model
	UserID      uint // foreign key reference
	User        User `gorm:"foreignKey:UserID"`
	Name        string
	Description string
	Price       int
	Rating      float64
	SellerInfo  string
	Ratings     []Rating `gorm:"foreignKey:ItemID"`
}

type Rating struct {
	gorm.Model
	UserID uint // foreign key reference
	User   User `gorm:"foreignKey:UserID"`
	ItemID uint // foreign key reference
	Item   Item `gorm:"foreignKey:ItemID"`
	Rating float64
}

type Comment struct {
	gorm.Model
	UserID  uint // foreign key reference
	User    User `gorm:"foreignKey:UserID"`
	ItemID  uint // foreign key reference
	Item    Item `gorm:"foreignKey:ItemID"`
	Comment string
}

type Order struct {
	gorm.Model
	UserID      uint // foreign key reference
	User        User `gorm:"foreignKey:UserID"`
	ItemID      uint // foreign key reference
	Item        Item `gorm:"foreignKey:ItemID"`
	OrderStatus string
}
