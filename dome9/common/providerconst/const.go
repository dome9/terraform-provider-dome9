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
const (
	AWS                = "1"
	Azure              = "2"
	GCP                = "3"
	OrganizationalUnit = "12"
)

const (
	US_EAST_1           = "0"
	US_WEST_1           = "1"
	EU_WEST_1           = "2"
	AP_SOUTHEAST_1      = "3"
	AP_NORTHEAST_1      = "4"
	US_WEST_2           = "5"
	SA_EAST_1           = "6"
	AZ_1_REGION_A_GEO_1 = "7"
	AZ_2_REGION_A_GEO_1 = "8"
	AZ_3_REGION_A_GEO_1 = "9"
	AP_SOUTHEAST_2      = "10"
	MELLANOX_REGION     = "11"
	US_GOV_WEST_1       = "12"
	EU_CENTRAL_1        = "13"
	AP_NORTHEAST_2      = "14"
	AP_SOUTH_1          = "15"
	US_EAST_2           = "16"
	CA_CENTRAL_1        = "17"
	EU_WEST_2           = "18"
	EU_WEST_3           = "19"
	EU_NORTH_1          = "20"
	CN_NORTH_1          = "21"
	CN_NORTHWEST_1      = "22"
	US_GOV_EAST_1       = "23"
	AP_EAST_1           = "24"
	ME_SOUTH_1          = "25"
	AF_SOUTH_1          = "26"
	EU_SOUTH_1          = "27"
	AP_NORTHEAST_3      = "28"
)

// AzureSecurityGroupRegions const
var AzureSecurityGroupRegions = []string{"centralus", "eastus", "eastus2", "usgovlowa", "usgovvirginia", "northcentralus", "southcentralus", "westus", "westus2", "westcentralus", "northeurope", "westeurope", "eastasia", "southeastasia", "japaneast", "japanwest", "brazilsouth", "australiaeast", "australiasoutheast", "centralindia", "southindia", "westindia", "chinaeast", "chinanorth", "canadacentral", "canadaeast", "germanycentral", "germanynortheast", "koreacentral", "uksouth", "ukwest", "koreasouth"}
var AzureSecurityGroupAccess = []string{"Allow", "Deny"}
var AzureSecurityGroupProtocol = []string{"UDP", "TCP", "ANY"}
var AzureSecurityGroupSourceScopeTypes = []string{"CIDR", "IPList", "Tag"}

// The 21 regions Dome9 manages in AWS cloud account
var AWSRegions = []string{"us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1", "ap_east_1", "me_south_1", "af_south_1", "eu_south_1", "ap_northeast_3"}
// The 2 regions Dome9 manages in AWSGOV cloud account
var AWSGOVRegions = []string{"us_gov_west_1", "us_gov_east_1"}
// The 23 regions Dome9 manages in AWS & AWSGOV cloud account
var AllAWSRegions = append(AWSGOVRegions, AWSRegions...)
// AWSRegions The 25 regions Dome9 manages in AWS cloud account (3 for awsGov & 2 for awsChina)
var AWSRegions = []string{"cn_northwest_1", "cn_north_1","us_gov_west_1", "us_gov_east_1", "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1", "ap_east_1", "me_south_1", "af_south_1", "eu_south_1"}
var CloudVendors = []string{"aws", "azure", "google"}
var ProtocolTypes = []string{"ALL", "HOPOPT", "ICMP", "IGMP", "GGP", "IPV4", "ST", "TCP", "CBT", "EGP", "IGP", "BBN_RCC_MON", "NVP2", "PUP", "ARGUS", "EMCON", "XNET", "CHAOS", "UDP", "MUX", "DCN_MEAS", "HMP", "PRM", "XNS_IDP", "TRUNK1", "TRUNK2", "LEAF1", "LEAF2", "RDP", "IRTP", "ISO_TP4", "NETBLT", "MFE_NSP", "MERIT_INP", "DCCP", "ThreePC", "IDPR", "XTP", "DDP", "IDPR_CMTP", "TPplusplus", "IL", "IPV6", "SDRP", "IPV6_ROUTE", "IPV6_FRAG", "IDRP", "RSVP", "GRE", "DSR", "BNA", "ESP", "AH", "I_NLSP", "SWIPE", "NARP", "MOBILE", "TLSP", "SKIP", "ICMPV6", "IPV6_NONXT", "IPV6_OPTS", "CFTP", "SAT_EXPAK", "KRYPTOLAN", "RVD", "IPPC", "SAT_MON", "VISA", "IPCV", "CPNX", "CPHB", "WSN", "PVP", "BR_SAT_MON", "SUN_ND", "WB_MON", "WB_EXPAK", "ISO_IP", "VMTP", "SECURE_VMTP", "VINES", "TTP", "NSFNET_IGP", "DGP", "TCF", "EIGRP", "OSPFIGP", "SPRITE_RPC", "LARP", "MTP", "AX25", "IPIP", "MICP", "SCC_SP", "ETHERIP", "ENCAP", "GMTP", "IFMP", "PNNI", "PIM", "ARIS", "SCPS", "QNX", "AN", "IPCOMP", "SNP", "COMPAQ_PEER", "IPX_IN_IP", "VRRP", "PGM", "L2TP", "DDX", "IATP", "STP", "SRP", "UTI", "SMP", "SM", "PTP", "ISIS", "FIRE", "CRTP", "CRUDP", "SSCOPMCE", "IPLT", "SPS", "PIPE", "SCTP", "FC", "RSVP_E2E_IGNORE", "MOBILITY_HEADER", "UDPLITE", "MPLS_IN_IP", "MANET", "HIP", "SHIM6", "WESP", "ROHC"}
var OperationMode = []string{"Read", "Manage"}
var SRLTypes = []string{"AWS", "Azure", "GCP", "OrganizationalUnit"}

