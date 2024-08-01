package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// DomainResponse is the struct for the domain page.
type DomainResponse struct {
	*Response
	Domain *model.Domain
}

// DomainDetailResponse is the struct for the domain detail page.
type DomainDetailResponse struct {
	*DomainResponse
	Details *DetailItems
}

// NewDomainDetailResponse is a constructor for the DomainDetailResponse struct.
func NewDomainDetailResponse(currentUser *model.User, domain *model.Domain) *DomainDetailResponse {
	header := &HeaderBlock{
		Title:       "Domain Detail",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Edit",
				Link:      fmt.Sprintf("/admin/domain/update/%d", domain.ID),
				Privilege: "domains.update",
			},
			{
				Label:     "Delete",
				Link:      fmt.Sprintf("/admin/domain/delete/%d", domain.ID),
				Privilege: "domains.delete",
			},
			{
				Label:     "List",
				Link:      "/admin/domain/list",
				Privilege: "domains.view",
			},
		},
	}
	details := &DetailItems{
		{Label: "ID", Value: &DetailValues{{Value: fmt.Sprintf("%d", domain.ID)}}},
		{Label: "Name", Value: &DetailValues{{Value: domain.Name, Link: domain.Name}}},
		{Label: "Created At", Value: &DetailValues{{Value: domain.CreatedAt}}},
		{Label: "Updated At", Value: &DetailValues{{Value: domain.UpdatedAt}}},
	}
	return &DomainDetailResponse{
		DomainResponse: &DomainResponse{
			Response: NewResponse("Domain Detail", currentUser, header),
			Domain:   domain,
		},
		Details: details,
	}
}

// DomainFormResponse is the struct for the domain form responses.
type DomainFormResponse struct {
	*DomainDetailResponse
	FormItems []*FormItem
}

// NewDomainFormResponse is a constructor for the DomainFormResponse struct.
func NewDomainFormResponse(title string, currentUser *model.User, domain *model.Domain) *DomainFormResponse {
	domainDetailResponse := NewDomainDetailResponse(currentUser, domain)
	domainDetailResponse.Header.Title = title
	domainDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	domainDetailResponse.Header.Buttons = []*ActionButton{{Label: "Back", Link: "/admin/domain/list", Privilege: "domains.view"}}
	formItems := []*FormItem{
		// Name.
		{Label: "Name", Type: "text", Name: "name", Value: domain.Name, Required: true},
	}
	return &DomainFormResponse{
		DomainDetailResponse: domainDetailResponse,
		FormItems:            formItems,
	}
}

// DomainListResponse is the struct for the domain list page.
type DomainListResponse struct {
	*Response
	Listing *Listing
}

// NewDomainListResponse is a constructor for the DomainListResponse struct.
func NewDomainListResponse(currentUser *model.User, domains *model.Domains) *DomainListResponse {
	header := &HeaderBlock{
		Title:       "Domain List",
		CurrentUser: currentUser,
		Buttons: []*ActionButton{
			{
				Label:     "Create",
				Link:      "/admin/domain/create",
				Privilege: "domains.create",
			},
		},
	}
	listingHeader := &ListingHeader{
		Headers: []string{"ID", "Name", "Actions"},
	}
	// create the rows
	listingRows := ListingRows{}
	userCanEdit := currentUser.HasPrivilege("domains.update")
	userCanDelete := currentUser.HasPrivilege("domains.delete")
	for _, domain := range *domains {
		columns := ListingColumns{}
		idColumn := &ListingColumn{&ListingColumnValues{{Value: fmt.Sprintf("%d", domain.ID)}}}
		columns = append(columns, idColumn)
		nameColumn := &ListingColumn{&ListingColumnValues{{Value: domain.Name}}}
		columns = append(columns, nameColumn)
		actionsColumn := ListingColumn{&ListingColumnValues{
			{Value: "View", Link: fmt.Sprintf("/admin/domain/view/%d", domain.ID)},
		}}
		if userCanEdit {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Update", Link: fmt.Sprintf("/admin/domain/update/%d", domain.ID)})
		}
		if userCanDelete {
			*actionsColumn.Values = append(*actionsColumn.Values, &ListingColumnValue{Value: "Delete", Link: fmt.Sprintf("/admin/domain/delete/%d", domain.ID), Form: true})
		}
		columns = append(columns, &actionsColumn)

		listingRows = append(listingRows, &ListingRow{Columns: &columns})
	}
	return &DomainListResponse{
		Response: NewResponse("Domain List", currentUser, header),
		Listing:  &Listing{Header: listingHeader, Rows: &listingRows},
	}
}
