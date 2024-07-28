package response

import (
	"fmt"

	"github.com/akosgarai/projectregister/pkg/model"
)

// DomainDetailResponse is the struct for the domain detail page.
type DomainDetailResponse struct {
	*Response
	Header *HeaderBlock
	Domain *model.Domain
}

// NewDomainDetailResponse is a constructor for the DomainDetailResponse struct.
func NewDomainDetailResponse(currentUser *model.User, domain *model.Domain) *DomainDetailResponse {
	return &DomainDetailResponse{
		Response: NewResponse("Domain Detail", currentUser),
		Header: &HeaderBlock{
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
			},
		},
		Domain: domain,
	}
}

// DomainFormResponse is the struct for the domain form responses.
type DomainFormResponse struct {
	*DomainDetailResponse
}

// NewDomainFormResponse is a constructor for the DomainFormResponse struct.
func NewDomainFormResponse(title string, currentUser *model.User, domain *model.Domain) *DomainFormResponse {
	domainDetailResponse := NewDomainDetailResponse(currentUser, domain)
	domainDetailResponse.Header.Title = title
	domainDetailResponse.Title = title
	// The buttons are unnecessary on the form page.
	domainDetailResponse.Header.Buttons = []*ActionButton{}
	return &DomainFormResponse{
		DomainDetailResponse: domainDetailResponse,
	}
}

// DomainListResponse is the struct for the domain list page.
type DomainListResponse struct {
	*Response
	Header  *HeaderBlock
	Domains []*model.Domain
}

// NewDomainListResponse is a constructor for the DomainListResponse struct.
func NewDomainListResponse(currentUser *model.User, domains []*model.Domain) *DomainListResponse {
	return &DomainListResponse{
		Response: NewResponse("Domain List", currentUser),
		Header: &HeaderBlock{
			Title:       "Domain List",
			CurrentUser: currentUser,
			Buttons: []*ActionButton{
				{
					Label:     "Create",
					Link:      "/admin/domain/create",
					Privilege: "domains.create",
				},
			},
		},
		Domains: domains,
	}
}
