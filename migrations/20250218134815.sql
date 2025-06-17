-- +migrate Up
-- Modify "datastores" table
ALTER TABLE "datastores" ADD COLUMN "project_default" boolean NULL DEFAULT false;
-- Modify "external_networks" table
ALTER TABLE "external_networks" ADD COLUMN "project_default" boolean NULL DEFAULT false;
-- Modify "networks" table
ALTER TABLE "networks" ADD COLUMN "project_default" boolean NULL DEFAULT false;
-- Modify "resource_pools" table
ALTER TABLE "resource_pools" ADD COLUMN "project_default" boolean NULL DEFAULT false;
-- Modify "tenants" table
ALTER TABLE "tenants" ADD COLUMN "suspended" boolean NULL DEFAULT false;
-- Modify "virtual_machines" table
ALTER TABLE "virtual_machines" ADD COLUMN "project_default" boolean NULL DEFAULT false;
