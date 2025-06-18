package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/ivagulin/dbmocks/migrations"
	"github.com/lib/pq"
	"github.com/rubenv/pgtest"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"log/slog"
	"os"
	"time"
)

type Repo struct {
	db  *gorm.DB
	dsn string
}

var (
	ErrNotFound   = errors.New("not found")
	ErrVMNotFound = errors.New("vm not found")

	Models = []any{
		&VCenter{},
		&VirtualMachine{},
		&Datastore{},
		&ResourcePool{},
		&Network{},
		&ExternalNetwork{},
		&HardDisk{},
		&VMNetwork{},
		&VMFloatingIP{},
		&VMSnapshot{},
		&DistributedVSwitch{},
		&DistributedPortGroupConfig{},
		&DistributedPortGroupVlanTrunkRange{},
		&OpenstackHypervisorsStats{},
		&SecurityGroup{},
		&SecurityGroupRule{},
	}
)

type ObjectsReport struct {
	PoolCount      int64
	DatastoreCount int64
	NetworkCount   int64
	TmplCount      int64
	VMCount        int64
	TntCount       int64
	VCenterCount   int64
}

type VCenter struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	URL                    string
	Login                  string
	Password               string
	Insecure               bool
	Type                   string `gorm:"not null;default:vcenter"`
	OSDomainName           string
	OSTenantName           string
	OSComputeEndpoint      string
	OSImageEndpoint        string
	OSBlockStorageEndpoint string
	OSNetworkEndpoint      string

	VirtualMachines []VirtualMachine `gorm:"constraint:OnDelete:CASCADE"`
	CustomerID      string           `gorm:"not null;default:null"`
}

type VirtualMachine struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	MoRef                 string `gorm:"not null;uniqueIndex:virtual_machines_mo_ref_v_center_id_uindex"`
	Name                  string
	IsTemplate            bool
	IsUserCreatedTemplate bool
	CPUs                  int32
	MemoryMB              int64
	GuestSystem           string
	StorageCapacity       int64
	PowerState            string
	IPAddress             string
	FloatingIP            *string
	IsBilled              bool `gorm:"not null;default:false"`
	BillingProjectID      string

	VCenterID      int64   `gorm:"not null;default:null;uniqueIndex:virtual_machines_mo_ref_v_center_id_uindex"`
	VCenter        VCenter `gorm:"constraint:OnDelete:CASCADE"`
	ResourcePoolID *int64
	ResourcePool   ResourcePool `gorm:"constraint:OnDelete:CASCADE"`
	DatacenterRef  string
	Tenants        []Tenant    `gorm:"many2many:tenants_virtual_machines;constraint:OnDelete:CASCADE"`
	HardDisks      []*HardDisk `gorm:"constraint:OnDelete:SET NULL"`

	OsconfigEnabled  bool
	OsconfigEndpoint string
	OsconfigToken    string
	OsconfigComName  string

	LOGaaSEnabled   bool
	IAMAddress      string
	IAMClientID     string
	IAMClientSecret string
	LogProjectID    string
	LogGroupID      string
	LOGaaSAddress   string

	ProjectDefault bool `gorm:"default:false"`
}

func (v *VirtualMachine) GetVCenter() VCenter {
	return v.VCenter
}

type Datastore struct {
	ID                   int64  `gorm:"primarykey"`
	MoRef                string `gorm:"not null;uniqueIndex:datastores_mo_ref_v_center_id_uindex"`
	Name                 string
	Type                 string
	Capacity             int64
	FreeSpace            int64
	DiskResourceSpecCode string

	VCenterID     int64   `gorm:"not null;default:null;uniqueIndex:datastores_mo_ref_v_center_id_uindex"`
	VCenter       VCenter `gorm:"constraint:OnDelete:CASCADE"`
	DatacenterRef string
	Tenants       []Tenant `gorm:"many2many:tenants_datastores;"`

	ProjectDefault bool `gorm:"default:false"`
}

func (v *Datastore) GetVCenter() VCenter {
	return v.VCenter
}

