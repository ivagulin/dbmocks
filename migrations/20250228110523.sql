-- +migrate Up
-- Drop index "idx_project_product_unique" from table: "tenants"
DROP INDEX IF EXISTS "idx_project_product_unique";
-- Modify "tenants" table
ALTER TABLE "tenants" ADD CONSTRAINT "uni_tenants_product_instance_id" UNIQUE ("product_instance_id"), ADD CONSTRAINT "uni_tenants_project_id" UNIQUE ("project_id");
