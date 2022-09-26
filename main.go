package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// Controller

// Get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succes get all users",
		"users":    users,
	})
}

// Get user by id
func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "url parameter harus angka",
		})
	}

	var user []User
	user = append(user, users[id-1])
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get user by id",
		"user":     user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "url parameter harus angka",
		})
	}

	users = append(users[:id-1], users[id:]...)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "delete data success",
		"users":    users,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "url parameter harus angka",
		})
	}

	user := User{}
	c.Bind(&user)

	users[id-1].Name = user.Name
	users[id-1].Email = user.Email
	users[id-1].Password = user.Password

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "update data success",
		"users":    user,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "create data success",
		"users":    user,
	})
}

func main() {
	e := echo.New()

	// routing
	// access by "localhost:8080/v1/users"
	userRoute := e.Group("/v1")
	userRoute.GET("/users", GetUsersController)
	userRoute.GET("/users/:id", GetUserController)
	userRoute.POST("/users", CreateUserController)
	userRoute.DELETE("/users/:id", DeleteUserController)
	userRoute.PUT("/users/:id", UpdateUserController)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}
