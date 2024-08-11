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
	if len(result) != 0 {
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
		if _, ok := (mapping)[field]; !ok {
			t.Errorf("The mapping does not contain the field: %s", field)
		}
	}

	// The mapping should not contain any other fields.
	if len(mapping) != len(fields) {
		t.Errorf("The mapping contains other fields.")
	}
}

// TestMapRowDefaultMapping tests the MapRow function with default data.
func TestMapRowDefaultMapping(t *testing.T) {
	mapping := NewApplicationImportMapping()
	importRow := mapping.MapRow([]string{"client", "project", "runtime", "pool", "domains", "framework", "database", "database_name", "database_user", "doc_root", "repository", "branch"})
	// The import row has to be empty.
	if importRow.Application != nil {
		t.Errorf("The application is not nil.")
	}
	if importRow.ErrorMessage != "" {
		t.Errorf("The error message is not empty.")
	}
	if len(importRow.RowData) != len(mapping) {
		t.Errorf("The row data length is not equal to the mapping length.")
	}
	for key, value := range importRow.RowData {
		if value != "" {
			t.Errorf("The value of the key %s is not empty.", key)
		}
	}
}

// TestMapRowWithMapping tests the MapRow function with mapping data.
func TestMapRowWithMapping(t *testing.T) {
	mapping := NewApplicationImportMapping()
	mapping["client"].ColumnIndex = 0
	mapping["project"].ColumnIndex = 1
	mapping["runtime"].CustomValue = "php"
	mapping["pool"].ColumnIndex = 3
	mapping["domains"].ColumnIndex = 4
	mapping["framework"].ColumnIndex = 5
	mapping["database"].ColumnIndex = 6
	mapping["database_name"].ColumnIndex = 7
	mapping["database_user"].ColumnIndex = 8
	mapping["doc_root"].ColumnIndex = 9
	mapping["repository"].ColumnIndex = 10
	mapping["branch"].ColumnIndex = 11
	importRow := mapping.MapRow([]string{"client", "project", "runtime", "pool", "domains", "framework", "database", "database_name", "database_user", "doc_root", "repository", "branch"})
	// The import row has to be filled with the mapping data.
	if importRow.Application != nil {
		t.Errorf("The application is not nil.")
	}
	if importRow.ErrorMessage != "" {
		t.Errorf("The error message is not empty.")
	}
	if len(importRow.RowData) != len(mapping) {
		t.Errorf("The row data length is not equal to the mapping length.")
	}
	if importRow.RowData["client"] != "client" {
		t.Errorf("The client value is not correct.")
	}
	if importRow.RowData["project"] != "project" {
		t.Errorf("The project value is not correct.")
	}
	if importRow.RowData["runtime"] != "php" {
		t.Errorf("The runtime value is not correct.")
	}
	if importRow.RowData["pool"] != "pool" {
		t.Errorf("The pool value is not correct.")
	}
	if importRow.RowData["domains"] != "domains" {
		t.Errorf("The domains value is not correct.")
	}
	if importRow.RowData["framework"] != "framework" {
		t.Errorf("The framework value is not correct.")
	}
	if importRow.RowData["database"] != "database" {
		t.Errorf("The database value is not correct.")
	}
	if importRow.RowData["database_name"] != "database_name" {
		t.Errorf("The database_name value is not correct.")
	}
	if importRow.RowData["database_user"] != "database_user" {
		t.Errorf("The database_user value is not correct.")
	}
	if importRow.RowData["doc_root"] != "doc_root" {
		t.Errorf("The doc_root value is not correct.")
	}
	if importRow.RowData["repository"] != "repository" {
		t.Errorf("The repository value is not correct.")
	}
	if importRow.RowData["branch"] != "branch" {
		t.Errorf("The branch value is not correct.")
	}
}
