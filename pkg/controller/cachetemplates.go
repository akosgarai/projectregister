package controller

// CacheTemplates builds the templates and stores them in templates.
func (c *Controller) CacheTemplates() {

	// Template for the login page.
	c.renderer.Template.AddTemplate("login.html", []string{c.renderer.GetTemplateDirectoryPath() + "/auth/login.html.tmpl"})

	// Template for the dashboard.
	c.renderer.Template.AddTemplate("dashboard.html", []string{c.renderer.GetTemplateDirectoryPath() + "/dashboard/index.html.tmpl"})

	// Template for application import
	c.renderer.Template.AddTemplate("application-import.html", []string{c.renderer.GetTemplateDirectoryPath() + "/application/import.html.tmpl"})
	// Template for application mapping
	c.renderer.Template.AddTemplate("application-mapping.html", []string{c.renderer.GetTemplateDirectoryPath() + "/application/mapping.html.tmpl"})
	// Template for application import results
	c.renderer.Template.AddTemplate("application-import-results.html", []string{c.renderer.GetTemplateDirectoryPath() + "/application/import-results.html.tmpl"})

	resources := []string{
		"user", "role", "client", "project", "domain", "environment",
		"runtime", "pool", "database", "server", "application",
	}

	for _, resource := range resources {
		// Template for the view.
		c.renderer.Template.AddTemplate(resource+"-view.html", []string{c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/view.html.tmpl"})
		// Template for the create.
		c.renderer.Template.AddTemplate(resource+"-create.html", []string{c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/create.html.tmpl"})
		// Template for the update.
		c.renderer.Template.AddTemplate(resource+"-update.html", []string{c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/update.html.tmpl"})
		// Template for the list.
		c.renderer.Template.AddTemplate(resource+"-list.html", []string{c.renderer.GetTemplateDirectoryPath() + "/" + resource + "/list.html.tmpl"})
	}
}
