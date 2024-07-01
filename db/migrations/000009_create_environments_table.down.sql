DROP TABLE environments;

DELETE FROM resources WHERE name IN ('environments.view', 'environments.create', 'environments.update', 'environments.delete');
