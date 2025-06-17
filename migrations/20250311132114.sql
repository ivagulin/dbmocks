-- +migrate Up
-- Drop index "hard_disks_virtual_machine_id_name_uindex" from table: "hard_disks"
DROP INDEX "hard_disks_virtual_machine_id_name_uindex";
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" DROP CONSTRAINT "fk_hard_disks_datastore", ADD COLUMN "mo_ref" text NULL, ADD CONSTRAINT "fk_hard_disks_datastore" FOREIGN KEY ("datastore_id") REFERENCES "datastores" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Create index "hard_disks_mo_ref_virtual_machine_id_name_uindex" to table: "hard_disks"
CREATE UNIQUE INDEX "hard_disks_mo_ref_virtual_machine_id_name_uindex" ON "hard_disks" ("virtual_machine_id", "name", "mo_ref");
