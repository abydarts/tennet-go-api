package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Wallet struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Wallet) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Title))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Wallet) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	
	return nil
}

func (p *Wallet) SaveWallet(db *gorm.DB) (*Wallet, error) {
	var err error
	err = db.Debug().Model(&Wallet{}).Create(&p).Error
	if err != nil {
		return &Wallet{}, err
	}
	return p, nil
}

func (p *Wallet) FindAllWallet(db *gorm.DB) (*[]Wallet, error) {
	var err error
	wallets := []Wallet{}
	err = db.Debug().Model(&Wallet{}).Limit(100).Find(&wallets).Error
	if err != nil {
		return &[]Wallet{}, err
	}
	if len(wallets) > 0 {

	}
	return &wallets, nil
}

func (p *Wallet) FindWalletByID(db *gorm.DB, pid uint64) (*Wallet, error) {
	var err error
	err = db.Debug().Model(&Wallet{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Wallet{}, err
	}
	
	return p, nil
}

func (p *Wallet) UpdateAWallet(db *gorm.DB) (*Wallet, error) {

	var err error

	err = db.Debug().Model(&Wallet{}).Where("id = ?", p.ID).Updates(Wallet{Name: p.Name, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Wallet{}, err
	}
	
	return p, nil
}

func (p *Wallet) DeleteAWallet(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Wallet{}).Where("id = ?", pid).Take(&Wallet{}).Delete(&Wallet{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Wallet not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}