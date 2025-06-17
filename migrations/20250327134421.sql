-- +migrate Up
-- Modify "datastores" table
ALTER TABLE "datastores" ALTER COLUMN "mo_ref" SET NOT NULL;
-- Modify "external_networks" table
ALTER TABLE "external_networks" ALTER COLUMN "mo_ref" SET NOT NULL;
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ALTER COLUMN "mo_ref" SET NOT NULL, ALTER COLUMN "v_center_id" SET NOT NULL;
-- Modify "networks" table
ALTER TABLE "networks" ALTER COLUMN "mo_ref" SET NOT NULL;
-- Modify "resource_pools" table
ALTER TABLE "resource_pools" ALTER COLUMN "mo_ref" SET NOT NULL;
-- Modify "virtual_machines" table
ALTER TABLE "virtual_machines" ALTER COLUMN "mo_ref" SET NOT NULL;
-- Modify "vm_networks" table
ALTER TABLE "vm_networks" ALTER COLUMN "virtual_machine_id" SET NOT NULL;
