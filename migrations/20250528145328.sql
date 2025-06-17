-- +migrate Up
-- Modify "vm_networks_security_groups" table
ALTER TABLE "vm_networks_security_groups" DROP CONSTRAINT "fk_vm_networks_security_groups_security_group", DROP CONSTRAINT "fk_vm_networks_security_groups_vm_network", ADD CONSTRAINT "fk_vm_networks_security_groups_security_group" FOREIGN KEY ("security_group_id") REFERENCES "security_groups" ("id") ON UPDATE NO ACTION ON DELETE CASCADE, ADD CONSTRAINT "fk_vm_networks_security_groups_vm_network" FOREIGN KEY ("vm_network_id") REFERENCES "vm_networks" ("id") ON UPDATE NO ACTION ON DELETE CASCADE;
