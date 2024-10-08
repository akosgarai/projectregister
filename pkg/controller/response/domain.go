package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/controller/response/components"
	"github.com/akosgarai/projectregister/pkg/model"
)

// NewDomainDetailResponse is a constructor for the DetailResponse struct for a domain.
func NewDomainDetailResponse(currentUser *model.User, domain *model.Domain) *DetailResponse {
	headerText := "Domain Detail"
	headerContent := components.NewContentHeader(headerText, newDetailHeaderButtons(currentUser, "domains", fmt.Sprintf("%d", domain.ID)))
	if currentUser.HasPrivilege("domains.update") {
		headerContent.Buttons = append(
			[]*components.Link{components.NewLink("Check", fmt.Sprintf("/admin/domain/check-ssl/%d", domain.ID))},
			headerContent.Buttons...,
		)
	}
	protocol := "http"
	if domain.HasSSL {
		protocol = "https"
	}
	details := &components.DetailItems{
		{Label: "ID", Value: &components.DetailValues{{Value: fmt.Sprintf("%d", domain.ID)}}},
		{Label: "Name", Value: &components.DetailValues{{Value: domain.Name, Link: fmt.Sprintf("%s://%s", protocol, domain.Name)}}},
		{Label: "Has SSL", Value: &components.DetailValues{{Value: fmt.Sprintf("%t", domain.HasSSL)}}},
		{Label: "Created At", Value: &components.DetailValues{{Value: domain.CreatedAt}}},
		{Label: "Updated At", Value: &components.DetailValues{{Value: domain.UpdatedAt}}},
	}
	return NewDetailResponse(headerText, currentUser, headerContent, details)
}

// NewCreateDomainResponse is a constructor for the FormResponse struct for a domain.
func NewCreateDomainResponse(currentUser *model.User) *FormResponse {
	return newDomainFormResponse("Create Domain", currentUser, &model.Domain{}, "/admin/domain/create", "POST", "Create")
}

// NewUpdateDomainResponse is a constructor for the FormResponse struct for a domain.
func NewUpdateDomainResponse(currentUser *model.User, domain *model.Domain) *FormResponse {
	return newDomainFormResponse("Update Domain", currentUser, domain, fmt.Sprintf("/admin/domain/update/%d", domain.ID), "POST", "Update")
}

// newDomainFormResponse is a constructor for the FormResponse struct for a domain.
func newDomainFormResponse(title string, currentUser *model.User, domain *model.Domain, action, method, submitLabel string) *FormResponse {
	headerContent := components.NewContentHeader(title, []*components.Link{components.NewLink("List", "/admin/domain/list")})
	formItems := []*components.FormItem{
		// Name.
		components.NewFormItem("Name", "name", "text", domain.Name, true, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: action,
		Method: method,
		Submit: submitLabel,
	}
	return NewFormResponse(title, currentUser, headerContent, form)
}

// NewDomainListResponse is a constructor for the ListingResponse struct of the domains.
func NewDomainListResponse(currentUser *model.User, domains *model.Domains, filter *model.DomainFilter) *ListingResponse {
	headerText := "Domain List"
	headerContent := components.NewContentHeader(headerText, []*components.Link{})
	if currentUser.HasPrivilege("domains.create") {
		headerContent.Buttons = append(headerContent.Buttons, components.NewLink("Create", "/admin/domain/create"))
	}
	listingHeader := &components.ListingHeader{
		Headers: []string{"ID", "Name", "Has SSL", "Actions"},
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
		hasSSLColumn := &components.ListingColumn{&components.ListingColumnValues{{Value: fmt.Sprintf("%t", domain.HasSSL)}}}
		columns = append(columns, hasSSLColumn)
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
	/* Create the search form. The only form item is the name. */
	formItems := []*components.FormItem{
		components.NewFormItem("Name", "name", "text", filter.Name, false, nil, nil),
	}
	form := &components.Form{
		Items:  formItems,
		Action: "/admin/domain/list",
		Method: "POST",
		Submit: "Search",
	}
	return NewListingResponse(headerText, currentUser, headerContent, &components.Listing{Header: listingHeader, Rows: &listingRows}, form)
}
