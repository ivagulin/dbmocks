-- +migrate Up
-- Drop index "hard_disks_mo_ref_virtual_machine_id_name_uindex" from table: "hard_disks"
DROP INDEX "hard_disks_mo_ref_virtual_machine_id_name_uindex";
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" DROP CONSTRAINT "fk_hard_disks_virtual_machine", ADD COLUMN "virtual_machine_name" text NULL, ADD COLUMN "v_center_id" bigint NULL, ADD COLUMN "project_default" boolean NULL DEFAULT false, ADD CONSTRAINT "fk_virtual_machines_hard_disks" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE NO ACTION ON DELETE SET NULL;
-- Create index "mo_ref_vcenter_id" to table: "hard_disks"
CREATE INDEX "mo_ref_vcenter_id" ON "hard_disks" ("v_center_id", "mo_ref");
-- Create "tenants_hard_disks" table
CREATE TABLE "tenants_hard_disks" (
  "tenant_id" bigint NOT NULL,
  "hard_disk_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "hard_disk_id"),
  CONSTRAINT "fk_tenants_hard_disks_hard_disk" FOREIGN KEY ("hard_disk_id") REFERENCES "hard_disks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_hard_disks_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
