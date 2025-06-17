-- +migrate Up
-- Modify "external_networks" table
ALTER TABLE "external_networks" ADD CONSTRAINT "fk_external_networks_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "hard_disks" table
ALTER TABLE "hard_disks" ADD CONSTRAINT "fk_hard_disks_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
-- Modify "networks" table
ALTER TABLE "networks" ADD CONSTRAINT "fk_networks_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
