package parser

import (
	"github.com/akosgarai/projectregister/pkg/model"
)

// ApplicationImportRow is the struct for the application import row results.
// It contains
// - the error message - if there is any
// - the row data - the original row data
// - the application - the imported application
type ApplicationImportRow struct {
	ErrorMessage string
	RowData      map[string]string
	Application  *model.Application
}

// NewApplicationImportRow is a constructor for the ApplicationImportRow struct.
// It returns a new ApplicationImportRow instance with default (empty) values.
func NewApplicationImportRow() *ApplicationImportRow {
	return &ApplicationImportRow{
		ErrorMessage: "",
		RowData:      map[string]string{},
		Application:  nil,
	}
}

// ApplicationImportResult is the struct for the application import results.
// It contains the application import rows.
// The key is the row number, the value is the ApplicationImportRow.
type ApplicationImportResult map[int]*ApplicationImportRow

// NewApplicationImportResult is a constructor for the ApplicationImportResult struct.
// It returns a new ApplicationImportResult instance with default (empty) values.
func NewApplicationImportResult() ApplicationImportResult {
	return ApplicationImportResult{}
}

// MappingRule is the struct for the mapping rule.
// It contains
// - the column index
// - the custom value On case of the column index is -1, the custom value will be used.
type MappingRule struct {
	ColumnIndex int
	CustomValue string
}

// NewMappingRule is a constructor for the MappingRule struct.
// It returns a new MappingRule instance with its default values.
func NewMappingRule() *MappingRule {
	return &MappingRule{
		ColumnIndex: -1,
		CustomValue: "",
	}
}

// ApplicationImportMapping is the struct for the application import mapping.
// It contains the mapping of the column names and the column indexes.
// The key is the column name, the value is the mapping rule.
type ApplicationImportMapping map[string]*MappingRule

// NewApplicationImportMapping is a constructor for the ApplicationImportMapping struct.
// It returns a new ApplicationImportMapping instance with default (empty) values.
// Maps the column names to the -1 index.
func NewApplicationImportMapping() ApplicationImportMapping {
	return ApplicationImportMapping{
		"client":        NewMappingRule(),
		"project":       NewMappingRule(),
		"runtime":       NewMappingRule(),
		"pool":          NewMappingRule(),
		"domains":       NewMappingRule(),
		"framework":     NewMappingRule(),
		"database":      NewMappingRule(),
		"database_name": NewMappingRule(),
		"database_user": NewMappingRule(),
		"doc_root":      NewMappingRule(),
		"repository":    NewMappingRule(),
		"branch":        NewMappingRule(),
	}
}

// MapRow maps the row data to the application import row.
// It uses the mapping to set the values.
func (m ApplicationImportMapping) MapRow(row []string) *ApplicationImportRow {
	importRow := NewApplicationImportRow()
	for key, rule := range m {
		if rule.ColumnIndex == -1 {
			importRow.RowData[key] = rule.CustomValue
		} else {
			importRow.RowData[key] = row[rule.ColumnIndex]
		}
	}
	return importRow
}
