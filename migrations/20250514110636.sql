-- +migrate Up
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" DROP COLUMN IF EXISTS "virtual_machine_name";
