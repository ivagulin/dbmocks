-- +migrate Up
UPDATE virtual_machines SET iam_client_secret = '';