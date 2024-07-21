DROP TABLE application_to_domains;
DROP TABLE applications;

DELETE FROM resources WHERE name IN ('applications.view', 'applications.create', 'applications.update', 'applications.delete');
