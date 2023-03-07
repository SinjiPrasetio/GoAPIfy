package model

import (
	"fmt"

	"gorm.io/gorm"
)

// Model is the interface that must be implemented by models
type Model interface {
	Create(data interface{}) (interface{}, error)
	FindByID(id uint) (interface{}, error)
	FindByKey(key string, value interface{}) (interface{}, error)
}

// Model is the concrete type that implements the Model interface
type model struct {
	db *gorm.DB
}

// NewModel creates a new instance of the model type with the specified database connection
func NewModel(db *gorm.DB) *model {
	return &model{db}
}

// AutoMigration automatically migrates the models to the database schema
func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
	)
	return err
}

// Create creates a new model with the specified data
func (m *model) Create(data interface{}) (interface{}, error) {
	err := m.db.Create(data).Error
	return data, err
}

// FindByID searches for a model with the specified ID and returns it
func (m *model) FindByID(id uint) (interface{}, error) {
	var model interface{}
	err := m.db.First(model, id).Error
	return model, err
}

// FindByKey searches for a model with the specified key/value pair and returns it
func (m *model) FindByKey(key string, value interface{}) (interface{}, error) {
	var model interface{}
	err := m.db.Where(fmt.Sprintf("%s = ?", key), value).Find(&model).Error
	return model, err
}
