-- +migrate Up
-- Create index "external_networks_mo_ref_v_center_id_uindex" to table: "external_networks"
CREATE UNIQUE INDEX IF NOT EXISTS "external_networks_mo_ref_v_center_id_uindex" ON "external_networks" ("mo_ref", "v_center_id");
-- Create index "networks_mo_ref_v_center_id_uindex" to table: "networks"
CREATE UNIQUE INDEX IF NOT EXISTS "networks_mo_ref_v_center_id_uindex" ON "networks" ("mo_ref", "v_center_id");
-- Modify "vm_networks" table
ALTER TABLE "vm_networks" ADD COLUMN IF NOT EXISTS "network_id" bigint NULL, ADD CONSTRAINT "fk_vm_networks_network" FOREIGN KEY ("network_id") REFERENCES "networks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
