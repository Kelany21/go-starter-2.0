package main

import (
	"github.com/bykovme/gotrans"
	"golang-ddd-starter/app"
	"golang-ddd-starter/app/validations"
	"golang-ddd-starter/infrastructure"
	"golang-ddd-starter/infrastructure/datebase_control"
	"golang-ddd-starter/providers"
)



func main()  {
	/**
	* start multi language
	 */
	err := gotrans.InitLocales("public/trans")
	if err != nil {
		panic(err)
	}
	/**
	* add custom role to validation
	*
	 */
	validations.Init()
	/**
	* connect with data base logic you can edit env file to
	* change any connection params
	 */
	infrastructure.ConnectToDatabase()
	app.InitApps(infrastructure.DB)
	/**
	* drop All tables and migrate
	* to stop delete tables make DROP_ALL_TABLES false in env file
	* if you need to stop auto migration just stop this line
	 */
	datebase_control.MigrateAllTable()
	/**
	* this function will open seeders folder look inside all files
	* search for seeders function and seed execute these function
	* if you need to stop seeding you can stop this line
	 */
	datebase_control.Seed()
	/**
	* Run gin framework
	* add middleware
	* run routing
	* serve app
	 */
	providers.Run()
}
