-- Add the has_ssl column to the domains table with a default value of false.
ALTER TABLE domains ADD COLUMN has_ssl BOOLEAN DEFAULT FALSE;
