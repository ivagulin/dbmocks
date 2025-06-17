-- +migrate Up
-- Create "v_centers" table
CREATE TABLE IF NOT EXISTS "v_centers" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "url" text NULL,
  "login" text NULL,
  "password" text NULL,
  "insecure" boolean NULL,
  "type" text NOT NULL DEFAULT 'vcenter',
  "os_domain_name" text NULL,
  "os_tenant_name" text NULL,
  "customer_id" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "distributed_v_switches" table
CREATE TABLE IF NOT EXISTS "distributed_v_switches" (
  "id" bigserial NOT NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "num_ports" integer NULL,
  "num_uplinks" integer NULL,
  "num_hosts" integer NULL,
  "description" text NULL,
  "v_center_id" bigint NOT NULL,
  "datacenter_ref" text NULL,
  PRIMARY KEY ("id")
);
-- Create "openstack_hypervisors_stats" table
CREATE TABLE IF NOT EXISTS "openstack_hypervisors_stats" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "v_center_id" bigint NOT NULL,
  "hv_count" bigint NULL,
  "running_vms_count" bigint NULL,
  "v_cpus" bigint NULL,
  "v_cpus_used" bigint NULL,
  "memory_mb" bigint NULL,
  "memory_mb_used" bigint NULL,
  "memory_mb_free" bigint NULL,
  "disk_gb" bigint NULL,
  "disk_gb_used" bigint NULL,
  "disk_gb_free" bigint NULL,
  "disk_gb_available" bigint NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_openstack_hypervisors_stats_v_center_id" UNIQUE ("v_center_id")
);
-- Create "datastores" table
CREATE TABLE IF NOT EXISTS "datastores" (
  "id" bigserial NOT NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "type" text NULL,
  "capacity" bigint NULL,
  "free_space" bigint NULL,
  "v_center_id" bigint NOT NULL,
  "datacenter_ref" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_datastores_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "networks" table
CREATE TABLE IF NOT EXISTS "networks" (
  "id" bigserial NOT NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "type" text NULL,
  "deletable" boolean NULL,
  "v_center_id" bigint NOT NULL,
  "datacenter_ref" text NULL,
  PRIMARY KEY ("id")
);
-- Create "distributed_port_group_configs" table
CREATE TABLE IF NOT EXISTS "distributed_port_group_configs" (
  "id" bigserial NOT NULL,
  "name" text NULL,
  "network_id" bigint NULL,
  "network_ref" text NULL,
  "vlan_type" integer NULL,
  "vlan_id" integer NULL,
  "port_binding" integer NULL,
  "auto_expand" boolean NULL,
  "num_ports" integer NULL,
  "promiscuous_mode" boolean NULL,
  "mac_address_changes" boolean NULL,
  "forged_transmits" boolean NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_networks_distributed_config" FOREIGN KEY ("network_id") REFERENCES "networks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "distributed_port_group_vlan_trunk_ranges" table
CREATE TABLE IF NOT EXISTS "distributed_port_group_vlan_trunk_ranges" (
  "id" bigserial NOT NULL,
  "start" integer NULL,
  "end" integer NULL,
  "distributed_port_group_config_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_distributed_port_group_configs_vlan_trunk_ranges" FOREIGN KEY ("distributed_port_group_config_id") REFERENCES "distributed_port_group_configs" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "virtual_machines" table
CREATE TABLE IF NOT EXISTS "virtual_machines" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "is_template" boolean NULL,
  "is_user_created_template" boolean NULL,
  "cpus" integer NULL,
  "memory_mb" bigint NULL,
  "guest_system" text NULL,
  "storage_capacity" bigint NULL,
  "power_state" text NULL,
  "ip_address" text NULL,
  "floating_ip" text NULL,
  "is_billed" boolean NOT NULL DEFAULT false,
  "v_center_id" bigint NOT NULL,
  "datacenter_ref" text NULL,
  "osconfig_enabled" boolean NULL,
  "osconfig_endpoint" text NULL,
  "osconfig_token" text NULL,
  "osconfig_com_name" text NULL,
  "lo_gaa_s_enabled" boolean NULL,
  "iam_address" text NULL,
  "iam_client_id" text NULL,
  "iam_client_secret" text NULL,
  "project_id" text NULL,
  "log_group_id" text NULL,
  "lo_gaa_s_address" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_v_centers_virtual_machines" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "virtual_machines_mo_ref_v_center_id_uindex" to table: "virtual_machines"
CREATE UNIQUE INDEX IF NOT EXISTS "virtual_machines_mo_ref_v_center_id_uindex" ON "virtual_machines" ("mo_ref", "v_center_id");
-- Create "hard_disks" table
CREATE TABLE IF NOT EXISTS "hard_disks" (
  "id" bigserial NOT NULL,
  "virtual_machine_id" bigint NULL,
  "name" text NULL,
  "size" bigint NULL,
  "controller" text NULL,
  "mode" text NULL,
  "filename" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_hard_disks_virtual_machine" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "hard_disks_virtual_machine_id_name_uindex" to table: "hard_disks"
CREATE UNIQUE INDEX IF NOT EXISTS "hard_disks_virtual_machine_id_name_uindex" ON "hard_disks" ("virtual_machine_id", "name");
-- Create "resource_pools" table
CREATE TABLE IF NOT EXISTS "resource_pools" (
  "id" bigserial NOT NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "cpu_limit" bigint NULL,
  "cores" bigint NULL,
  "memory_limit" bigint NULL,
  "v_center_id" bigint NOT NULL,
  "datacenter_ref" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_resource_pools_v_center" FOREIGN KEY ("v_center_id") REFERENCES "v_centers" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create "tenants" table
CREATE TABLE IF NOT EXISTS "tenants" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "project_id" uuid NULL,
  "product_instance_id" uuid NULL,
  "customer_id" uuid NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_tenants_product_instance_id" UNIQUE ("product_instance_id"),
  CONSTRAINT "uni_tenants_project_id" UNIQUE ("project_id")
);
-- Create "tenants_datastores" table
CREATE TABLE IF NOT EXISTS "tenants_datastores" (
  "tenant_id" bigint NOT NULL,
  "datastore_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "datastore_id"),
  CONSTRAINT "fk_tenants_datastores_datastore" FOREIGN KEY ("datastore_id") REFERENCES "datastores" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_datastores_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "external_networks" table
CREATE TABLE IF NOT EXISTS "external_networks" (
  "id" bigserial NOT NULL,
  "mo_ref" text NULL,
  "name" text NULL,
  "v_center_id" bigint NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "tenants_external_networks" table
CREATE TABLE IF NOT EXISTS "tenants_external_networks" (
  "tenant_id" bigint NOT NULL,
  "external_network_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "external_network_id"),
  CONSTRAINT "fk_tenants_external_networks_external_network" FOREIGN KEY ("external_network_id") REFERENCES "external_networks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_external_networks_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "tenants_networks" table
CREATE TABLE IF NOT EXISTS "tenants_networks" (
  "tenant_id" bigint NOT NULL,
  "network_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "network_id"),
  CONSTRAINT "fk_tenants_networks_network" FOREIGN KEY ("network_id") REFERENCES "networks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_networks_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "tenants_resource_pools" table
CREATE TABLE IF NOT EXISTS "tenants_resource_pools" (
  "tenant_id" bigint NOT NULL,
  "resource_pool_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "resource_pool_id"),
  CONSTRAINT "fk_tenants_resource_pools_resource_pool" FOREIGN KEY ("resource_pool_id") REFERENCES "resource_pools" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_resource_pools_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "tenants_virtual_machines" table
CREATE TABLE IF NOT EXISTS "tenants_virtual_machines" (
  "tenant_id" bigint NOT NULL,
  "virtual_machine_id" bigint NOT NULL,
  PRIMARY KEY ("tenant_id", "virtual_machine_id"),
  CONSTRAINT "fk_tenants_virtual_machines_tenant" FOREIGN KEY ("tenant_id") REFERENCES "tenants" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "fk_tenants_virtual_machines_virtual_machine" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
-- Create "vm_networks" table
CREATE TABLE IF NOT EXISTS "vm_networks" (
  "id" bigserial NOT NULL,
  "virtual_machine_id" bigint NULL,
  "name" text NULL,
  "adapter" text NULL,
  "power" boolean NULL,
  "ip_address" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_vm_networks_virtual_machine" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "vm_networks_virtual_machine_id_name_uindex" to table: "vm_networks"
CREATE UNIQUE INDEX IF NOT EXISTS "vm_networks_virtual_machine_id_name_uindex" ON "vm_networks" ("virtual_machine_id", "name");
-- Create "vm_floating_ips" table
CREATE TABLE IF NOT EXISTS "vm_floating_ips" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "virtual_machine_id" bigint NOT NULL,
  "external_network_id" bigint NOT NULL,
  "vm_network_id" bigint NOT NULL,
  "floating_ip" text NOT NULL,
  "floating_ip_ref" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_vm_floating_ips_floating_ip_ref" UNIQUE ("floating_ip_ref"),
  CONSTRAINT "fk_vm_floating_ips_external_network" FOREIGN KEY ("external_network_id") REFERENCES "external_networks" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_vm_floating_ips_virtual_machine" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT "fk_vm_floating_ips_vm_network" FOREIGN KEY ("vm_network_id") REFERENCES "vm_networks" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create "vm_snapshots" table
CREATE TABLE IF NOT EXISTS "vm_snapshots" (
  "id" bigserial NOT NULL,
  "virtual_machine_id" bigint NULL,
  "name" text NULL,
  "description" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_vm_snapshots_virtual_machine" FOREIGN KEY ("virtual_machine_id") REFERENCES "virtual_machines" ("id") ON UPDATE CASCADE ON DELETE CASCADE
);
-- Create index "vm_snapshots_by_vmid_name_unique" to table: "vm_snapshots"
CREATE UNIQUE INDEX IF NOT EXISTS "vm_snapshots_by_vmid_name_unique" ON "vm_snapshots" ("virtual_machine_id", "name");
