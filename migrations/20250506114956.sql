-- +migrate Up
-- Create "security_groups" table
CREATE TABLE "security_groups" (
  "id" bigserial NOT NULL,
  "v_center_id" bigint NOT NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "az" text NULL,
  "mo_ref" text NULL,
  "tags" character varying[] NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_security_groups_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "security_group_rules" table
CREATE TABLE "security_group_rules" (
  "id" bigserial NOT NULL,
  "security_group_id" bigint NULL,
  "mo_ref" text NULL,
  "direction" integer NULL,
  "protocol" text NULL,
  "port_range_min" bigint NULL,
  "port_range_max" bigint NULL,
  "remote_group_id" text NULL,
  "remote_ip_prefix" text NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_security_groups_rules" FOREIGN KEY ("security_group_id") REFERENCES "security_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "tenants_security_groups" table
CREATE TABLE "tenants_security_groups" (
  "tenant_id" bigint NOT NULL,
  "security_group_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "security_group_id"),
  CONSTRAINT "fk_tenants_security_groups_security_group" FOREIGN KEY ("security_group_id") REFERENCES "security_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_security_groups_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
