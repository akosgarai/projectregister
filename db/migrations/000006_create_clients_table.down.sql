DROP TABLE clients;

DELETE FROM resources WHERE name IN ('clients.view', 'clients.create', 'clients.update', 'clients.delete');
