-- +migrate Up
-- Modify "tenants" table
ALTER TABLE "tenants" ADD COLUMN "product_internal_name" text NULL;
