package config

type AuthConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func LoadAuthConfig() AuthConfig {
	return AuthConfig{
		AccessSecret:  "1122334455",
		RefreshSecret: "445566778899",
	}
}