type ResourcePool struct {
	ID                  int64  `gorm:"primarykey"`
	MoRef               string `gorm:"not null;uniqueIndex:resource_pools_mo_ref_v_center_id_uindex"`
	Name                string
	CPULimit            int64
	Cores               int64
	MemoryLimit         int64
	VCenterID           int64   `gorm:"not null;default:null;uniqueIndex:resource_pools_mo_ref_v_center_id_uindex"`
	VCenter             VCenter `gorm:"constraint:OnDelete:CASCADE"`
	DatacenterRef       string
	Tenants             []Tenant `gorm:"many2many:tenants_resource_pools;"`
	CPUResourceSpecCode string
	RAMResourceSpecCode string
	ProjectDefault      bool `gorm:"default:false"`
}

func (v *ResourcePool) GetVCenter() VCenter {
	return v.VCenter
}

type Tenant struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ProjectID         uuid.UUID `gorm:"type:uuid;unique"`
	ProductInstanceID uuid.UUID `gorm:"type:uuid;unique"`
	CustomerID        uuid.UUID `gorm:"type:uuid;default:null"`

	Suspended           bool `gorm:"default:false"`
	ProductInternalName string

	ResourcePools    []ResourcePool    `gorm:"many2many:tenants_resource_pools;constraint:OnDelete:CASCADE"`
	Networks         []Network         `gorm:"many2many:tenants_networks;constraint:OnDelete:CASCADE"`
	ExternalNetworks []ExternalNetwork `gorm:"many2many:tenants_external_networks;constraint:OnDelete:CASCADE"`
	Datastores       []Datastore       `gorm:"many2many:tenants_datastores;constraint:OnDelete:CASCADE"`
	VirtualMachines  []VirtualMachine  `gorm:"many2many:tenants_virtual_machines;constraint:OnDelete:CASCADE"`
	HardDisks        []HardDisk        `gorm:"many2many:tenants_hard_disks;constraint:OnDelete:CASCADE"`
	SecurityGroups   []SecurityGroup   `gorm:"many2many:tenants_security_groups;constraint:OnDelete:CASCADE"`
}

type Network struct {
	ID                int64  `gorm:"primarykey"`
	MoRef             string `gorm:"not null;uniqueIndex:networks_mo_ref_v_center_id_uindex"`
	Name              string
	Type              string
	Deletable         bool
	DistributedConfig *DistributedPortGroupConfig `gorm:"foreignKey:NetworkID;constraint:OnDelete:CASCADE;"` // Ensure cascading delete

	VCenterID     int64   `gorm:"not null;default:null;uniqueIndex:networks_mo_ref_v_center_id_uindex"`
	VCenter       VCenter `gorm:"constraint:OnDelete:CASCADE"`
	DatacenterRef string
	Tenants       []Tenant `gorm:"many2many:tenants_networks;"`

	ProjectDefault bool `gorm:"default:false"`
}

func (v *Network) GetVCenter() VCenter {
	return v.VCenter
}

type ExternalNetwork struct {
	ID        int64  `gorm:"primarykey"`
	MoRef     string `gorm:"not null;uniqueIndex:external_networks_mo_ref_v_center_id_uindex"`
	Name      string
	VCenterID int64    `gorm:"not null;default:null;uniqueIndex:external_networks_mo_ref_v_center_id_uindex"`
	VCenter   VCenter  `gorm:"constraint:OnDelete:CASCADE"`
	Tenants   []Tenant `gorm:"many2many:tenants_external_networks;"`

	ProjectDefault      bool `gorm:"default:false"`
	FIPResourceSpecCode string
}

func (v *ExternalNetwork) GetVCenter() VCenter {
	return v.VCenter
}

