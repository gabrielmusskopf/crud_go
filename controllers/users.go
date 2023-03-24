package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
    "strconv"

    "github.com/gabrielgmusskopf/crud_go/services"
)

type UserController struct {}

type CreateUserInput struct {
    Name string `json:"name"`
}

func (u *UserController) GetAll(c *gin.Context)  {
    us := services.GetAllUsers()

    c.JSON(http.StatusOK, gin.H{
        "data": us,
    }) 

    return
}

func (u *UserController) GetById(c *gin.Context)  {
    id := getIdParam(c)

    user, error := services.GetUserById(id); if error != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": error,
        }) 

        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data": user,
    }) 

    return
}


func (u *UserController) Create(c *gin.Context)  {

    var input CreateUserInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := services.CreateUser(input.Name)

    c.JSON(http.StatusOK, gin.H{
        "data": user,
    }) 

    return
}

func (u *UserController) Delete(c *gin.Context)  {

    id := getIdParam(c)

    error := services.DeleteUserById(id); if error != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": error,
        }) 

        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "user deleted.",
    }) 
}

func (u *UserController) Update(c *gin.Context)  {

    var input CreateUserInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := getIdParam(c)

    user, error := services.UpdateUser(id, input.Name); if error != nil{
        c.JSON(http.StatusBadRequest, gin.H{
            "message": error,
        }) 

        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "user updated.",
        "data": user,
    }) 
}

func getIdParam(c *gin.Context) int {
    id, err := strconv.Atoi(c.Param("id")); if err != nil {
        panic(err)
    }

    return id
}
