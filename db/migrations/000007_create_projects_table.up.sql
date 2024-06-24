CREATE TABLE projects (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX projects_name_unique ON projects (name);

INSERT INTO resources (name) VALUES ('projects.view'), ('projects.create'), ('projects.update'), ('projects.delete');

-- Add the new resources to the admin role
-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
	INSERT INTO role_to_resources (role_id, resource_id)
		SELECT admin_role_id.id, resources.id FROM resources, admin_role_id WHERE resources.name IN ('projects.view', 'projects.create', 'projects.update', 'projects.delete');
