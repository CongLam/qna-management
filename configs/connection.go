package configs

import (
	"os"

	model "qna-management/models"
	"qna-management/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	log := utils.GetLogger()
	databaseURI := make(chan string, 1)

	if os.Getenv("GO_ENV") != "production" {
		databaseURI <- utils.GodotEnv("DATABASE_URI_DEV")
	} else {
		databaseURI <- os.Getenv("DATABASE_URI_PROD")
	}

	db, err := gorm.Open(postgres.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer log.Info("Connection to Database Failed")
		log.Fatal(err.Error())
	}
	log.Info("Connection to Database Successfully")
	if os.Getenv("GO_ENV") == "production" {
		log.Info("Connection to Database Successfully")
	}

	err = db.AutoMigrate(
		&model.EntityUsers{},
	)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
