DROP TABLE projects;

DELETE FROM resources WHERE name IN ('projects.view', 'projects.create', 'projects.update', 'projects.delete');
