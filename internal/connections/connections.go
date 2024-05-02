package connections

import (
	"github.com/tadilbek11kz/ePharma-backend/internal/app/config"

	pharmacy "github.com/tadilbek11kz/ePharma-backend/pkg/pharmacy"
	product "github.com/tadilbek11kz/ePharma-backend/pkg/product"
	user "github.com/tadilbek11kz/ePharma-backend/pkg/user"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connections struct {
	Postgres *gorm.DB
}

func (c *Connections) Close() {
	if c.Postgres != nil {
		db, _ := c.Postgres.DB()
		db.Close()
	}
}

func New(cfg *config.TomlConfig) (*Connections, error) {
	pg, err := connectPostgres(cfg.DatabaseUrl)
	if err != nil {
		return nil, err
	}

	conns := &Connections{
		Postgres: pg,
	}

	return conns, nil
}

func connectPostgres(connectUrl string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	err = sqlDB.Ping()
	if err != nil {
		println(err.Error())
		panic(err)
	}

	db.AutoMigrate(
		&user.User{},
		&pharmacy.Pharmacy{},
		&product.Product{},
	)

	runMigrations(db)

	return db, nil
}
