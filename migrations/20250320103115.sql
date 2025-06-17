-- +migrate Up
-- Drop index "mo_ref_vcenter_id" from table: "hard_disks"
DROP INDEX IF EXISTS "mo_ref_vcenter_id";
-- Create index "mo_ref_vcenter_id" to table: "hard_disks"
CREATE UNIQUE INDEX IF NOT EXISTS "mo_ref_vcenter_id" ON "hard_disks" ("v_center_id", "mo_ref");
