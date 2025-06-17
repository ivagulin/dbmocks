-- +migrate Up
-- Modify "security_groups" table
ALTER TABLE "security_groups" ADD COLUMN "created_at" timestamptz NULL, ADD COLUMN "updated_at" timestamptz NULL, ADD COLUMN "project_default" boolean NULL DEFAULT false, ADD COLUMN "default" boolean NULL DEFAULT false;
-- Create "vm_networks_security_groups" table
CREATE TABLE "vm_networks_security_groups" (
  "security_group_id" bigint NOT NULL,
  "vm_network_id" bigint NOT NULL,
  PRIMARY KEY ("security_group_id", "vm_network_id"),
  CONSTRAINT "fk_vm_networks_security_groups_security_group" FOREIGN KEY ("security_group_id") REFERENCES "security_groups" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_vm_networks_security_groups_vm_network" FOREIGN KEY ("vm_network_id") REFERENCES "vm_networks" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
