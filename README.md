## Ubuntu 24.04 (Yandex Cloud)

```shell
go test -timeout=30m -p=1 -count=5 ./...
?       github.com/ivagulin/dbmocks/migrations  [no test files]
ok      github.com/ivagulin/dbmocks/pgtestpkg   109.782s
?       github.com/ivagulin/dbmocks/repo        [no test files]
ok      github.com/ivagulin/dbmocks/testcontainers_tmpfs        104.616s
ok      github.com/ivagulin/dbmocks/testcontainerspkg   432.450s
?       github.com/ivagulin/dbmocks/testinfra   [no test files]

go test -timeout=30m -p=1 -count=3 ./...
?       github.com/ivagulin/dbmocks/migrations  [no test files]
ok      github.com/ivagulin/dbmocks/pgtestpkg   65.956s
?       github.com/ivagulin/dbmocks/repo        [no test files]
ok      github.com/ivagulin/dbmocks/testcontainers_tmpfs        63.031s
ok      github.com/ivagulin/dbmocks/testcontainerspkg   262.272s
?       github.com/ivagulin/dbmocks/testinfra   [no test files]
```

### Окружение

Yandex Cloud:

- OS: Ubuntu 24.04
- CPU: 16 CPU 100%, Intel Ice Lake
- RAM: 32 GiB
- DISK: SSD IO 186 GB (56000/11200 IOPS, 220/164 MB/s)

<details>

<summary>Подробнее</summary>

```shell
lsb_release -a
No LSB modules are available.
Distributor ID: Ubuntu
Description:    Ubuntu 24.04.2 LTS
Release:        24.04
Codename:       noble

lscpu
Architecture:             x86_64
  CPU op-mode(s):         32-bit, 64-bit
  Address sizes:          40 bits physical, 57 bits virtual
  Byte Order:             Little Endian
CPU(s):                   16
  On-line CPU(s) list:    0-15
Vendor ID:                GenuineIntel
  BIOS Vendor ID:         QEMU
  Model name:             Intel Xeon Processor (Icelake)
    BIOS Model name:      pc-q35-yc-2.12  CPU @ 2.0GHz
    BIOS CPU family:      1
    CPU family:           6
    Model:                106
    Thread(s) per core:   2
    Core(s) per socket:   8
    Socket(s):            1
    Stepping:             0
    BogoMIPS:             3990.63
    Flags:                fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx pdpe1gb rdtscp lm constant_tsc rep_good nopl xtopology nonstop_tsc cpuid tsc_known_freq pni pclmulqdq ssse3 fma 
                          cx16 pcid sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx f16c rdrand hypervisor lahf_lm abm 3dnowprefetch cpuid_fault ssbd ibrs ibpb stibp ibrs_enhanced fsgsbase bmi1 avx2 smep bmi2 erms invpcid avx512f a
                          vx512dq rdseed adx smap avx512ifma clflushopt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 wbnoinvd arat avx512vbmi umip pku ospke avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bitalg avx512_vpopcntd
                          q la57 rdpid fsrm md_clear arch_capabilities
Virtualization features:  
  Hypervisor vendor:      KVM
  Virtualization type:    full
Caches (sum of all):      
  L1d:                    512 KiB (16 instances)
  L1i:                    512 KiB (16 instances)
  L2:                     32 MiB (8 instances)
  L3:                     16 MiB (1 instance)
NUMA:                     
  NUMA node(s):           1
  NUMA node0 CPU(s):      0-15
Vulnerabilities:          
  Gather data sampling:   Not affected
  Itlb multihit:          Not affected
  L1tf:                   Not affected
  Mds:                    Not affected
  Meltdown:               Not affected
  Mmio stale data:        Mitigation; Clear CPU buffers; SMT Host state unknown
  Reg file data sampling: Not affected
  Retbleed:               Not affected
  Spec rstack overflow:   Not affected
  Spec store bypass:      Mitigation; Speculative Store Bypass disabled via prctl
  Spectre v1:             Mitigation; usercopy/swapgs barriers and __user pointer sanitization
  Spectre v2:             Mitigation; Enhanced / Automatic IBRS; IBPB conditional; RSB filling; PBRSB-eIBRS SW sequence; BHI SW loop, KVM SW loop
  Srbds:                  Not affected
  Tsx async abort:        Not affected
  
free -h
               total        used        free      shared  buff/cache   available
Mem:            31Gi       1.3Gi        26Gi       182Mi       4.7Gi        30Gi
Swap:             0B          0B          0B

go version
go version go1.23.10 linux/amd64

docker info
Client: Docker Engine - Community
 Version:    28.3.1
 Context:    default
 Debug Mode: false
 Plugins:
  buildx: Docker Buildx (Docker Inc.)
    Version:  v0.25.0
    Path:     /usr/libexec/docker/cli-plugins/docker-buildx
  compose: Docker Compose (Docker Inc.)
    Version:  v2.38.1
    Path:     /usr/libexec/docker/cli-plugins/docker-compose

Server:
 Containers: 0
  Running: 0
  Paused: 0
  Stopped: 0
 Images: 2
 Server Version: 28.3.1
 Storage Driver: overlay2
  Backing Filesystem: extfs
  Supports d_type: true
  Using metacopy: false
  Native Overlay Diff: true
  userxattr: false
 Logging Driver: json-file
 Cgroup Driver: systemd
 Cgroup Version: 2
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local splunk syslog
 CDI spec directories:
  /etc/cdi
  /var/run/cdi
 Swarm: inactive
 Runtimes: io.containerd.runc.v2 runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 05044ec0a9a75232cad458027ca83437aae3f4da
 runc version: v1.2.5-0-g59923ef
 init version: de40ad0
 Security Options:
  apparmor
  seccomp
   Profile: builtin
  cgroupns
 Kernel Version: 6.8.0-62-generic
 Operating System: Ubuntu 24.04.2 LTS
 OSType: linux
 Architecture: x86_64
 CPUs: 16
 Total Memory: 31.34GiB
 Name: compute-vm-16-32-186-ssd-io-m3-1751676682773
 ID: 9bac4856-d65a-42b6-9048-367a314948d6
 Docker Root Dir: /var/lib/docker
 Debug Mode: false
 Experimental: false
 Insecure Registries:
  ::1/128
  127.0.0.0/8
 Live Restore Enabled: false
 
psql --version
psql (PostgreSQL) 16.9 (Ubuntu 16.9-1.pgdg24.04+1)

docker images | grep postgres
postgres              16        f465f4fe2ba0   4 weeks ago    436MB
```

