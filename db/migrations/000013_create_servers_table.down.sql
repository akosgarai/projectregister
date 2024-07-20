DROP TABLE server_to_pool;
DROP TABLE server_to_runtime;
DROP TABLE servers;

DELETE FROM resources WHERE name IN ('servers.view', 'servers.create', 'servers.update', 'servers.delete');
