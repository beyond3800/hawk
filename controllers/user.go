package controllers

import (
	"fmt"

	"github.com/beyond3800/hawk/core/database"
	"github.com/beyond3800/hawk/core/hawk"
	"github.com/beyond3800/hawk/models"
)

type UserController struct{
}

func (u *UserController) CreateUser(c *hawk.Context) {
	var user models.Users
	
	if err := c.BindAndValidate(&user); err != nil {
		c.ValidationError(err)
		return
	}
	// _,err :=database.HawkDB().Table("users").Insert(map[string]any{
	// 	"first_name": user.Name,
	// 	"last_name":"bass",
	// 	"email": user.Email,
	// 	"password": user.Password,
	// })
	// if err != nil{
	// 	c.JSON(501 , err)
	// 	return
	// }

}

func (u *UserController) Show(c *hawk.Context){
	var users []models.User
	
	err :=database.HawkDB().Table("users").
    Select("*").
	Where("id",5).
    Get(&users)

	if err != nil {
		fmt.Println("Error fetching user:", err)
	}
	// c.Cookie("schoolToken")
	fmt.Printf("User: %+v\n", users)
	
}
