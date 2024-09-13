package db

import (
	"context"
	"fmt"

	"github.com/ThisJohan/snapp-assignment/pkg/di"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const Service = "db"

type Config struct {
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port     int    `env:"POSTGRES_PORT" envDefault:"5432"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"postgres"`
	DBName   string `env:"POSTGRES_DB" envDefault:"postgres"`
	SSLMode  string `env:"POSTGRES_DB" envDefault:"disable"`
}

func (c *Config) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMode)
}

func OpenDBConnection(c Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.dsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InjectDB(ctx context.Context) *gorm.DB {
	return di.Inject[*gorm.DB](ctx, Service)
}
