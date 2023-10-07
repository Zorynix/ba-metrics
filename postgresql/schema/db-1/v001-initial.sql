DROP SCHEMA IF EXISTS ba_schema CASCADE;

CREATE SCHEMA IF NOT EXISTS ba_schema;

CREATE TABLE IF NOT EXISTS ba_schema.links (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    url text NOT NULL
);