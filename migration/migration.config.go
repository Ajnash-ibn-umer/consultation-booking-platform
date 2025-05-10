package migration

import (
	"ajnash-ibn-ummer/wewithyou/config"
	"ajnash-ibn-ummer/wewithyou/pkg/auth"
	"ajnash-ibn-ummer/wewithyou/pkg/users"
)

func MigrateEntities() {
	db := config.GetDB()

	db.AutoMigrate(
		&users.User{},
		&auth.Role{},
		&auth.RolePermission{})

}
