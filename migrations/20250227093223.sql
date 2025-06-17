-- +migrate Up
-- Drop index "external_networks_mo_ref_v_center_id_uindex" from table: "external_networks"
DROP INDEX IF EXISTS "external_networks_mo_ref_v_center_id_uindex";
-- Modify "tenants" table
ALTER TABLE "tenants" DROP CONSTRAINT IF EXISTS "uni_tenants_product_instance_id", DROP CONSTRAINT IF EXISTS "uni_tenants_project_id";
-- Create index "idx_project_product_unique" to table: "tenants"
CREATE UNIQUE INDEX IF NOT EXISTS "idx_project_product_unique" ON "tenants" ("project_id", "product_instance_id");
