package providerconst

// Provider environment variable
const (
	ProviderAccessIDEnvVariable  = "DOME9_ACCESS_ID"
	ProviderSecretKeyEnvVariable = "DOME9_SECRET_KEY"
)

// SDK parameters names
const (
	ProviderAccessID  = "dome9_access_id"
	ProviderSecretKey = "dome9_secret_key"
	ProviderBaseURL   = "base_url"
)

// GCP onboarding
const (
	GCPCloudAccountType                    = "service_account"
	GCPCloudAccountAuthUri                 = "https://accounts.google.com/o/oauth2/auth"
	GCPCloudAccountTokenUri                = "https://oauth2.googleapis.com/token"
	GCPCloudAccountAuthProviderX509CertUrl = "https://www.googleapis.com/oauth2/v1/certs"
)

// AWS security group protection mode
const (
	FullManage = "FullManage"
	ReadOnly   = "ReadOnly"
)

// IAM entity
const (
	IAMSafeEntityTypeUser             = "User"
	IAMSafeEntityTypeRole             = "Role"
	IAMSafeEntityProtect              = "Protect"
	IAMSafeEntityProtectWithElevation = "ProtectWithElevation"
)

// Azure security group
var AzureSecurityGroupRegions = []string{"centralus", "eastus", "eastus2", "usgovlowa", "usgovvirginia", "northcentralus", "southcentralus", "westus", "westus2", "westcentralus", "northeurope", "westeurope", "eastasia", "southeastasia", "japaneast", "japanwest", "brazilsouth", "australiaeast", "australiasoutheast", "centralindia", "southindia", "westindia", "chinaeast", "chinanorth", "canadacentral", "canadaeast", "germanycentral", "germanynortheast", "koreacentral", "uksouth", "ukwest", "koreasouth"}
var AzureSecurityGroupAccess = []string{"Allow", "Deny"}
var AzureSecurityGroupProtocol = []string{"UDP", "TCP", "ANY"}
var AzureSecurityGroupSourceScopeTypes = []string{"CIDR", "IPList", "Tag"}

// The 16 regions Dome9 manages in AWS cloud account
var AWSRegions = []string{"us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1"}
var CloudVendors = []string{"AWS", "Azure", "GCP"}
var ProtocolTypes = []string{"ALL", "HOPOPT", "ICMP", "IGMP", "GGP", "IPV4", "ST", "TCP", "CBT", "EGP", "IGP", "BBN_RCC_MON", "NVP2", "PUP", "ARGUS", "EMCON", "XNET", "CHAOS", "UDP", "MUX", "DCN_MEAS", "HMP", "PRM", "XNS_IDP", "TRUNK1", "TRUNK2", "LEAF1", "LEAF2", "RDP", "IRTP", "ISO_TP4", "NETBLT", "MFE_NSP", "MERIT_INP", "DCCP", "ThreePC", "IDPR", "XTP", "DDP", "IDPR_CMTP", "TPplusplus", "IL", "IPV6", "SDRP", "IPV6_ROUTE", "IPV6_FRAG", "IDRP", "RSVP", "GRE", "DSR", "BNA", "ESP", "AH", "I_NLSP", "SWIPE", "NARP", "MOBILE", "TLSP", "SKIP", "ICMPV6", "IPV6_NONXT", "IPV6_OPTS", "CFTP", "SAT_EXPAK", "KRYPTOLAN", "RVD", "IPPC", "SAT_MON", "VISA", "IPCV", "CPNX", "CPHB", "WSN", "PVP", "BR_SAT_MON", "SUN_ND", "WB_MON", "WB_EXPAK", "ISO_IP", "VMTP", "SECURE_VMTP", "VINES", "TTP", "NSFNET_IGP", "DGP", "TCF", "EIGRP", "OSPFIGP", "SPRITE_RPC", "LARP", "MTP", "AX25", "IPIP", "MICP", "SCC_SP", "ETHERIP", "ENCAP", "GMTP", "IFMP", "PNNI", "PIM", "ARIS", "SCPS", "QNX", "AN", "IPCOMP", "SNP", "COMPAQ_PEER", "IPX_IN_IP", "VRRP", "PGM", "L2TP", "DDX", "IATP", "STP", "SRP", "UTI", "SMP", "SM", "PTP", "ISIS", "FIRE", "CRTP", "CRUDP", "SSCOPMCE", "IPLT", "SPS", "PIPE", "SCTP", "FC", "RSVP_E2E_IGNORE", "MOBILITY_HEADER", "UDPLITE", "MPLS_IN_IP", "MANET", "HIP", "SHIM6", "WESP", "ROHC"}
var OperationMode = []string{"Read", "Manage"}

var IAMEntityProtectType = []string{IAMSafeEntityTypeUser, IAMSafeEntityTypeRole}
var IAMEntityProtectionMode = []string{IAMSafeEntityProtect, IAMSafeEntityProtectWithElevation}
