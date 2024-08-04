package controller

// CacheTemplates builds the templates and stores them in templates.
func (c *Controller) CacheTemplates() {

	headerTemplate := c.renderer.GetTemplateDirectoryPath() + "/frontend-components/header.html.tmpl"
	formItemsTemplate := c.renderer.GetTemplateDirectoryPath() + "/frontend-components/form-items.html.tmpl"
	detailItemsTemplate := c.renderer.GetTemplateDirectoryPath() + "/frontend-components/detail-items.html.tmpl"
	listingItemsTemplate := c.renderer.GetTemplateDirectoryPath() + "/frontend-components/listing.html.tmpl"

	// Template for the login page.
	c.renderer.Template.AddTemplate("login.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/auth/login.html.tmpl"})

	// Template for the dashboard.
	c.renderer.Template.AddTemplate("dashboard.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/dashboard/index.html.tmpl"})

	// Template for application import
	c.renderer.Template.AddTemplate("application-import.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/application/import.html.tmpl"})
	// Template for application mapping
	c.renderer.Template.AddTemplate("application-mapping.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/application/mapping.html.tmpl"})
	// Template for application import results
	c.renderer.Template.AddTemplate("application-import-results.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/application/import-results.html.tmpl"})

	// Template for the view.
	c.renderer.Template.AddTemplate("detail-page.html", []string{headerTemplate, detailItemsTemplate})
	// Template for the list.
	c.renderer.Template.AddTemplate("listing-page.html", []string{headerTemplate, listingItemsTemplate})
	// Template for the update.
	c.renderer.Template.AddTemplate("form-page.html", []string{headerTemplate, formItemsTemplate})
}
