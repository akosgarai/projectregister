-- Add a role_id column to the users table with a default value of the user role.
ALTER TABLE users ADD COLUMN role_id INT NOT NULL DEFAULT 2;

-- add a foreign key constraint to the role_id column
ALTER TABLE users ADD CONSTRAINT users_role_id_foreign FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;

-- If the admin user does not exist, insert the admin user into the users table.
INSERT INTO users (name, email, password, role_id) VALUES ('Admin', 'system@admin', '$2a$10$8QIzpaZqZEEI3RVKKjGnh.GJ3DqLEIewuuRMGGCnRD3VW3v7ZodUW', (SELECT id FROM roles WHERE name = 'admin'))
ON CONFLICT (email) DO UPDATE SET role_id = (SELECT id FROM roles WHERE name = 'admin');
