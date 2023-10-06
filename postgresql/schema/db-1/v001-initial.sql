DROP SCHEMA IF EXISTS service_template_schema CASCADE;

CREATE SCHEMA IF NOT EXISTS service_template_schema;

CREATE TABLE IF NOT EXISTS service_template_schema.users (
    name TEXT PRIMARY KEY,
    count INTEGER DEFAULT(1)
);
