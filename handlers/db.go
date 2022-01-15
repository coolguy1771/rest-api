package handlers

import (
	"github.com/coolguy1771/rest-api/models"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ClientInterface resembles a db interface to interact with an underlying db
type ClientInterface interface {
	Ping() error
	Connect(connectionString string) error
	GetUserByID(id int) *models.User
	SetUser(user *models.User) error
	GetUnitByID(id int) *models.Unit
	SetUnit(unit *models.Unit) error
}

// Client is a custom db client
type Client struct {
	Client *gorm.DB
}

// Ping allows the db to be pinged.
func (c *Client) Ping() error {
	return c.Client.DB().Ping()
}

// Connect establishes a connection to the database and auto migrates the database schema
func (c *Client) Connect(connectionString string) error {
	var err error
	// Create the database connection
	c.Client, err = gorm.Open("sqlite", connectionString)

	// End the program with an error if it could not connect to the database
	if err != nil {
		return err
	}
	c.Client.LogMode(false)
	c.autoMigrate()
	return nil
}

// autoMigrate creates the default database schema
func (c *Client) autoMigrate() {
	c.Client.AutoMigrate(&models.User{})
	c.Client.AutoMigrate(&models.Unit{})
}

// GetUserByID queries an user from the database
func (c *Client) GetUserByID(id int) *models.User {
	user := &models.User{}

	c.Client.Where("id = ?", id).First(&user).Scan(user)

	return user
}

// SetUser writes an user to the database
func (c *Client) SetUser(user *models.User) error {
	// Upsert by trying to create and updating on conflict
	if err := c.Client.Create(&user).Error; err != nil {
		return c.Client.Model(&user).Where("id = ?", user.ID).Update(&user).Error
	}
	return nil
}

// GetUnitByID queries an unit from the database
func (c *Client) GetUnitByID(id int) *models.Unit {
	unit := &models.Unit{}

	c.Client.Where("id = ?", id).First(&unit).Scan(unit)

	return unit
}

// SetUnit writes an unit to the database
func (c *Client) SetUnit(unit *models.Unit) error {
	// Upsert by trying to create and updating on conflict
	if err := c.Client.Create(&unit).Error; err != nil {
		return c.Client.Model(&unit).Where("id = ?", unit.ID).Update(&unit).Error
	}
	return nil
}
