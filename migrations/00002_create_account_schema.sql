-- +goose Up
CREATE SCHEMA accounts;

SET
SEARCH_PATH TO accounts, PUBLIC;

CREATE TABLE accounts (
  id           TEXT         NOT NULL,
  email        VARCHAR(255) UNIQUE,
  phone_number VARCHAR(15)  UNIQUE,
  created_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  deleted_at   TIMESTAMP,
  PRIMARY KEY (id)
);

CREATE INDEX idx_accounts_deleted_at ON accounts (deleted_at);
-- +goose Down
DROP SCHEMA IF EXISTS accounts CASCADE;

