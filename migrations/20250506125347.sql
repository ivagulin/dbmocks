-- +migrate Up
-- Modify "v_centers" table
ALTER TABLE "v_centers" ADD COLUMN IF NOT EXISTS "os_compute_endpoint" text NULL, ADD COLUMN IF NOT EXISTS "os_image_endpoint" text NULL, ADD COLUMN IF NOT EXISTS "os_block_storage_endpoint" text NULL, ADD COLUMN IF NOT EXISTS "os_network_endpoint" text NULL;
-- Modify "vm_networks" table
ALTER TABLE "vm_networks" DROP COLUMN IF EXISTS "network_id", ADD COLUMN IF NOT EXISTS "network_ref" text NULL;
