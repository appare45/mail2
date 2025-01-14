package config

type Config struct {
	Smtp SmtpConfig
}

func defaultConfig() Config {
	c := Config{}
	c.Smtp.defaultConfig()
	return c
}
