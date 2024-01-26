package database

import (
	"github.com/kilianp07/lifegame/config"
	"github.com/kilianp07/lifegame/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

// NewController creates a new instance of the Controller struct with the provided configuration.
// It initializes a connection to the database using the specified DBPath.
// Returns a pointer to the Controller and an error if any occurred during the initialization.
func NewController(conf config.DBConf) (*Controller, error) {
	var (
		db  *gorm.DB
		err error
		c   = &Controller{}
	)

	if db, err = gorm.Open(sqlite.Open(conf.DBPath), &gorm.Config{}); err != nil {
		return nil, err
	}

	c.DB = db

	return c, nil
}

func (c *Controller) Init() error {
	return c.migrate()
}

// migrate is a method that performs database migration for the Controller.
// It uses the AutoMigrate function from the DB package to automatically create or update the User table in the database.
// Returns an error if the migration fails.
func (c *Controller) migrate() error {
	return c.DB.AutoMigrate(&model.User{})
}

// CreateUser creates a new user in the database.
// It takes a pointer to a User struct as a parameter and returns an error if any.
func (c *Controller) CreateUser(user *model.User) error {
	return c.DB.Create(user).Error
}

// Close closes the database connection.
// It returns an error if there was a problem closing the connection.
func (c *Controller) Close() error {
	db, err := c.DB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
