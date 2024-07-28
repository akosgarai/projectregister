package controller

import (
	"github.com/akosgarai/projectregister/pkg/resources"
)

// CacheTemplates builds the templates and stores them in templates.
func (c *Controller) CacheTemplates() {

	headerTemplate := c.renderer.GetTemplateDirectoryPath() + "/frontend-components/header.html.tmpl"

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

	for _, resource := range resources.Resources {
		// Template for the view.
		c.renderer.Template.AddTemplate(resource+"-view.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/view.html.tmpl"})
		// Template for the create.
		c.renderer.Template.AddTemplate(resource+"-create.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/create.html.tmpl"})
		// Template for the update.
		c.renderer.Template.AddTemplate(resource+"-update.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/update.html.tmpl"})
		// Template for the list.
		c.renderer.Template.AddTemplate(resource+"-list.html", []string{headerTemplate, c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/list.html.tmpl"})
	}
}
