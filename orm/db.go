package orm

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {

	/* SET TIME ZONE */
	time.LoadLocation("Asia/Bangkok")
	/* SET TIME ZONE */

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbIsPaseTime := os.Getenv("DB_IS_PARSE_TIME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?parseTime=" + dbIsPaseTime + "&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connect db already!")

	// Migrate the schema
	Db.AutoMigrate(
	// &models.Users{},

	)

}
