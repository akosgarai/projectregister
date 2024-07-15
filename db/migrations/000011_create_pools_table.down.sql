DROP TABLE pools;

DELETE FROM resources WHERE name IN ('pools.view', 'pools.create', 'pools.update', 'pools.delete');
