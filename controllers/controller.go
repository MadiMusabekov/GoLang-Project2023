package controllers

import (
	Init "GoProject2023/init"
	"GoProject2023/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ItemsShow(c *gin.Context) {
	var items []models.Item
	Init.DB.Find(&items)

	c.JSON(200, gin.H{
		"items": items,
	})
}

func CommentsShow(c *gin.Context) {
	var comments []models.Comment
	Init.DB.Find(&comments)

	c.JSON(200, gin.H{
		"comments": comments,
	})
}
func RatingsShow(c *gin.Context) {
	var ratings []models.Rating
	Init.DB.Find(&ratings)

	c.JSON(200, gin.H{
		"ratings": ratings,
	})
}

func UserRegistration(c *gin.Context) {
	var body struct {
		//ID          uint
		Items       []models.Item
		Username    string
		Email       string
		Password    string
		SocialMedia string
		Cash        int
		Comments    []models.Comment `gorm:"foreignKey:UserID"`
		Orders      []models.Order   `gorm:"foreignKey:UserID"`
		Ratings     []models.Rating  `gorm:"foreignKey:UserID"`
	}
	c.Bind(&body)
	user := models.User{
		//ID: body.ID,
		Items: body.Items, Username: body.Username,
		Email: body.Email, Password: body.Password, SocialMedia: body.SocialMedia,
		Cash: body.Cash, Comments: make([]models.Comment, 0), Orders: make([]models.Order, 0), Ratings: make([]models.Rating, 0),
	}
	Init.DB.Create(&user)
	c.JSON(200, gin.H{
		"user": user,
	})
}
func UsersShow(c *gin.Context) {
	var users []models.User
	Init.DB.Preload("Comments").Find(&users)
	c.JSON(200, gin.H{
		"users": users,
	})
}

//	func Rating(c *gin.Context) {
//		var body struct {
//			UserID string
//			ItemID string
//			Rating int
//		}
//		c.Bind(&body)
//		var newRating float64
//		var user models.User
//		var item models.Item
//		Init.DB.First(&user, body.UserID)
//		Init.DB.First(&item, body.ItemID)
//		previousRatings := []float64{float64(item.Rating), newRating}
//		sum := 0.0
//		for _, r := range previousRatings {
//			sum += r
//		}
//		avgRating := sum / float64(len(previousRatings))
//		Init.DB.Model(&item).Update("Rating", avgRating)
//		//c.JSON(200, gin.H{
//		//  "users": users,
//		//})
//	}
func AddComment(c *gin.Context) {
	var body struct {
		UserID  uint
		ItemID  uint
		Comment string
	}
	c.Bind(&body)
	comment := models.Comment{
		UserID:  body.UserID,
		ItemID:  body.ItemID,
		Comment: body.Comment,
	}
	Init.DB.Create(&comment)

	//old version
	//var user models.User
	//Init.DB.Model(&models.User{}).Where("id = ?", body.UserID).Association("Comments").Append(&comment)

	//new version
	var user models.User
	Init.DB.First(&user, body.UserID)
	Init.DB.Model(&user).Association("Comments").Append(&comment)

	//old ver
	//var item models.Item
	//Init.DB.Model(&models.Item{}).Where("id = ?", body.ItemID).Association("Comments").Append(&comment)

	//new ver
	var item models.Item
	Init.DB.First(&item, body.ItemID)
	Init.DB.Model(&item).Association("Comments").Append(&comment)

	c.JSON(200, gin.H{
		"message": "Commented successfully",
	})
}

func RateItem(c *gin.Context) {
	var body struct {
		UserID uint
		ItemID uint
		Rating float64
	}

	c.Bind(&body)

	// Get the user and item from the database
	var user models.User
	var item models.Item
	Init.DB.First(&user, body.UserID)
	Init.DB.First(&item, body.ItemID)

	// Create a new rating
	newRating := models.Rating{
		UserID: body.UserID,
		ItemID: body.ItemID,
		Rating: body.Rating,
	}

	// Add the rating to the item and save it
	Init.DB.Model(&item).Association("Ratings").Append(&newRating)
	Init.DB.Save(&item)

	// Add the rating to the user and save it
	Init.DB.Model(&user).Association("Ratings").Append(&newRating)
	Init.DB.Save(&user)

	var previousRatings []float64
	var ratings []float64
	Init.DB.Where("item_id = ?", body.ItemID).Pluck("rating", &ratings)

	// here we are iterating through an array
	//consisting of rating column, so that we can apply
	// an average rating to an item during each rating
	sum := 0.0
	for _, r := range ratings {
		sum += r
	}
	avgRating := sum / float64(len(previousRatings))

	Init.DB.Model(&item).Update("Rating", avgRating)

	// Return a success message
	c.JSON(200, gin.H{
		"message": "Item rated successfully",
	})
}

func ItemCreate(c *gin.Context) {
	var body struct {
		UserID      uint
		Name        string
		Description string
		Price       int
		Rating      float64
		SellerInfo  string
		//Comments    []models.Comment `gorm:"foreignKey:ItemID;references:ID"`
		//Orders      []models.Order   `gorm:"foreignKey:ItemID;references:ID"`
		Ratings []models.Rating `gorm:"foreignKey:ItemID;references:ID"`
	}
	c.Bind(&body)
	item := models.Item{
		UserID: body.UserID,
		Name:   body.Name, Description: body.Description,
		Price: body.Price, Rating: body.Rating,
		SellerInfo: body.SellerInfo,
		//Comments:   body.Comments,
		//Orders:     body.Orders,
		Ratings: body.Ratings,
	}
	//Init.DB.Create(&item)
	if err := Init.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"item": item,
	})
}
