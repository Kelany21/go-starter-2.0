package datebase_control

import (
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/infrastructure"
	"golang-ddd-starter/infrastructure/seeder"
	"os"
	"strconv"
)

func MigrateAllTable() {
	deleteTables, _ := strconv.ParseBool(os.Getenv("DROP_ALL_TABLES"))
	if deleteTables {
		DbDrop()
	}
	infrastructure.DB.AutoMigrate(&models.User{})
}

/**
* drop all tables
 */

type query struct {
	Query string
}

func DbDrop() {
	var query []query
	infrastructure.DB.Table("information_schema.tables").Select("concat('DROP TABLE IF EXISTS `', table_name, '`;') as query").Where("table_schema = ? " , "starter").Find(&query)
	for _ , q :=range query{
		infrastructure.DB.Exec(q.Query)
	}
}

func Seed() {
	seeder.UserSeeder()
}