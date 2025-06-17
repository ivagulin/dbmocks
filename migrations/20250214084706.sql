-- +migrate Up
-- Modify "datastores" table
ALTER TABLE "datastores" DROP CONSTRAINT "fk_datastores_v_center", ADD CONSTRAINT "fk_datastores_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "virtual_machines" DROP CONSTRAINT "fk_v_centers_virtual_machines", ADD CONSTRAINT "fk_v_centers_virtual_machines" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE "resource_pools" DROP CONSTRAINT "fk_resource_pools_v_center", ADD CONSTRAINT "fk_resource_pools_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;