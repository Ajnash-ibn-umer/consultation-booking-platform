package services

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/pkg/dto"
	"ajnash-ibn-ummer/wewithyou/pkg/models"
	"errors"
	"fmt"
)

func CreateUser(user *dto.CreateUserDto, uid uint) (*common.Resp, error) {

	db := config.GetDB()
	validator := config.ValidatorInt()
	if err := validator.Var(user.Email, "required,email"); err != nil {
		return common.NewErrorResponse(&common.E{
			Message: "Email is not valid Email",
			Errors:  err,
		}), err
	}

	if user.Name == "" {
		fmt.Print("User Name is Required")
		return common.NewErrorResponse(&common.E{
			Message: "User Name is Required",
			Errors:  errors.New("User name is required"),
		}), errors.New("User name is required")
	}

	hashedPassword, errHash := HashPassword(user.Password)
	if errHash != nil {
		return common.NewErrorResponse(&common.E{
			Message: "Password storing Failed",
			Errors:  errHash,
		}), errHash
	}

	fmt.Println("user", user)
	newUser := models.User{
		Name:         user.Name,
		Email:        user.Email,
		Password:     hashedPassword,
		UserName:     user.UserName,
		AuthMethod:   1,
		IsActive:     false,
		RoleID:       nil,
		UserType:     1,
		ProfileImage: "",
		Age:          user.Age,
		Gender:       "Male",
	}

	err := db.Create(&newUser).Error
	if err != nil {

		return common.NewErrorResponse(&common.E{
			Message: err.Error(),
			Errors:  err,
		}), err

	}
	return common.NewSuccessResponse(&common.S{
		Message: "Success",
		Data:    &newUser,
	}), nil
}
