CREATE TABLE applications (
	id SERIAL PRIMARY KEY,

	client_id INT NOT NULL,
	FOREIGN KEY (client_id) REFERENCES clients (id),

	project_id INT NOT NULL,
	FOREIGN KEY (project_id) REFERENCES projects (id),

	env_id INT NOT NULL,
	FOREIGN KEY (env_id) REFERENCES environments (id),

	database_id INT NOT NULL,
	FOREIGN KEY (database_id) REFERENCES databases (id),

	runtime_id INT NOT NULL,
	FOREIGN KEY (runtime_id) REFERENCES runtimes (id),

	pool_id INT NOT NULL,
	FOREIGN KEY (pool_id) REFERENCES pools (id),

	repository TEXT NOT NULL DEFAULT '',
	branch TEXT NOT NULL DEFAULT '',

	db_name TEXT NOT NULL DEFAULT '',
	db_user TEXT NOT NULL DEFAULT '',

	framework TEXT NOT NULL DEFAULT '',
	document_root TEXT NOT NULL,

	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX applications_client_id_project_id_env_id_unique ON applications (client_id, project_id, env_id);
CREATE UNIQUE INDEX applications_repository_branch_unique ON applications (repository, branch);

INSERT INTO resources (name) VALUES ('applications.view'), ('applications.create'), ('applications.update'), ('applications.delete');

-- Add the new resources to the admin role
-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
	INSERT INTO role_to_resources (role_id, resource_id)
		SELECT admin_role_id.id, resources.id FROM resources, admin_role_id WHERE resources.name IN ('applications.view', 'applications.create', 'applications.update', 'applications.delete');

CREATE TABLE application_to_domains (
	application_id INT NOT NULL,
	domain_id INT NOT NULL,
	PRIMARY KEY (application_id, domain_id),
	FOREIGN KEY (application_id) REFERENCES applications (id),
	FOREIGN KEY (domain_id) REFERENCES domains (id)
);
