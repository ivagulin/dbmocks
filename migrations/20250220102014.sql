-- +migrate Up
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ADD COLUMN IF NOT EXISTS "bootable" boolean NULL;