</details>

<details>

<summary>Настройка окружения</summary>

```shell
sudo apt update
sudo apt install -y \
  git \
  wget \
  curl \
  ca-certificates \
  gnupg \
  lsb-release \
  apt-transport-https

#go
rm -rf /usr/local/go
wget https://go.dev/dl/go1.23.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.23.10.linux-amd64.tar.gz

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile

go version

#docker
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done

# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

docker pull postgres:16

#postgres
sudo apt install -y postgresql-common curl ca-certificates
sudo /usr/share/postgresql-common/pgdg/apt.postgresql.org.sh

sudo apt update
sudo apt install -y postgresql-16 postgresql-client-16 libpq-dev

psql --version
```

</details>

## MacOS

```shell
go test -timeout=30m -p=1 -count=3 ./...
?       github.com/ivagulin/dbmocks/migrations  [no test files]
ok      github.com/ivagulin/dbmocks/pgtestpkg   179.647s
?       github.com/ivagulin/dbmocks/repo        [no test files]
ok      github.com/ivagulin/dbmocks/testcontainers_tmpfs        100.401s
ok      github.com/ivagulin/dbmocks/testcontainerspkg   112.284s
?       github.com/ivagulin/dbmocks/testinfra   [no test files]

go test -timeout=30m -p=1 -count=3 ./...
?       github.com/ivagulin/dbmocks/migrations  [no test files]
ok      github.com/ivagulin/dbmocks/pgtestpkg   179.349s
?       github.com/ivagulin/dbmocks/repo        [no test files]
ok      github.com/ivagulin/dbmocks/testcontainers_tmpfs        106.758s
ok      github.com/ivagulin/dbmocks/testcontainerspkg   117.221s
?       github.com/ivagulin/dbmocks/testinfra   [no test files]
```

### Окружение

Apple MacBook Pro 14, M2 Pro, 32 GiB, 512 GB, macOS 13.7.4 (22H420)

<details>

<summary>Подробнее</summary>

```shell
go version
go version go1.23.5 darwin/arm64

docker version
Client:
 Version:           28.0.1
 API version:       1.48
 Go version:        go1.23.6
 Git commit:        068a01e
 Built:             Wed Feb 26 10:38:16 2025
 OS/Arch:           darwin/arm64
 Context:           desktop-linux

Server: Docker Desktop 4.39.0 (184744)
 Engine:
  Version:          28.0.1
  API version:      1.48 (minimum version 1.24)
  Go version:       go1.23.6
  Git commit:       bbd0a17
  Built:            Wed Feb 26 10:40:57 2025
  OS/Arch:          linux/arm64
  Experimental:     false
 containerd:
  Version:          1.7.25
  GitCommit:        bcc810d6b9066471b0b6fa75f557a15a1cbf31bb
 runc:
  Version:          1.2.4
  GitCommit:        v1.2.4-0-g6c52b3f
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0
  
postgres --version
postgres (PostgreSQL) 16.9 (Homebrew)
```

</details>
