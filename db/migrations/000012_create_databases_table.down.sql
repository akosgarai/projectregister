DROP TABLE databases;

DELETE FROM resources WHERE name IN ('databases.view', 'databases.create', 'databases.update', 'databases.delete');
