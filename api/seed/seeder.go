package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/abydarts/tennet-go-api/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "user1",
		Email:    "user@email.com",
		Password: "123456",
	}
}

var wallets = []models.Wallet{
	models.Wallet{
		Name:   "E-money"
	}
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Wallet{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		
		err = db.Debug().Model(&models.Wallet{}).Create(&wallets[i]).Error
		if err != nil {
			log.Fatalf("cannot seed wallet table: %v", err)
		}
	}
}