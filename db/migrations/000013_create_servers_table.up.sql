CREATE TABLE servers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	description TEXT,
	remote_address VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX servers_name_unique ON servers (name);
CREATE UNIQUE INDEX servers_remote_address_unique ON servers (remote_address);

INSERT INTO resources (name) VALUES ('servers.view'), ('servers.create'), ('servers.update'), ('servers.delete');

-- Add the new resources to the admin role
-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
	INSERT INTO role_to_resources (role_id, resource_id)
		SELECT admin_role_id.id, resources.id FROM resources, admin_role_id WHERE resources.name IN ('servers.view', 'servers.create', 'servers.update', 'servers.delete');

CREATE TABLE server_to_pool (
	server_id INT NOT NULL,
	pool_id INT NOT NULL,
	PRIMARY KEY (server_id, pool_id),
	FOREIGN KEY (server_id) REFERENCES servers (id),
	FOREIGN KEY (pool_id) REFERENCES pools (id)
);

CREATE TABLE server_to_runtime (
	server_id INT NOT NULL,
	runtime_id INT NOT NULL,
	PRIMARY KEY (server_id, runtime_id),
	FOREIGN KEY (server_id) REFERENCES servers (id),
	FOREIGN KEY (runtime_id) REFERENCES runtimes (id)
);
