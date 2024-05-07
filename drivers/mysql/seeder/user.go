package seeder

import (
	"errors"
	"lawan-tambang-liar/entities"
	"lawan-tambang-liar/utils"

	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	if err := db.First(&entities.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Create array of entities.User
		hash, _ := utils.HashPassword("user")
		users := []entities.User{
			{
				Username: "user_1",
				Password: hash,
				Email:    "user_1@gmail.com",
			},
			{
				Username: "user_2",
				Password: hash,
				Email:    "user_2@gmail.com",
			},
			{
				Username: "user_3",
				Password: hash,
				Email:    "user_3@gmail.com",
			},
			{
				Username: "user_4",
				Password: hash,
				Email:    "user_4@gmail.com",
			},
			{
				Username: "user_5",
				Password: hash,
				Email:    "user_5@gmail.com",
			},
			{
				Username: "user_6",
				Password: hash,
				Email:    "user_6@gmail.com",
			},
			{
				Username: "user_7",
				Password: hash,
				Email:    "user_7@gmail.com",
			},
		}

		if err := db.CreateInBatches(&users, len(users)).Error; err != nil {
			panic(err)
		}
	}
}
