CREATE TABLE environment_to_servers (
	environment_id INT NOT NULL,
	server_id INT NOT NULL,
	PRIMARY KEY (environment_id, server_id),
	FOREIGN KEY (environment_id) REFERENCES environments (id),
	FOREIGN KEY (server_id) REFERENCES servers (id)
);
