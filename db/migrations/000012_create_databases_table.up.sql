CREATE TABLE databases (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX databases_name_unique ON databases (name);

INSERT INTO resources (name) VALUES ('databases.view'), ('databases.create'), ('databases.update'), ('databases.delete');

-- Add the new resources to the admin role
-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
	INSERT INTO role_to_resources (role_id, resource_id)
		SELECT admin_role_id.id, resources.id FROM resources, admin_role_id WHERE resources.name IN ('databases.view', 'databases.create', 'databases.update', 'databases.delete');
