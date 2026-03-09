package vo

import "time"

// ClusterProvider 集群提供程序类型
type ClusterProvider string

const (
	ClusterProviderDocker ClusterProvider = "docker" // Docker容器化集群
)

// WorkerStatusPublic Worker状态信息
type WorkerStatusPublic struct {
	ID          int64             `json:"id"`          // Worker唯一标识符
	Name        string            `json:"name"`        // Worker名称
	WorkerUUID  string            `json:"workerUuid"`  // Worker全局唯一ID
	MachineID   string            `json:"machineId "`  // 机器ID
	Hostname    string            `json:"hostname"`    // 主机名
	IP          string            `json:"ip"`          // Worker IP地址
	Port        int               `json:"port"`        // Worker服务端口
	MetricsPort int               `json:"metricsPort"` // Worker指标端口
	Labels      map[string]string `json:"labels"`      // Worker标签

	// 系统资源字段
	SystemReserved *SystemReserved `json:"systemReserved "` // 系统保留资源
	Status         *WorkerStatus   `json:"status "`
	HasGPU         bool            `json:"hasGpu"` // 是否拥有GPU

	// 节点状态字段
	State         WorkerState  `json:"state"`          // Worker状态
	Unreachable   bool         `json:"unreachable"`    // Worker是否可达
	HeartbeatTime time.Time    `json:"heartbeatTime "` // 最后心跳时间
	StateMessage  string       `json:"stateMessage "`  // 状态描述信息
	Maintenance   *Maintenance `json:"maintenance "`   // 维护模式信息

	//信息变化
	Changed bool `json:"changed"` //节点信息是否改变

	// 其他信息
	Provider ClusterProvider `json:"provider"` // 集群提供程序
}

type WorkerStatus struct {
	*SystemInfo     `json:"systemInfo"` // 系统信息
	GPUDevices      []*GPUDeviceInfo    `json:"gpuDevices"` // GPU设备列表
	LastCollectTime time.Time           `json:"lastCollectTime"`
}

// UtilizationInfo 利用率信息基类
type UtilizationInfo struct {
	Total           int64   `json:"total"`            // 总量（核数 / bytes）
	UtilizationRate float64 `json:"utilizationRate "` // 利用率百分比，范围0-100
}

// CPUInfo CPU信息
type CPUInfo struct {
	UtilizationInfo
}

// MemoryInfo 内存信息
type MemoryInfo struct {
	UtilizationInfo
	Used            int64 `json:"used"`            // 已使用内存
	Allocated       int64 `json:"allocated "`      // 可分配内存
	Cached          int64 `json:"cached "`         // 缓存内存
	IsUnifiedMemory bool  `json:"isUnifiedMemory"` // 是否为统一内存
}

// GPUCoreInfo GPU核心信息
type GPUCoreInfo struct {
	UtilizationInfo
}

// MountPoint 文件系统
type MountPoint struct {
	Name       string `json:"name"`
	MountPoint string `json:"mountPoint"` // 挂载路径
	MountFrom  string `json:"mountFrom"`  // 挂载源
	Total      int64  `json:"total "`     // 总容量（字节）
	Used       int64  `json:"used "`      // 已使用容量
	Free       int64  `json:"free "`      // 可用容量
	Available  int64  `json:"available "` // 实际可用容量
}

// FileSystemInfo 文件系统信息（挂载点列表）
type FileSystemInfo []MountPoint

// SwapInfo 交换分区信息
type SwapInfo struct {
	UtilizationInfo
	Used int64 `json:"used "` // 已使用交换分区
}

// OperatingSystemInfo 操作系统信息
type OperatingSystemInfo struct {
	Name    string `json:"name"`    // 操作系统名称
	Version string `json:"version"` // 操作系统版本
}

// KernelInfo 内核信息
type KernelInfo struct {
	Name         string `json:"name"`         // 内核名称
	Release      string `json:"release"`      // 内核版本号
	Version      string `json:"version"`      // 内核详细版本
	Architecture string `json:"architecture"` // 内核架构
}

// UptimeInfo 系统运行时间信息
type UptimeInfo struct {
	Uptime   float64 `json:"uptime "`  // 运行时间（秒）
	BootTime string  `json:"bootTime"` // 启动时间
}

// SystemInfo Worker状态信息
type SystemInfo struct {
	CPU        *CPUInfo             `json:"cpu "`        // CPU信息
	Memory     *MemoryInfo          `json:"memory "`     // 内存信息
	Swap       *SwapInfo            `json:"swap "`       // 交换分区信息
	FileSystem FileSystemInfo       `json:"filesystem "` // 文件系统信息
	OS         *OperatingSystemInfo `json:"os "`         // 操作系统信息
	Kernel     *KernelInfo          `json:"kernel "`     // 内核信息
	Uptime     *UptimeInfo          `json:"uptime "`     // 运行时间信息
}

