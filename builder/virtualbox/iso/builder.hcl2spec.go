// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package iso

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug               *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	HTTPDir                   *string           `mapstructure:"http_directory" cty:"http_directory" hcl:"http_directory"`
	HTTPPortMin               *int              `mapstructure:"http_port_min" cty:"http_port_min" hcl:"http_port_min"`
	HTTPPortMax               *int              `mapstructure:"http_port_max" cty:"http_port_max" hcl:"http_port_max"`
	HTTPAddress               *string           `mapstructure:"http_bind_address" cty:"http_bind_address" hcl:"http_bind_address"`
	ISOChecksum               *string           `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum" hcl:"iso_checksum"`
	RawSingleISOUrl           *string           `mapstructure:"iso_url" required:"true" cty:"iso_url" hcl:"iso_url"`
	ISOUrls                   []string          `mapstructure:"iso_urls" cty:"iso_urls" hcl:"iso_urls"`
	TargetPath                *string           `mapstructure:"iso_target_path" cty:"iso_target_path" hcl:"iso_target_path"`
	TargetExtension           *string           `mapstructure:"iso_target_extension" cty:"iso_target_extension" hcl:"iso_target_extension"`
	FloppyFiles               []string          `mapstructure:"floppy_files" cty:"floppy_files" hcl:"floppy_files"`
	FloppyDirectories         []string          `mapstructure:"floppy_dirs" cty:"floppy_dirs" hcl:"floppy_dirs"`
	FloppyLabel               *string           `mapstructure:"floppy_label" cty:"floppy_label" hcl:"floppy_label"`
	BootGroupInterval         *string           `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval" hcl:"boot_keygroup_interval"`
	BootWait                  *string           `mapstructure:"boot_wait" cty:"boot_wait" hcl:"boot_wait"`
	BootCommand               []string          `mapstructure:"boot_command" cty:"boot_command" hcl:"boot_command"`
	Format                    *string           `mapstructure:"format" required:"false" cty:"format" hcl:"format"`
	ExportOpts                []string          `mapstructure:"export_opts" required:"false" cty:"export_opts" hcl:"export_opts"`
	OutputDir                 *string           `mapstructure:"output_directory" required:"false" cty:"output_directory" hcl:"output_directory"`
	OutputFilename            *string           `mapstructure:"output_filename" required:"false" cty:"output_filename" hcl:"output_filename"`
	Headless                  *bool             `mapstructure:"headless" required:"false" cty:"headless" hcl:"headless"`
	VRDPBindAddress           *string           `mapstructure:"vrdp_bind_address" required:"false" cty:"vrdp_bind_address" hcl:"vrdp_bind_address"`
	VRDPPortMin               *int              `mapstructure:"vrdp_port_min" required:"false" cty:"vrdp_port_min" hcl:"vrdp_port_min"`
	VRDPPortMax               *int              `mapstructure:"vrdp_port_max" cty:"vrdp_port_max" hcl:"vrdp_port_max"`
	ShutdownCommand           *string           `mapstructure:"shutdown_command" required:"false" cty:"shutdown_command" hcl:"shutdown_command"`
	ShutdownTimeout           *string           `mapstructure:"shutdown_timeout" required:"false" cty:"shutdown_timeout" hcl:"shutdown_timeout"`
	PostShutdownDelay         *string           `mapstructure:"post_shutdown_delay" required:"false" cty:"post_shutdown_delay" hcl:"post_shutdown_delay"`
	DisableShutdown           *bool             `mapstructure:"disable_shutdown" required:"false" cty:"disable_shutdown" hcl:"disable_shutdown"`
	ACPIShutdown              *bool             `mapstructure:"acpi_shutdown" required:"false" cty:"acpi_shutdown" hcl:"acpi_shutdown"`
	Type                      *string           `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string           `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string           `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int              `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string           `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string           `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string           `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string           `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHCiphers                []string          `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool             `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHPrivateKeyFile         *string           `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHPty                    *bool             `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string           `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string           `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool             `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool             `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int              `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string           `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int              `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool             `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string           `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string           `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool             `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string           `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHFileTransferMethod     *string           `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string           `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int              `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string           `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string           `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string           `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string           `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string          `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string          `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte            `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte            `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string           `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string           `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string           `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool             `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int              `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string           `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool             `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool             `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool             `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	HostPortMin               *int              `mapstructure:"host_port_min" required:"false" cty:"host_port_min" hcl:"host_port_min"`
	HostPortMax               *int              `mapstructure:"host_port_max" required:"false" cty:"host_port_max" hcl:"host_port_max"`
	SkipNatMapping            *bool             `mapstructure:"skip_nat_mapping" required:"false" cty:"skip_nat_mapping" hcl:"skip_nat_mapping"`
	SSHHostPortMin            *int              `mapstructure:"ssh_host_port_min" required:"false" cty:"ssh_host_port_min" hcl:"ssh_host_port_min"`
	SSHHostPortMax            *int              `mapstructure:"ssh_host_port_max" cty:"ssh_host_port_max" hcl:"ssh_host_port_max"`
	SSHSkipNatMapping         *bool             `mapstructure:"ssh_skip_nat_mapping" required:"false" cty:"ssh_skip_nat_mapping" hcl:"ssh_skip_nat_mapping"`
	CpuCount                  *int              `mapstructure:"cpus" required:"false" cty:"cpus" hcl:"cpus"`
	MemorySize                *int              `mapstructure:"memory" required:"false" cty:"memory" hcl:"memory"`
	Sound                     *string           `mapstructure:"sound" required:"false" cty:"sound" hcl:"sound"`
	USB                       *bool             `mapstructure:"usb" required:"false" cty:"usb" hcl:"usb"`
	VBoxManage                [][]string        `mapstructure:"vboxmanage" required:"false" cty:"vboxmanage" hcl:"vboxmanage"`
	VBoxManagePost            [][]string        `mapstructure:"vboxmanage_post" required:"false" cty:"vboxmanage_post" hcl:"vboxmanage_post"`
	VBoxVersionFile           *string           `mapstructure:"virtualbox_version_file" required:"false" cty:"virtualbox_version_file" hcl:"virtualbox_version_file"`
	BundleISO                 *bool             `mapstructure:"bundle_iso" required:"false" cty:"bundle_iso" hcl:"bundle_iso"`
	GuestAdditionsMode        *string           `mapstructure:"guest_additions_mode" required:"false" cty:"guest_additions_mode" hcl:"guest_additions_mode"`
	DiskSize                  *uint             `mapstructure:"disk_size" required:"false" cty:"disk_size" hcl:"disk_size"`
	GuestAdditionsPath        *string           `mapstructure:"guest_additions_path" required:"false" cty:"guest_additions_path" hcl:"guest_additions_path"`
	GuestAdditionsSHA256      *string           `mapstructure:"guest_additions_sha256" required:"false" cty:"guest_additions_sha256" hcl:"guest_additions_sha256"`
	GuestAdditionsURL         *string           `mapstructure:"guest_additions_url" required:"false" cty:"guest_additions_url" hcl:"guest_additions_url"`
	GuestAdditionsInterface   *string           `mapstructure:"guest_additions_interface" required:"false" cty:"guest_additions_interface" hcl:"guest_additions_interface"`
	GuestOSType               *string           `mapstructure:"guest_os_type" required:"false" cty:"guest_os_type" hcl:"guest_os_type"`
	HardDriveDiscard          *bool             `mapstructure:"hard_drive_discard" required:"false" cty:"hard_drive_discard" hcl:"hard_drive_discard"`
	HardDriveInterface        *string           `mapstructure:"hard_drive_interface" required:"false" cty:"hard_drive_interface" hcl:"hard_drive_interface"`
	SATAPortCount             *int              `mapstructure:"sata_port_count" required:"false" cty:"sata_port_count" hcl:"sata_port_count"`
	NVMePortCount             *int              `mapstructure:"nvme_port_count" required:"false" cty:"nvme_port_count" hcl:"nvme_port_count"`
	HardDriveNonrotational    *bool             `mapstructure:"hard_drive_nonrotational" required:"false" cty:"hard_drive_nonrotational" hcl:"hard_drive_nonrotational"`
	ISOInterface              *string           `mapstructure:"iso_interface" required:"false" cty:"iso_interface" hcl:"iso_interface"`
	KeepRegistered            *bool             `mapstructure:"keep_registered" required:"false" cty:"keep_registered" hcl:"keep_registered"`
	SkipExport                *bool             `mapstructure:"skip_export" required:"false" cty:"skip_export" hcl:"skip_export"`
	VMName                    *string           `mapstructure:"vm_name" required:"false" cty:"vm_name" hcl:"vm_name"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"http_bind_address":            &hcldec.AttrSpec{Name: "http_bind_address", Type: cty.String, Required: false},
		"iso_checksum":                 &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_url":                      &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                     &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":              &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":         &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"floppy_files":                 &hcldec.AttrSpec{Name: "floppy_files", Type: cty.List(cty.String), Required: false},
		"floppy_dirs":                  &hcldec.AttrSpec{Name: "floppy_dirs", Type: cty.List(cty.String), Required: false},
		"floppy_label":                 &hcldec.AttrSpec{Name: "floppy_label", Type: cty.String, Required: false},
		"boot_keygroup_interval":       &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"format":                       &hcldec.AttrSpec{Name: "format", Type: cty.String, Required: false},
		"export_opts":                  &hcldec.AttrSpec{Name: "export_opts", Type: cty.List(cty.String), Required: false},
		"output_directory":             &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"output_filename":              &hcldec.AttrSpec{Name: "output_filename", Type: cty.String, Required: false},
		"headless":                     &hcldec.AttrSpec{Name: "headless", Type: cty.Bool, Required: false},
		"vrdp_bind_address":            &hcldec.AttrSpec{Name: "vrdp_bind_address", Type: cty.String, Required: false},
		"vrdp_port_min":                &hcldec.AttrSpec{Name: "vrdp_port_min", Type: cty.Number, Required: false},
		"vrdp_port_max":                &hcldec.AttrSpec{Name: "vrdp_port_max", Type: cty.Number, Required: false},
		"shutdown_command":             &hcldec.AttrSpec{Name: "shutdown_command", Type: cty.String, Required: false},
		"shutdown_timeout":             &hcldec.AttrSpec{Name: "shutdown_timeout", Type: cty.String, Required: false},
		"post_shutdown_delay":          &hcldec.AttrSpec{Name: "post_shutdown_delay", Type: cty.String, Required: false},
		"disable_shutdown":             &hcldec.AttrSpec{Name: "disable_shutdown", Type: cty.Bool, Required: false},
		"acpi_shutdown":                &hcldec.AttrSpec{Name: "acpi_shutdown", Type: cty.Bool, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"host_port_min":                &hcldec.AttrSpec{Name: "host_port_min", Type: cty.Number, Required: false},
		"host_port_max":                &hcldec.AttrSpec{Name: "host_port_max", Type: cty.Number, Required: false},
		"skip_nat_mapping":             &hcldec.AttrSpec{Name: "skip_nat_mapping", Type: cty.Bool, Required: false},
		"ssh_host_port_min":            &hcldec.AttrSpec{Name: "ssh_host_port_min", Type: cty.Number, Required: false},
		"ssh_host_port_max":            &hcldec.AttrSpec{Name: "ssh_host_port_max", Type: cty.Number, Required: false},
		"ssh_skip_nat_mapping":         &hcldec.AttrSpec{Name: "ssh_skip_nat_mapping", Type: cty.Bool, Required: false},
		"cpus":                         &hcldec.AttrSpec{Name: "cpus", Type: cty.Number, Required: false},
		"memory":                       &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"sound":                        &hcldec.AttrSpec{Name: "sound", Type: cty.String, Required: false},
		"usb":                          &hcldec.AttrSpec{Name: "usb", Type: cty.Bool, Required: false},
		"vboxmanage":                   &hcldec.AttrSpec{Name: "vboxmanage", Type: cty.List(cty.List(cty.String)), Required: false},
		"vboxmanage_post":              &hcldec.AttrSpec{Name: "vboxmanage_post", Type: cty.List(cty.List(cty.String)), Required: false},
		"virtualbox_version_file":      &hcldec.AttrSpec{Name: "virtualbox_version_file", Type: cty.String, Required: false},
		"bundle_iso":                   &hcldec.AttrSpec{Name: "bundle_iso", Type: cty.Bool, Required: false},
		"guest_additions_mode":         &hcldec.AttrSpec{Name: "guest_additions_mode", Type: cty.String, Required: false},
		"disk_size":                    &hcldec.AttrSpec{Name: "disk_size", Type: cty.Number, Required: false},
		"guest_additions_path":         &hcldec.AttrSpec{Name: "guest_additions_path", Type: cty.String, Required: false},
		"guest_additions_sha256":       &hcldec.AttrSpec{Name: "guest_additions_sha256", Type: cty.String, Required: false},
		"guest_additions_url":          &hcldec.AttrSpec{Name: "guest_additions_url", Type: cty.String, Required: false},
		"guest_additions_interface":    &hcldec.AttrSpec{Name: "guest_additions_interface", Type: cty.String, Required: false},
		"guest_os_type":                &hcldec.AttrSpec{Name: "guest_os_type", Type: cty.String, Required: false},
		"hard_drive_discard":           &hcldec.AttrSpec{Name: "hard_drive_discard", Type: cty.Bool, Required: false},
		"hard_drive_interface":         &hcldec.AttrSpec{Name: "hard_drive_interface", Type: cty.String, Required: false},
		"sata_port_count":              &hcldec.AttrSpec{Name: "sata_port_count", Type: cty.Number, Required: false},
		"nvme_port_count":              &hcldec.AttrSpec{Name: "nvme_port_count", Type: cty.Number, Required: false},
		"hard_drive_nonrotational":     &hcldec.AttrSpec{Name: "hard_drive_nonrotational", Type: cty.Bool, Required: false},
		"iso_interface":                &hcldec.AttrSpec{Name: "iso_interface", Type: cty.String, Required: false},
		"keep_registered":              &hcldec.AttrSpec{Name: "keep_registered", Type: cty.Bool, Required: false},
		"skip_export":                  &hcldec.AttrSpec{Name: "skip_export", Type: cty.Bool, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
	}
	return s
}
