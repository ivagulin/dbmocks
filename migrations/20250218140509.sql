-- +migrate Up
-- Modify "datastores" table
ALTER TABLE "datastores" ADD COLUMN IF NOT EXISTS "disk_resource_spec_code" text NULL;
-- Create index "datastores_mo_ref_v_center_id_uindex" to table: "datastores"
CREATE UNIQUE INDEX IF NOT EXISTS "datastores_mo_ref_v_center_id_uindex" ON "datastores" ("mo_ref", "v_center_id");
-- Modify "external_networks" table
ALTER TABLE "external_networks" ADD COLUMN IF NOT EXISTS "f_ip_resource_spec_code" text NULL;
-- Create index "external_networks_mo_ref_v_center_id_uindex" to table: "external_networks"
CREATE UNIQUE INDEX IF NOT EXISTS "external_networks_mo_ref_v_center_id_uindex" ON "external_networks" ("mo_ref", "v_center_id");
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ADD COLUMN IF NOT EXISTS "datastore_id" bigint NULL, ADD CONSTRAINT "fk_hard_disks_datastore" FOREIGN KEY ("datastore_id") REFERENCES "datastores" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "resource_pools" table
ALTER TABLE "resource_pools" ADD COLUMN IF NOT EXISTS "cpu_resource_spec_code" text NULL, ADD COLUMN IF NOT EXISTS "ram_resource_spec_code" text NULL;
-- Create index "resource_pools_mo_ref_v_center_id_uindex" to table: "resource_pools"
CREATE UNIQUE INDEX IF NOT EXISTS "resource_pools_mo_ref_v_center_id_uindex" ON "resource_pools" ("mo_ref", "v_center_id");
-- Rename a column from "project_id" to "log_project_id"
ALTER TABLE "virtual_machines" RENAME COLUMN "project_id" TO "log_project_id";
-- Modify "virtual_machines" table
ALTER TABLE "virtual_machines" ADD COLUMN IF NOT EXISTS "resource_pool_id" bigint NULL, ADD CONSTRAINT "fk_virtual_machines_resource_pool" FOREIGN KEY ("resource_pool_id") REFERENCES "resource_pools" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
