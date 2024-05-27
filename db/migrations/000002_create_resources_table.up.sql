CREATE TABLE resources (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL
);

CREATE UNIQUE INDEX resources_name_unique ON resources (name);

INSERT INTO resources (name) VALUES ('users.view'), ('users.create'), ('users.update'), ('users.delete');
INSERT INTO resources (name) VALUES ('roles.view'), ('roles.create'), ('roles.update'), ('roles.delete');
