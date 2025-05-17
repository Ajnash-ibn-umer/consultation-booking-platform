package services

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/pkg/dto"
	"ajnash-ibn-ummer/wewithyou/pkg/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func SignIn(dto *dto.SignInDto) (*common.Resp, error) {

	db := config.GetDB()
	var userData models.User
	resultErr := db.Where(&models.User{Email: dto.Email}).Find(&userData).Error
	if resultErr != nil || userData.ID == 0 {
		return common.NewErrorResponse(
			&common.E{Message: "User not found", Errors: resultErr},
		), resultErr
	}
	fmt.Print("user datan-----", userData, "\n ---------")

	passwordCheck, errInhash := CheckPasswordHash(dto.Password, userData.Password)

	if passwordCheck != true {
		return common.NewErrorResponse(
			&common.E{Message: "Password doesn`t match", Errors: errInhash},
		), errInhash
	}

	userData.Password = ""
	return common.NewSuccessResponse(
		&common.S{Message: "User got it", Data: userData},
	), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil, err
}
