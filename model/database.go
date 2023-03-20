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
	Delete() error
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

// modelLoad specifies the model to be used for subsequent database operations.
// It takes in a pointer to the model object and returns a pointer to the model object.
// This method sets the current model to the specified model object, allowing it to be used for subsequent database operations.
func (m *model) modelLoad(model interface{}) *model {
	m.db = m.memoryDB.Model(model)
	return m
}

// Load loads data into the specified model object.
// It takes in a pointer to the model object, loads it into the model, and returns a pointer to the model object.
// This method loads the specified model object into the current model, allowing it to be used for subsequent database operations.
func (m *model) Load(model interface{}) *model {
	m.modelLoad(&model)
	m.tempData = model
	return m
}

// Find searches for a model with the specified ID and returns it.
// It takes in a uint representing the ID to search for, and returns the found data and an error, if any.
// If the data is found, the function returns the data with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Find(id uint) (interface{}, error) {
	err := m.db.First(m.tempData, id).Error
	return m.tempData, err
}

// Where applies the specified query to the model.
// It takes in a query interface and a variable number of arguments, and returns a pointer to the model object.
// This method applies the specified query to the current query on the model, with the specified arguments.
func (m *model) Where(query interface{}, args ...interface{}) *model {
	m.db = m.db.Where(query, args...)
	return m
}

// Get retrieves the records that match the current query.
// It takes no input parameters and returns an error, if any.
// If the retrieval is successful, the function returns nil. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Get() error {
	return m.db.Find(m.tempData).Error
}

// OrderBy specifies the order in which the records should be retrieved.
// It takes in two strings representing the column and the mode of the order, respectively, and returns a pointer to the model object.
// This method orders the records in the database by the specified column, with the specified mode of order (ASC or DESC).
func (m *model) OrderBy(column string, mode string) *model {
	mode = strings.ToUpper(mode)
	order := fmt.Sprintf("%s %s", column, mode)
	m.db = m.db.Order(order)
	return m
}

// Limit specifies the maximum number of records to retrieve.
// It takes in an integer representing the maximum number of records to retrieve, and returns a pointer to the model object.
// This method limits the number of records that can be retrieved from the database to the specified limit.
func (m *model) Limit(limit int) *model {
	m.db = m.db.Limit(limit)
	return m
}

// Delete : deletes the specified record from the database.
// It takes no input parameters and returns an error, if any.
// If the deletion is successful, the function returns nil. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Delete() error {
	err := m.db.Delete(m.tempData).Error
	return err
}

// Count returns the total number of records that match the current query.
// It takes no input parameters and returns the total number of records as an int64 and an error, if any.
// If the counting is successful, the function returns the total number of records with a nil error. If an error occurs, it returns an error object with the corresponding error message.
func (m *model) Count() (int64, error) {
	var count int64
	err := m.db.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Get retrieves the records that match the current query.
// It takes no input parameters and returns an error, if any.
// If the retrieval is successful, the function returns nil. If an error occurs, it returns an error object with the corresponding error message.
// It takes in a pointer to the model object, the page number, and the number of records per page, and returns a pointer to a Pagination object and an error, if any.
// If the pagination is successful, the function returns a pointer to the Pagination object with a nil error. If an error occurs, it returns an error object with the corresponding error message.
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
// It takes no input parameters and returns an error, if any.
// If the creation or update is successful, the function returns nil. If an error occurs, it returns an error object with the corresponding error message.
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

// The FindByIDPreload function retrieves a record from a database table by its ID, and preloads any specified associations using the GORM library in Go.
// The function takes in the ID of the record, a pointer to a struct that represents the database table,
// and an optional list of associations to preload. The function returns the populated struct and an error.
func (m *model) FindByIDPreload(id uint, model interface{}, preloads ...string) (interface{}, error) {
	query := m.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	err := query.First(model, id).Error
	if err != nil {
		return model, err
	}
	return model, nil
}
