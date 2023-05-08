package models

type User struct {
	Id          int
	Items       []Items
	Username    string
	Email       string
	Password    string
	SocialMedia string
	Cash        int
}
type Admin struct {
	User
}
type Client struct {
	User
}

type Items struct {
	Id          string
	Name        string
	Description string
	Price       int
	Rating      int
	SellerInfo  string
}
type Rating struct {
	UserID string
	ItemID string
	Rating int
}

type Comments struct {
	UserID  string
	ItemID  string
	Comment string
}
type Roles struct {
}

type Order struct {
	UserId      string
	ItemId      string
	OrderStatus string
}
