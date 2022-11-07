package database

import (
	"gitlab.com/meta-node/mail/core/entities"
)

func Migration(dbConn *dbConnection) {
	dbConn.DB.AutoMigrate(
		entities.User{},
		entities.Email{},
		entities.File{},
		entities.EmailManager{},

		// RBAC
		entities.Role{},
		entities.Permission{},
	)
}