var IAMEntityProtectType = []string{IAMSafeEntityTypeUser, IAMSafeEntityTypeRole}
var IAMEntityProtectionMode = []string{IAMSafeEntityProtect, IAMSafeEntityProtectWithElevation}

//PermissionTrafficOptions Used in role permission construction
var PermissionTrafficOptions = []string{"All Services", "All Traffic"}
var SRLStructure = []string{"type", "main_id", "rg", "region", "sg", "security_group_id", "traffic"}

//SRlType SRL construction variables
var SRlType = map[string]string{
	"AWS":                AWS,
	"Azure":              Azure,
	"GCP":                GCP,
	"OrganizationalUnit": OrganizationalUnit,
}

var AWSRegionsEnum = map[string]string{
	"us_east_1":           US_EAST_1,
	"us_west_1":           US_WEST_1,
	"eu_west_1":           EU_WEST_1,
	"ap_southeast_1":      AP_SOUTHEAST_1,
	"ap_northeast_1":      AP_NORTHEAST_1,
	"us_west_2":           US_WEST_2,
	"sa_east_1":           SA_EAST_1,
	"az_1_region_a_geo_1": AZ_1_REGION_A_GEO_1,
	"az_2_region_a_geo_1": AZ_2_REGION_A_GEO_1,
	"az_3_region_a_geo_1": AZ_3_REGION_A_GEO_1,
	"ap_southeast_2":      AP_SOUTHEAST_2,
	"mellanox_region":     MELLANOX_REGION,
	"us_gov_west_1":       US_GOV_WEST_1,
	"eu_central_1":        EU_CENTRAL_1,
	"ap_northeast_2":      AP_NORTHEAST_2,
	"ap_south_1":          AP_SOUTH_1,
	"us_east_2":           US_EAST_2,
	"ca_central_1":        CA_CENTRAL_1,
	"eu_west_2":           EU_WEST_2,
	"eu_west_3":           EU_WEST_3,
	"eu_north_1":          EU_NORTH_1,
	"cn_north_1":          CN_NORTH_1,
	"cn_northwest_1":      CN_NORTHWEST_1,
	"us_gov_east_1":       US_GOV_EAST_1,
	"ap_east_1":           AP_EAST_1,
	"me_south_1":          ME_SOUTH_1,
	"af_south_1":          AF_SOUTH_1,
	"eu_south_1":          EU_SOUTH_1,
	"ap_northeast_3":      AP_NORTHEAST_3,
}

var PermissionTrafficType = map[string]string{
	"All Services": "",
	"All Traffic":  "-1",
}
