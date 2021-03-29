package config

type Config struct {
	MJ_APIKEY_PUBLIC  string
	MJ_APIKEY_PRIVATE string
	From              string
	Name              string
	Username          string
	Password          string
}

// LoadConfig to initialize the config parameters
func LoadConfig() Config {
	var cfg Config

	// fetch environment values for MJ_APIKEY_PRIVATE, MJ_APIKEY_PUBLIC if empty, use below for testing purpose
	// from security perspective, MJ_APIKEY_PRIVATE should not be present in the repo
	// instead can be fetched as environment variable or some other secured place
	cfg.MJ_APIKEY_PRIVATE = "44e6855e189d2db06f791b868dc3113d"
	cfg.MJ_APIKEY_PUBLIC = "41d2e0f3d61a6182c53dd52fd14c0b26"

	cfg.From = "saurabhagr.developer@gmail.com"
	cfg.Name = "Saurabh Agarwal"

	// config info for API basic authentication
	// Note: below info is just kept for demo/test purpose
	// from security perspective, in realtime,  these can be fetched from environment or some secured place
	cfg.Username = "abc"
	cfg.Password = "123"

	return cfg
}
