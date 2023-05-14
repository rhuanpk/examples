package database

import (
	"fmt"
	"log"

	"backend/internal/consts"
	"backend/internal/models"

	// <user>:<password>@<protocol>(<host>:<port>)/<database>[?parseTime=true]
	// root:root@tcp(localhost:3306)/test
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Instance is the singleton instance of database.
var (
	Instance      *gorm.DB
	databaseError error
)

func init() {

	stringConnection := fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s",
		consts.DBUSER,
		consts.DBPASSWORD,
		consts.DBPROTOCOL,
		consts.DBHOST,
		consts.DBPORT,
		consts.DBNAME,
	)

	Instance, databaseError = gorm.Open(mysql.Open(stringConnection), &gorm.Config{})
	if databaseError != nil {
		log.Fatalln("can't establish database connection!")
	}

	err := Instance.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalln("can't database auto migration!")
	}

}
