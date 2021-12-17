// this package contains all database operations

package db

import (
	"github.com/jinzhu/gorm"
	// postgres blank import for gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/coolguy1771/rest-api/pkg/types"
)

const (
	pageSize = 10
)

// ClientInterface resembles a db interface to interact with an underlying db
type ClientInterface interface {
	Ping() error
	Connect(connectionString string) error
	GetUserByID(id int) *types.User
	SetUser(user *types.User) error
	GetUsers(pageID int) *types.UserList
	GetUnitByID(id int) *types.Unit
	SetUnit(unit *types.Unit) error
	GetUnits(pageID int) *types.UnitList
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
	c.Client, err = gorm.Open(
		"postgres",
		connectionString,
	)

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
	c.Client.AutoMigrate(&types.User{})
	c.Client.AutoMigrate(&types.Unit{})
}

// GetUserByID queries an user from the database
func (c *Client) GetUserByID(id int) *types.User {
	user := &types.User{}

	c.Client.Where("id = ?", id).First(&user).Scan(user)

	return user
}

// SetUser writes an user to the database
func (c *Client) SetUser(user *types.User) error {
	// Upsert by trying to create and updating on conflict
	if err := c.Client.Create(&user).Error; err != nil {
		return c.Client.Model(&user).Where("id = ?", user.ID).Update(&user).Error
	}
	return nil
}

// GetUsers returns all users from the database
func (c *Client) GetUsers(pageID int) *types.UserList {
	users := &types.UserList{}
	c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&users.Items)
	if len(users.Items) == pageSize+1 {
		users.NextPageID = users.Items[len(users.Items)-1].ID
		users.Items = users.Items[:pageSize]
	}
	return users
}

// GetUnitByID queries an unit from the database
func (c *Client) GetUnitByID(id int) *types.Unit {
	unit := &types.Unit{}

	c.Client.Where("id = ?", id).First(&unit).Scan(unit)

	return unit
}

// SetUnit writes an unit to the database
func (c *Client) SetUnit(unit *types.Unit) error {
	// Upsert by trying to create and updating on conflict
	if err := c.Client.Create(&unit).Error; err != nil {
		return c.Client.Model(&unit).Where("id = ?", unit.ID).Update(&unit).Error
	}
	return nil
}

// GetUnit returns all units from the database
func (c *Client) GetUnits(pageID int) *types.UnitList {
	unit := &types.UnitList{}
	c.Client.Find(&unit.Units).Where("id >= ?", pageID).Order("id").Limit(pageSize + 1)
	if len(unit.Units) == pageSize+1 {
		unit.NextPageID = unit.Units[len(unit.Units)-1].ID
		unit.Units = unit.Units[:pageSize+1]
	}
	return unit
}
