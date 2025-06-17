-- +migrate Up
-- Modify "virtual_machines" table
ALTER TABLE "virtual_machines" ADD COLUMN "billing_project_id" text NULL;
