BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS "user"
(
  id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
  name varchar NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp NULL
);

INSERT INTO migration VALUES ('00002_create_user.up.sql', NOW());
COMMIT;
