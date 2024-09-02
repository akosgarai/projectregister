CREATE TABLE frameworks (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	score INTEGER DEFAULT 0,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX frameworks_name_unique ON frameworks (name);

INSERT INTO resources (name) VALUES ('frameworks.view'), ('frameworks.create'), ('frameworks.update'), ('frameworks.delete');

-- Add the new resources to the admin role
-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
	INSERT INTO role_to_resources (role_id, resource_id)
		SELECT admin_role_id.id, resources.id FROM resources, admin_role_id WHERE resources.name IN ('frameworks.view', 'frameworks.create', 'frameworks.update', 'frameworks.delete');

-- Insert the current frameworks into the table
-- It could be extracted from the applications table, from the framework column
INSERT INTO frameworks (name)
	SELECT DISTINCT framework FROM applications WHERE framework IS NOT NULL;

-- Add the new column to the applications table
ALTER TABLE applications ADD COLUMN framework_id INTEGER;
ALTER TABLE applications ADD FOREIGN KEY (framework_id) REFERENCES frameworks (id);
-- Update the applications table with the new framework_id
UPDATE applications SET framework_id = frameworks.id FROM frameworks WHERE applications.framework = frameworks.name;
-- Drop the old column
ALTER TABLE applications DROP COLUMN framework;
