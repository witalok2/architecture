BEGIN;

CREATE TABLE IF NOT EXISTS migration
(
    name character varying(100) NOT NULL PRIMARY KEY,
    applied_at timestamp without time zone NOT NULL
);

INSERT INTO migration VALUES ('00001_migration.up.sql', NOW());
COMMIT;
