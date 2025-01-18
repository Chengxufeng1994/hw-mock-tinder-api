-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS postgis;

CREATE SCHEMA auths;

SET
SEARCH_PATH TO auths, PUBLIC;

CREATE TABLE otps (
  id            SERIAL,
  account_id    TEXT NOT NULL,
  code          VARCHAR(6) NOT NULL,
  is_verified   BOOLEAN DEFAULT FALSE,
  expiration    TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA IF EXISTS auths CASCADE;

DROP EXTENSION IF EXISTS postgis;
-- +goose StatementEnd

