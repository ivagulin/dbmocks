-- +migrate Up
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL;
