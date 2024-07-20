CREATE TABLE environment_to_databases (
	environment_id INT NOT NULL,
	database_id INT NOT NULL,
	PRIMARY KEY (environment_id, database_id),
	FOREIGN KEY (environment_id) REFERENCES environments (id),
	FOREIGN KEY (database_id) REFERENCES databases (id)
);
