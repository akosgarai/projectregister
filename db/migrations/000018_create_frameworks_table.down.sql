-- create the framework column again
ALTER TABLE applications ADD COLUMN framework TEXT NOT NULL DEFAULT '';
-- add the string values to the new column
UPDATE applications SET framework = frameworks.name FROM frameworks WHERE applications.framework_id = frameworks.id;
-- drop the old column
ALTER TABLE applications DROP COLUMN framework_id;

DROP TABLE frameworks;

DELETE FROM resources WHERE name IN ('frameworks.view', 'frameworks.create', 'frameworks.update', 'frameworks.delete');
