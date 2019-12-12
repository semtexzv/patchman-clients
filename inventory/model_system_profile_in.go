/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// SystemProfileIn Representation of the system profile fields
type SystemProfileIn struct {
	Arch string `json:"arch,omitempty"`
	BiosReleaseDate string `json:"bios_release_date,omitempty"`
	BiosVendor string `json:"bios_vendor,omitempty"`
	BiosVersion string `json:"bios_version,omitempty"`
	CloudProvider string `json:"cloud_provider,omitempty"`
	CoresPerSocket int32 `json:"cores_per_socket,omitempty"`
	CpuFlags []string `json:"cpu_flags,omitempty"`
	DiskDevices []DiskDevice `json:"disk_devices,omitempty"`
	DnfModules []DnfModule `json:"dnf_modules,omitempty"`
	EnabledServices []string `json:"enabled_services,omitempty"`
	InfrastructureType string `json:"infrastructure_type,omitempty"`
	InfrastructureVendor string `json:"infrastructure_vendor,omitempty"`
	InsightsClientVersion string `json:"insights_client_version,omitempty"`
	InsightsEggVersion string `json:"insights_egg_version,omitempty"`
	InstalledPackages []string `json:"installed_packages,omitempty"`
	InstalledProducts []InstalledProduct `json:"installed_products,omitempty"`
	InstalledServices []string `json:"installed_services,omitempty"`
	KatelloAgentRunning bool `json:"katello_agent_running,omitempty"`
	KernelModules []string `json:"kernel_modules,omitempty"`
	LastBootTime string `json:"last_boot_time,omitempty"`
	NetworkInterfaces []NetworkInterface `json:"network_interfaces,omitempty"`
	NumberOfCpus int32 `json:"number_of_cpus,omitempty"`
	NumberOfSockets int32 `json:"number_of_sockets,omitempty"`
	OsKernelVersion string `json:"os_kernel_version,omitempty"`
	OsRelease string `json:"os_release,omitempty"`
	RunningProcesses []string `json:"running_processes,omitempty"`
	SatelliteManaged bool `json:"satellite_managed,omitempty"`
	SubscriptionAutoAttach string `json:"subscription_auto_attach,omitempty"`
	SubscriptionStatus string `json:"subscription_status,omitempty"`
	SystemMemoryBytes int32 `json:"system_memory_bytes,omitempty"`
	YumRepos []YumRepo `json:"yum_repos,omitempty"`
}
