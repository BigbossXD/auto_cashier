package orm

import (
	"fmt"
	"os"

	"github.com/BigbossXD/auto_cashier/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func InitDB() {

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
	fmt.Println("Connect Database Already...!")

	// Migrate the schema
	Db.AutoMigrate(
		&models.CashierConfigs{},
	)

	// Init Configs
	moneyValue := []float32{1000, 500, 100, 50, 20, 10, 5, 1, 0.25}
	maximumAmount := []int32{10, 20, 15, 20, 30, 20, 20, 20, 50}
	for i, v := range moneyValue {
		configs := &models.CashierConfigs{}
		result := Db.Where("money_value = ?", v).First(&configs)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				configs.MoneyValue = v
				configs.MaximumAmount = maximumAmount[i]
				configs.CurrentAmount = 0
				Db.Save(configs)
			} else {
				panic(result.Error.Error())
			}
		}

	}

}