type HardDisk struct {
	ID               int64 `gorm:"primarykey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	VirtualMachineID *int64 `gorm:"constraint:OnDelete:SET NULL"`
	VirtualMachine   *VirtualMachine
	VCenterID        int64   `gorm:"not null;uniqueIndex:mo_ref_vcenter_id"`
	VCenter          VCenter `gorm:"constraint:OnDelete:CASCADE"`
	DatastoreID      int64
	Datastore        Datastore `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Tenants          []Tenant  `gorm:"many2many:tenants_hard_disks;constraint:OnDelete:CASCADE"`
	Name             string
	Size             int64
	Controller       *string
	Mode             *string
	Filename         string
	Bootable         *bool
	MoRef            string `gorm:"not null;uniqueIndex:mo_ref_vcenter_id"`
	ProjectDefault   bool   `gorm:"default:false"`
	BillingProjectID string
}

func (v *HardDisk) GetVCenter() VCenter {
	return v.VCenter
}

type VMNetwork struct {
	ID               int64          `gorm:"primarykey"`
	VirtualMachineID int64          `gorm:"not null;uniqueIndex:vm_networks_virtual_machine_id_name_uindex"`
	VirtualMachine   VirtualMachine `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	NetworkRef       string
	Name             string `gorm:"uniqueIndex:vm_networks_virtual_machine_id_name_uindex"`
	Adapter          string
	Power            bool
	IPAddress        string
	SecurityGroups   []SecurityGroup `gorm:"many2many:vm_networks_security_groups;constraint:OnDelete:CASCADE"`
}

type VMFloatingIP struct {
	ID                int64 `gorm:"primarykey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	VirtualMachineID  int64           `gorm:"not null;default:null"`
	VirtualMachine    VirtualMachine  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExternalNetworkID int64           `gorm:"not null;default:null"`
	ExternalNetwork   ExternalNetwork `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	VMNetworkID       int64           `gorm:"not null;default:null"`
	VMNetwork         VMNetwork       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FloatingIP        string          `gorm:"not null;default:null"`
	FloatingIPRef     string          `gorm:"unique;not null;default:null"`
}

type VMSnapshot struct {
	ID               int64          `gorm:"primarykey"`
	VirtualMachineID int64          `gorm:"uniqueIndex:vm_snapshots_by_vmid_name_unique"`
	VirtualMachine   VirtualMachine `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name        string `gorm:"uniqueIndex:vm_snapshots_by_vmid_name_unique"` // Наименование снимка виртулаьной машины.
	Description string // Описание снимка виртулаьной машины.
}

type DistributedVSwitch struct {
	ID    int64 `gorm:"primarykey"`
	MoRef string
	Name  string

	NumPorts   int32
	NumUplinks int32

	NumHosts    int32
	Description string

	VCenterID     int64 `gorm:"not null;default:null"`
	DatacenterRef string
}

type VlanType int32

const (
	VlanNone  VlanType = 0
	Vlan      VlanType = 1
	VlanTrunk VlanType = 2
)

type DistributedPortGroupVlanTrunkRange struct {
	ID                           int64 `gorm:"primarykey"`
	Start                        int32
	End                          int32
	DistributedPortGroupConfigID int64 `gorm:"not null;default:null"`
}

type DistributedPortGroupConfig struct {
	ID         int64 `gorm:"primarykey"`
	Name       string
	NetworkID  int64
	NetworkRef string

	VlanType        VlanType
	VlanID          int32
	VlanTrunkRanges []DistributedPortGroupVlanTrunkRange `gorm:"foreignKey:DistributedPortGroupConfigID;constraint:OnDelete:CASCADE;"`

	PortBinding int32
	AutoExpand  bool
	NumPorts    int32

	PromiscuousMode   bool
	MACAddressChanges bool
	ForgedTransmits   bool
}

