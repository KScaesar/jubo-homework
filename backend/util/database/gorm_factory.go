package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/KScaesar/jubo-homework/backend/util/errors"
)

func NewGormPgsql(cfg *DbConfig) (*WrapperGorm, error) {
	cfg.setDefault()

	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s lock_timeout=5000 idle_in_transaction_session_timeout=10000 sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	gormDB, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		return nil, errors.Join3rdParty(errors.ErrSystem, err)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, errors.Join3rdParty(errors.ErrSystem, err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, errors.Join3rdPartyWithMsg(errors.ErrSystem, err, "ping test connect")
	}

	sqlDB.SetMaxOpenConns(cfg.MaxConn)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConn)

	if cfg.DebugMode {
		gormDB = gormDB.Debug()
	}

	return NewWrapperGorm(gormDB), nil
}
