package seeder

import (
	"github.com/google/uuid"
	"golang-ddd-starter/domian/models"
	"golang-ddd-starter/helpers"
	"golang-ddd-starter/infrastructure"
	"syreclabs.com/go/faker"
)

/***
*	Seed Function must Have the same file Name then Add Seeder key word
* 	Example :  file is user function must be UserSeeder
 */
func UserSeeder() {
	newUser(true)
	for i := 0; i < 10; i++ {
		newUser(false)
	}
}

/**
* fake data and create data base
 */
func newUser(admin bool) {
	data := models.User{
		Email:    faker.Internet().Email(),
		Password: faker.Internet().Password(8, 14),
		Name:     faker.Internet().UserName(),
		Block:    2,
		Role:     1,
	}
	if admin {
		data.UUID, _ = uuid.Parse("e8780570-f510-4543-a1c7-a28df8ba25dc")
		data.Role = 2
		data.Email = "admin@admin.com"
		data.Password, _ = helpers.HashPassword("admin@admin.com")
		data.Token = "$2a$10$VqGA1Sr85Df6zX36w/iKv.JMOXkshmaWytk.njJFL6wGZiazVHd9i"
	}
	infrastructure.DB.Create(&data)
}
