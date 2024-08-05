package parser

import (
	"testing"
)

// TestNewApplicationImportRow tests the NewApplicationImportRow function.
func TestNewApplicationImportRow(t *testing.T) {
	row := NewApplicationImportRow()
	if row.ErrorMessage != "" {
		t.Errorf("The error message is not empty.")
	}
	if len(row.RowData) != 0 {
		t.Errorf("The row data is not empty.")
	}
	if row.Application != nil {
		t.Errorf("The application is not nil.")
	}
}

// TestNewApplicationImportResult tests the NewApplicationImportResult function.
func TestNewApplicationImportResult(t *testing.T) {
	result := NewApplicationImportResult()
	if len(*result) != 0 {
		t.Errorf("The result is not empty.")
	}
}

// TestNewMappingRule tests the NewMappingRule function.
func TestNewMappingRule(t *testing.T) {
	rule := NewMappingRule()
	if rule.ColumnIndex != -1 {
		t.Errorf("The column index is not -1.")
	}
	if rule.CustomValue != "" {
		t.Errorf("The custom value is not empty.")
	}
}

// TestNewApplicationImportMapping tests the NewApplicationImportMapping function.
func TestNewApplicationImportMapping(t *testing.T) {
	mapping := NewApplicationImportMapping()
	// every mapping should have a default rule.
	// The mapping has to contain default rules for the following fields:
	// client, project, runtime, pool, domains, framework, database, database_name, database_user, doc_root, repository, branch
	fields := []string{"client", "project", "runtime", "pool", "domains", "framework", "database", "database_name", "database_user", "doc_root", "repository", "branch"}
	for _, field := range fields {
		if _, ok := (*mapping)[field]; !ok {
			t.Errorf("The mapping does not contain the field: %s", field)
		}
	}

	// The mapping should not contain any other fields.
	if len(*mapping) != len(fields) {
		t.Errorf("The mapping contains other fields.")
	}
}