type OpenstackHypervisorsStats struct {
	ID              int64 `gorm:"primarykey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	VCenterID       int64 `gorm:"not null;default:null;unique"`
	HVCount         *int64
	RunningVMsCount *int64
	VCPUs           *int64
	VCPUsUsed       *int64
	MemoryMB        *int64
	MemoryMBUsed    *int64
	MemoryMBFree    *int64
	DiskGB          *int64
	DiskGBUsed      *int64
	DiskGBFree      *int64
	DiskGBAvailable *int64
}

type SecurityGroup struct {
	ID             int64  `gorm:"primaryKey"`
	VCenterID      int64  `gorm:"not null;default:null"`
	Name           string `gorm:"not null"`
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	AZ             string
	Rules          []SecurityGroupRule `gorm:"foreignKey:SecurityGroupID;constraint:OnDelete:CASCADE"`
	MoRef          string
	Tags           pq.StringArray `gorm:"type:varchar[]"`
	VCenter        VCenter        `gorm:"constraint:OnDelete:CASCADE"`
	Tenants        []Tenant       `gorm:"many2many:tenants_security_groups;constraint:OnDelete:CASCADE"`
	ProjectDefault bool           `gorm:"default:false"`
	VMNetworks     []VMNetwork    `gorm:"many2many:vm_networks_security_groups;constraint:OnDelete:CASCADE"`
	Default        bool           `gorm:"default:false"`
}

func (v *SecurityGroup) GetVCenter() VCenter {
	return v.VCenter
}

type SecurityGroupRule struct {
	ID              int64 `gorm:"primaryKey"`
	SecurityGroupID int64 `gorm:"constraint:OnDelete:CASCADE"`
	MoRef           string
	Direction       SecurityGroupRuleDirectionType // e.g., "ingress" or "egress"
	Protocol        string                         // e.g., "tcp", "udp", "icmp"
	PortRangeMin    *int64                         // nullable
	PortRangeMax    *int64                         // nullable
	RemoteGroupID   *string
	RemoteIPPrefix  *string
	Description     string
}

type SecurityGroupRuleDirectionType int32

const (
	Invalid SecurityGroupRuleDirectionType = iota
	Ingress
	Egress
)

var SecurityGroupRuleDirectionTypeByNameMap = map[string]SecurityGroupRuleDirectionType{
	"invalid": Invalid,
	"ingress": Ingress,
	"egress":  Egress,
}

var SecurityGroupRuleDirectionNameByTypeMap = map[SecurityGroupRuleDirectionType]string{
	Invalid: "invalid",
	Ingress: "ingress",
	Egress:  "egress",
}

func newGenericRepo(dialector gorm.Dialector) *Repo {
	gormLogger := logger.New(
		slog.NewLogLogger(slog.Default().Handler(), slog.LevelWarn),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		slog.Error("unable to open database", "err", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("unable to access db api", "err", err)
		os.Exit(1)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(time.Minute)

	err = MigrateDB(sqlDB)
	if err != nil {
		slog.Error("unable to apply automigrate", "err", err)
		os.Exit(1)
	}

	return &Repo{
		db: db,
	}
}

func MigrateDB(db *sql.DB) error {
	migrationsSource := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrations.FS,
		Root:       "/",
	}
	migrate.SetTable("migrations")
	if applied, err := migrate.Exec(db, "postgres", migrationsSource, migrate.Up); err != nil {
		return err
	} else {
		slog.Info("applied migrations", "count", applied)
		return nil
	}
}

func NewTestRepo() (*Repo, func()) {
	db, err := pgtest.Start()
	if err != nil {
		panic(fmt.Errorf("cannot init test repo: %w", err))
	}
	pgUser := func(u string) string {
		if u != "" {
			return u + "@"
		}
		return ""
	}
	dsn := fmt.Sprintf("postgresql://%s/%s?host=%s&sslmode=disable", pgUser(db.User), db.Name, db.Host)
	r := newGenericRepo(postgres.Open(dsn))
	r.dsn = dsn

	return r, func() {
		if err := db.Stop(); err != nil {
			log.Fatalf("exiting due error while stopping test repo %s", err.Error())
		}
	}
}

func (r *Repo) CleanTables(ctx context.Context) error {
	for _, model := range Models {
		err := r.db.WithContext(ctx).Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model).Error
		if err != nil {
			return fmt.Errorf("deleting model: %w", err)
		}
	}
	return nil
}
