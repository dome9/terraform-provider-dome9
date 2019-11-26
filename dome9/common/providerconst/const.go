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

// The 16 regions Dome9 manages in AWS cloud account
var AWSRegions = []string{"us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1"}

var CloudVendors = []string{"AWS", "Azure", "GCP"}
