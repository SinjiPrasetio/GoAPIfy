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
		&EmailVerification{},
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
	With(relation string) *model
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

// Load loads data into the specified model object.
// It takes in a pointer to the model object, loads it into the model, and returns a pointer to the model object.
// This method loads the specified model object into the current model, allowing it to be used for subsequent database operations.
func (m *model) Load(model interface{}) *model {
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
func (m *model) With(relation string) *model {
	m.db = m.db.Preload(relation)
	return m
}

// BeginTransaction starts a new database transaction and returns a pointer to
// the model instance, allowing for method chaining. It saves the current database
// connection state in memoryDB before starting the transaction.
func (m *model) BeginTransaction() *model {
	tx := m.db.Begin()
	m.memoryDB = m.db
	m.db = tx
	return m
}

// CommitTransaction commits the current database transaction and returns an error
// if any issues occur during the commit operation.
func (m *model) CommitTransaction() error {
	return m.db.Commit().Error
}

// RollbackTransaction rolls back the current database transaction and returns an
// error if any issues occur during the rollback operation.
func (m *model) RollbackTransaction() error {
	return m.db.Rollback().Error
}

// ApplyScope applies a given scope function to the model's query and returns a pointer
// to the model instance. The scope function should take a *gorm.DB as an input and
// return a *gorm.DB.
func (m *model) ApplyScope(scope func(*gorm.DB) *gorm.DB) *model {
	m.db = scope(m.db)
	return m
}

// GroupBy applies the GROUP BY clause to the model's query using the given column
// and returns a pointer to the model instance.
func (m *model) GroupBy(column string) *model {
	m.db = m.db.Group(column)
	return m
}

// Having applies the HAVING clause to the model's query using the given query and
// arguments, and returns a pointer to the model instance.
func (m *model) Having(query interface{}, args ...interface{}) *model {
	m.db = m.db.Having(query, args...)
	return m
}

// Distinct retrieves distinct records from the model's query using the given columns
// and returns a pointer to the model instance.
func (m *model) Distinct(columns ...string) *model {
	columnsInterface := make([]interface{}, len(columns))
	for i, col := range columns {
		columnsInterface[i] = col
	}
	m.db = m.db.Distinct(columnsInterface...)
	return m
}

// UpdateColumn updates a single column in the model's query with the given value and
// returns an error if any issues occur during the update operation. This method does
// not trigger callbacks or validations.
func (m *model) UpdateColumn(column string, value interface{}) error {
	return m.db.UpdateColumn(column, value).Error
}

// UpdateColumns updates multiple columns in the model's query with the given values map
// and returns an error if any issues occur during the update operation. This method does
// not trigger callbacks or validations.
func (m *model) UpdateColumns(values map[string]interface{}) error {
	return m.db.UpdateColumns(values).Error
}

// Max finds the maximum value of the specified column for the records that match
// the current query and stores the result in the given result interface. It returns
// an error if any issues occur during the operation.
func (m *model) Max(column string, result interface{}) error {
	return m.db.Select("MAX(?) as max", column).Scan(result).Error
}

// Min finds the minimum value of the specified column for the records that match
// the current query and stores the result in the given result interface. It returns
// an error if any issues occur during the operation.
func (m *model) Min(column string, result interface{}) error {
	return m.db.Select("MIN(?) as min", column).Scan(result).Error
}

// WithCount retrieves the count of a related model using the specified relationship
// and stores the result in the model's query. It returns a pointer to the model instance.
func (m *model) WithCount(relation string) *model {
	m.db = m.db.Joins(fmt.Sprintf("%s", relation)).Select(fmt.Sprintf("COUNT(%s.id) as %s_count", relation, relation)).Group(fmt.Sprintf("%s.id", relation))
	return m
}

// WhereHas applies the specified query to a related model using the given relationship.
// It takes in a relationship name, a query function, and returns a pointer to the model instance.
func (m *model) WhereHas(relation string, query func(*gorm.DB) *gorm.DB) *model {
	m.db = m.db.Joins(relation).Where(query)
	return m
}

// OrWhereHas applies the specified query to a related model using the given relationship,
// but adds an OR condition to the query.
// It takes in a relationship name, a query function, and returns a pointer to the model instance.
func (m *model) OrWhereHas(relation string, query func(*gorm.DB) *gorm.DB) *model {
	m.db = m.db.Joins(relation).Or(query)
	return m
}
