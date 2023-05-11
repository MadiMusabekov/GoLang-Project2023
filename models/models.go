package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username    string
	Email       string
	Password    string
	SocialMedia string
	Cash        int
	Items       []Item    `gorm:"foreignKey:UserID"`
	Comments    []Comment `gorm:"foreignKey:UserID"`
	Orders      []Order   `gorm:"foreignKey:UserID"`
	Ratings     []Rating  `gorm:"foreignKey:UserID"`
}

type Item struct {
	gorm.Model
	UserID uint // foreign key reference
	//User        User `gorm:"foreignKey:UserID"`
	Name        string
	Description string
	Price       int
	Rating      float64
	Comment     string
	SellerInfo  string
	Comments    []Comment `gorm:"foreignKey:ItemID"`
	Orders      []Order   `gorm:"foreignKey:ItemID"`
	Ratings     []Rating  `gorm:"foreignKey:ItemID"`
}

type Rating struct {
	gorm.Model
	UserID uint // foreign key reference
	//User   User `gorm:"foreignKey:UserID"`
	ItemID uint // foreign key reference
	//Item   Item `gorm:"foreignKey:ItemID"`
	Rating float64
}

type Comment struct {
	gorm.Model
	UserID uint // foreign key reference
	//User    User `gorm:"foreignKey:UserID"`
	ItemID uint // foreign key reference
	//Item    Item `gorm:"foreignKey:ItemID"`
	Comment string
}

type Order struct {
	gorm.Model
	UserID uint // foreign key reference
	//User        User `gorm:"foreignKey:UserID"`
	ItemID uint // foreign key reference
	//Item        Item `gorm:"foreignKey:ItemID"`
	OrderStatus bool
}
