package db

import (
	"github.com/doornoc/dsbd-wg/pkg/core"
	"github.com/doornoc/dsbd-wg/pkg/core/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&core.Client{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Add(client []*core.Client) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}

	return db.Create(client).Error
}

func Delete(publicKey string) error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Delete(&[]core.Client{{PublicKey: publicKey}}).Error
}

func DeleteAll() error {
	db, err := ConnectDB()
	if err != nil {
		return err
	}
	return db.Exec("DELETE FROM client").Error
}