// SystemReserved 系统保留资源信息
type SystemReserved struct {
	RAM  int64 `json:"ram "`  // 保留的系统内存
	VRAM int64 `json:"vram "` // 保留的GPU内存
}

// RPCServer RPC服务器信息
type RPCServer struct {
	PID      int `json:"pid "`      // 进程ID
	Port     int `json:"port "`     // 端口号
	GPUIndex int `json:"gpuIndex "` // GPU索引
}

// WorkerState Worker状态枚举
type WorkerState string

const (
	WorkerStateNotReady     WorkerState = "not_ready"    // 不可用（心跳丢失）
	WorkerStateReady        WorkerState = "ready"        // 就绪
	WorkerStateOffline      WorkerState = "offline"      // 离线
	WorkerStateUnreachable  WorkerState = "unreachable"  // 不可达（健康检查失败）
	WorkerStatePending      WorkerState = "pending"      // 待处理
	WorkerStateProvisioning WorkerState = "provisioning" // 配置中
	WorkerStateInitializing WorkerState = "initializing" // 初始化中
	WorkerStateDeleting     WorkerState = "deleting"     // 删除中
	WorkerStateError        WorkerState = "error"        // 错误
	WorkerStateMaintenance  WorkerState = "maintenance"  // 维护模式
)

// GPUVendor GPU厂商枚举
type GPUVendor string

const (
	GPUVendorNvidia GPUVendor = "nvidia" // NVIDIA
	GPUVendorAMD    GPUVendor = "amd"    // AMD
	GPUVendorAscend GPUVendor = "ascend" // Ascend
	GPUVendorOther  GPUVendor = "other"  // 其他

)

// GPUBackend GPU运行时后端类型枚举
type GPUBackend string

const (
	GPUBackendCUDA  GPUBackend = "cuda"
	GPUBackendROCm  GPUBackend = "rocm"
	GPUBackendCANN  GPUBackend = "cann"
	GPUBackendOther GPUBackend = "other"
)

// GPUDeviceInfo GPU设备详细信息
type GPUDeviceInfo struct {
	UUID              string          `json:"uuid "`              // GPU唯一标识符
	Vendor            GPUVendor       `json:"vendor "`            // GPU厂商，如nvidia, amd, ascend等
	Type              GPUBackend      `json:"type "`              // GPU运行时后端类型，如cuda, rocm, cann等
	Manufacturer      string          `json:"manufacturer "`      // GPU制造商
	SerialNumber      string          `json:"serialNumber "`      // GPU序列号
	Index             int             `json:"index "`             // GPU逻辑索引，从0开始计数
	DeviceIndex       int             `json:"deviceIndex"`        // GPU设备索引，对应/dev路径下的设备
	DeviceChipIndex   int             `json:"deviceChipIndex"`    // GPU芯片索引，用于标识同一卡上的不同芯片
	ArchFamily        string          `json:"archFamily "`        // GPU架构系列
	Name              string          `json:"name"`               // GPU型号，如NVIDIA A100-SXM4-40GB等
	DriverVersion     string          `json:"driverVersion "`     // GPU驱动版本
	RuntimeVersion    string          `json:"runtimeVersion "`    // GPU运行时版本，如CUDA版本
	ComputeCapability string          `json:"computeCapability "` // GPU计算能力版本
	Core              *GPUCoreInfo    `json:"core "`              // GPU核心信息
	Memory            *MemoryInfo     `json:"memory "`            // GPU内存信息
	Temperature       float64         `json:"temperature "`       // GPU温度（摄氏度）
	Network           *GPUNetworkInfo `json:"network "`           // GPU网络信息（主要用于Ascend设备）
	Power             *GPUPower       `json:"power "`             // GPU功耗信息
	Health            string          `json:"health"`             // GPU健康状态
	InUse             bool            `json:"inUse"`              // 是否正在使用
}
type GPUPower struct {
	Draw  float64 `json:"draw "`  //当前功耗
	Limit float64 `json:"limit "` //限制功率
}

// GPUNetworkInfo GPU网络信息（主要用于Ascend设备）
type GPUNetworkInfo struct {
	Status  string `json:"status"`  // 网络状态 (up/down)
	Inet    string `json:"inet"`    // IPv4地址
	Netmask string `json:"netmask"` // 子网掩码
	Mac     string `json:"mac"`     // MAC地址
	Gateway string `json:"gateway"` // 默认网关
	Iface   string `json:"iface "`  // 网络接口名称
	MTU     int64  `json:"mtu "`    // 最大传输单元
}

// Maintenance 维护模式信息
type Maintenance struct {
	Enabled bool   `json:"enabled"`  // 是否启用维护模式
	Message string `json:"message "` // 维护消息
}

// WorkerUpdate Worker可更新字段
type WorkerUpdate struct {
	Name        string            `json:"name"`         // Worker名称
	Labels      map[string]string `json:"labels "`      // Worker标签
	Maintenance *Maintenance      `json:"maintenance "` // 维护模式信息
}
