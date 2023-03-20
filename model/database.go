package model

import (
	"fmt"
	"math"
	"strings"

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
	Find(id uint) (interface{}, error)
	Where(query interface{}, args ...interface{}) *model
	Save() error
	Delete(data interface{}) (interface{}, error)
	Load(model interface{}) *model
	Count() (int64, error)
	With(relation string) error
}

// model is the concrete type that implements the Model interface.
// It provides the actual implementation of the methods defined in the interface.
type model struct {
	db       *gorm.DB
	tempData interface{}
	memoryDB *gorm.DB
}

// Pagination represents pagination information.
type Pagination struct {
	Records        interface{}
	TotalRecords   int
	TotalPages     int
	CurrentPage    int
	RecordsPerPage int
}

// NewModel creates a new instance of the model type with the specified database connection.
func NewModel(db *gorm.DB) *model {
	var v interface{}
	return &model{db, v, db}
}

// Model specifies the model to be used for subsequent database operations.
func (m *model) modelLoad(model interface{}) *model {
	m.db = m.memoryDB.Model(model)
	return m
}

// Load loads data into the specified model object.
func (m *model) Load(model interface{}) *model {
	m.modelLoad(&model)
	m.tempData = model
	return m
}

// Find searches for a record with the specified ID and returns it.
func (m *model) Find(id uint) (interface{}, error) {
	err := m.db.First(m.tempData, id).Error
	return m.tempData, err
}

// Where applies the specified query to the model.
func (m *model) Where(query interface{}, args ...interface{}) *model {
	m.db = m.db.Where(query, args...)
	return m
}

// Get retrieves the records that match the current query.
func (m *model) Get() error {
	return m.db.Find(m.tempData).Error
}

// OrderBy specifies the order in which the records should be retrieved.
func (m *model) OrderBy(column string, mode string) *model {
	mode = strings.ToUpper(mode)
	order := fmt.Sprintf("%s %s", column, mode)
	m.db = m.db.Order(order)
	return m
}

// Limit specifies the maximum number of records to retrieve.
func (m *model) Limit(limit int) *model {
	m.db = m.db.Limit(limit)
	return m
}

// Delete deletes the specified record from the database.
func (m *model) Delete(data interface{}) (interface{}, error) {
	err := m.db.Delete(data).Error
	return data, err
}

// Count returns the total number of records that match the current query.
func (m *model) Count() (int64, error) {
	var count int64
	err := m.db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Paginate retrieves the records that match the current query and returns pagination information.
func (m *model) Paginate(model interface{}, page int, perPage int) (*Pagination, error) {
	var totalRecords int64
	m.db.Model(model).Count(&totalRecords)

	offset := (page - 1) * perPage
	m.db = m.db.Offset(offset).Limit(perPage)

	err := m.db.Find(model).Error
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(int(totalRecords)) / float64(perPage)))

	pagination := &Pagination{
		Records:        model,
		TotalRecords:   int(totalRecords),
		CurrentPage:    page,
		TotalPages:     totalPages,
		RecordsPerPage: perPage,
	}

	return pagination, nil
}

// Save creates a new record if the model has no ID, or updates an existing record if the model has an ID.
func (m *model) Save() error {
	err := m.db.Save(m.tempData).Error
	if err != nil {
		return err
	}
	return nil
}

// With adds an eager load for the specified relation.
// It takes in a string representing the name of the relation, and applies an eager load for that relation to the current query.
// If an error occurs during the eager load, the function returns an error object with the corresponding error message. Otherwise, it returns nil.
func (m *model) With(relation string) error {
	err := m.db.Preload(relation).Error
	if err != nil {
		return err
	}
	return nil
}
