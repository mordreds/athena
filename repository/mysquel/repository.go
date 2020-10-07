package mysquel

import (
	online_application "athena/internal/student/online-application"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var err error

func DBSetup() {
	DB, err = gorm.Open(mysql.Open(DbURL(BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed Setting DB:: ", err)
	}

	DB.AutoMigrate(&online_application.OnlineApplication{})
}