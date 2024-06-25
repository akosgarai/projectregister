DROP TABLE domains;

DELETE FROM resources WHERE name IN ('domains.view', 'domains.create', 'domains.update', 'domains.delete');
