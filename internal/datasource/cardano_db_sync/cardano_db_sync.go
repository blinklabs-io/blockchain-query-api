package cardano_db_sync

import (
	"errors"
	"fmt"
	"github.com/cloudstruct/blockchain-query-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CardanoDbSync struct {
	cfg *config.DatasourceCardanoDbSync
	db  *gorm.DB
}

var cardanoDbSync *CardanoDbSync

func New(cfg *config.Config) (*CardanoDbSync, error) {
	c := &CardanoDbSync{cfg: &cfg.Datasource.CardanoDbSync}
	err := c.Connect()
	if err != nil {
		return nil, err
	}
	cardanoDbSync = c
	return c, nil
}

func (c *CardanoDbSync) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.cfg.Host,
		c.cfg.Port,
		c.cfg.Username,
		c.cfg.Password,
		c.cfg.DbName,
		c.cfg.SslMode,
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func GetHandle() *gorm.DB {
	return cardanoDbSync.db
}

func IsRecordNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
