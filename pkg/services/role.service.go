package services

import (
	"ajnash-ibn-ummer/wewithyou/common"
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/pkg/models"
	"fmt"
	"net/http"
)

func RoleCreate(roleName *string) (*common.Resp, error) {

	db := config.GetDB()
	roleData := models.Role{
		Name: *roleName,
	}

	err := db.Create(&roleData).Error
	fmt.Print("role resp", &roleData)
	rolePermissions := []*models.RolePermission{
		{
			RoleID:   int(roleData.ID),
			Actions:  "R,W",
			Path:     "*",
			Resource: "*",
		},
		{
			RoleID:   int(roleData.ID),
			Actions:  "R,W",
			Path:     "/admin",
			Resource: "admin",
		},
	}
	errInPermission := db.Create(&rolePermissions).Error

	if err != nil {
		return common.NewErrorResponse(&common.E{
			Message: "Role creation failed",
			Status:  http.StatusBadGateway,
			Errors:  err,
		}), err
	}
	if errInPermission != nil {
		return common.NewErrorResponse(&common.E{
			Message: "Role Permission creation failed",
			Status:  http.StatusBadGateway,
			Errors:  errInPermission,
		}), err
	}
	return common.NewSuccessResponse(&common.S{
		Message: "Role created",
		Status:  http.StatusOK,
		Data:    &roleData,
	}), nil
}

func RoleList() (*common.Resp, error) {

	db := config.GetDB()

	var roles []*models.Role
	err := db.Preload("Permissions").Find(&roles).Error

	if err != nil {
		return common.NewErrorResponse(&common.E{
			Message: "Lisitng failed",
			Status:  http.StatusBadRequest,
			Errors:  err,
		}), err
	}

	return common.NewSuccessResponse(&common.S{
		Message: "Role listed",
		Status:  http.StatusOK,
		Data:    &roles,
	}), nil
}
