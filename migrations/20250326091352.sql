-- +migrate Up
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ADD COLUMN "billing_project_id" text NULL;
