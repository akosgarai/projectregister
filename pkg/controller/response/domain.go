package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewDomainDetailResponse is a constructor for the DetailResponse struct for a domain.
func NewDomainDetailResponse(currentUser *model.User, domain *model.Domain) *DetailResponse {
	headerText := "Domain Detail"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("domains.update") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Edit", fmt.Sprintf("/admin/domain/update/%d", domain.ID)))
	}
	if currentUser.HasPrivilege("domains.delete") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Delete", fmt.Sprintf("/admin/domain/delete/%d", domain.ID)))
	}
	if currentUser.HasPrivilege("domains.view") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("List", "/admin/domain/list"))
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", domain.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: domain.Name, Link: domain.Name}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: domain.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: domain.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// DomainFormResponse is the struct for the domain form responses.
type DomainFormResponse struct {
	*DetailResponse
	FormItems []*FormItem
}

// NewDomainFormResponse is a constructor for the DomainFormResponse struct.
func NewDomainFormResponse(title string, currentUser *model.User, domain *model.Domain) *DomainFormResponse {
	domainDetailResponse := NewDomainDetailResponse(currentUser, domain)
	domainDetailResponse.Header.Title = title
	domainDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	domainDetailResponse.Header.Buttons = []*components.Link{components.NewLink("List", "/admin/domain/list")}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: domain.Name, Required: true},
	}
	return &DomainFormResponse{
		DetailResponse: domainDetailResponse,
		FormItems:      formItems,
	}
}

// NewDomainListResponse is a constructor for the ListingResponse struct of the domains.
func NewDomainListResponse(currentUser *model.User, domains *model.Domains) *ListingResponse {
	headerText := "Domain List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("domains.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/domain/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := components.ListingRows{}
	userCanEdit := currentUser.HasPrivilege("domains.update")
	userCanDelete := currentUser.HasPrivilege("domains.delete")
	for _, domain := range *domains {
		columns := components.ListingColumns{}
		idColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%d", domain.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: domain.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := components.ListingColumn{&components.ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/domain/update/%d", domain.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &components.ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/domain/delete/%d", domain.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &components.ListingRow{Columns: &columns})
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows})
}
