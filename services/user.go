package services

import (
    "errors"

    "github.com/gabrielgmusskopf/crud_go/models"
    "github.com/gabrielgmusskopf/crud_go/db"
)

func CreateUser(name string) *models.User {
    user := models.User{Name: name}
    db.DB.Create(&user)

    return &user
}

func GetAllUsers() *[]models.User {
    var us *[]models.User
    db.DB.Find(&us)

    return us
}

func GetUserById(id int) (*models.User, error) {
    var user models.User
    result := db.DB.Where("id = ?", id).Find(&user)

    if result.Error != nil {
        return nil, errors.New("user not found")
    }

    return &user, nil
}

func DeleteUserById(id int) error {
    result := db.DB.Delete(&models.User{}, id); if result.Error != nil {
        return errors.New("user not foud")
    }

    return nil
}

func UpdateUser(id int, name string) (*models.User, error) {
    var user models.User
    result := db.DB.Where("id = ?", id).Find(&user)

    if result.Error != nil {
        return nil, errors.New("user not foud")
    }

    db.DB.Model(&user).Updates(models.User{Name: name})

    return &user, nil
}

