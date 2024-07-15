DROP TABLE runtimes;

DELETE FROM resources WHERE name IN ('runtimes.view', 'runtimes.create', 'runtimes.update', 'runtimes.delete');
