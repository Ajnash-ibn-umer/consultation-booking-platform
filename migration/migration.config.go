package migration

import (
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/pkg/models"
)

func MigrateEntities() {
	db := config.GetDB()

	db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.RolePermission{})

}
