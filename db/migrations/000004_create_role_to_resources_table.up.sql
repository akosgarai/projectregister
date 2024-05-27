CREATE TABLE role_to_resources (
	id SERIAL PRIMARY KEY,
	role_id INT NOT NULL,
	resource_id INT NOT NULL
);

CREATE UNIQUE INDEX role_to_resources_role_id_resource_id_unique ON role_to_resources (role_id, resource_id);

ALTER TABLE role_to_resources ADD CONSTRAINT role_to_resources_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE role_to_resources ADD CONSTRAINT role_to_resources_resource_id_foreign FOREIGN KEY (resource_id) REFERENCES resources (id) ON DELETE CASCADE;

-- add every resource to admin
WITH admin_role_id AS (SELECT id FROM roles WHERE name = 'admin')
INSERT INTO role_to_resources (role_id, resource_id) SELECT admin_role_id.id, resources.id FROM resources, admin_role_id;

-- add every *.view resource to user
WITH user_role_id AS (SELECT id FROM roles WHERE name = 'user')
INSERT INTO role_to_resources (role_id, resource_id) SELECT user_role_id.id, resources.id FROM resources, user_role_id WHERE resources.name LIKE '%.view';
