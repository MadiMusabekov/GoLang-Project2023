package model

type User struct {
	Id          int
	Username    string
	Email       string
	Password    string
	SocialMedia string
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
	User_ID string
	Item_ID string
	Rating  int
}

type Comments struct {
	User_ID string
	Item_ID string
	Comment string
}
type Roles struct {
}
