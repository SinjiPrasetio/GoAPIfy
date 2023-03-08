package model

import (
	"fmt"

	"gorm.io/gorm"
)

// AutoMigration automatically migrates the models to the database schema.
// It takes in a pointer to the GORM database object, and returns an error if one occurs.
// New models should be added here to ensure their corresponding database tables are created or updated during migration.
func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
	)
	return err
}

// Model is the interface that must be implemented by models.
// It defines the common methods that will be used across different models.
type Model interface {
	Create(data interface{}) (interface{}, error)
	FindByID(id uint, model interface{}) (interface{}, error)
	FindByKey(key string, value string, model interface{}) (interface{}, error)
	Update(data interface{}) (interface{}, error)
	Delete(data interface{}) (interface{}, error)
}

// model is the concrete type that implements the Model interface.
// It provides the actual implementation of the methods defined in the interface.
type model struct {
	db *gorm.DB
}

// NewModel creates a new instance of the model type with the specified database connection.
func NewModel(db *gorm.DB) *model {
	return &model{db}
}

// Create creates a new model with the specified data.
// It takes in a pointer to the model object and the data to be created, and returns the created data and an error, if any.
// If the creation is successful, the function returns the created data with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Create(data interface{}) (interface{}, error) {
	err := m.db.Create(data).Error
	return data, err
}

// FindByID searches for a model with the specified ID and returns it.
// It takes in a pointer to the model object, the ID to search for, and returns the found data and an error, if any.
// If the data is found, the function returns the data with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) FindByID(id uint, model interface{}) (interface{}, error) {
	err := m.db.First(model, id).Error
	return model, err
}

// FindByKey searches for a model with the specified key/value pair and returns it.
// It takes in a pointer to the model object, the key and value to search for, and a pointer to the model where the found data will be stored.
// The function returns the found data and an error, if any. If the data is found, it is stored in the model and returned with a nil error.
// If an error occurs, it returns an error object with the corresponding error message.
func (m *model) FindByKey(key string, value string, model interface{}) (interface{}, error) {
	err := m.db.Where(fmt.Sprintf("%s = ?", key), value).Find(model).Error
	return model, err
}

// Update updates an existing record in the database by saving the provided data to the model's database object.
// It takes in a pointer to the model object and the data to be updated, and returns the updated data and an error, if any.
// If the update is successful, the function returns the updated data with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Update(data interface{}) (interface{}, error) {
	err := m.db.Save(data).Error
	return data, err
}

// Delete deletes the specified record from the database.
// It takes in a pointer to the model object and the data to be deleted, and returns the deleted data and an error, if any.
// If the deletion is successful, the function returns the deleted data with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Delete(data interface{}) (interface{}, error) {
	err := m.db.Delete(data).Error
	return data, err
}
