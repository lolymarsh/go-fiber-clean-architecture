package config

type App struct {
	Port         string `mapstructure:"PORT"`
	BodyLimit    int    `mapstructure:"BODY_LIMIT"`
	IdleTimeout  int64  `mapstructure:"IDLE_TIMEOUT"`
	AllowOrigins string `mapstructure:"ALLOW_ORIGINS"`
}

type Database struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Name     string `mapstructure:"NAME"`
}

type Redis struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Password string `mapstructure:"PASSWORD"`
	DB       int    `mapstructure:"DB"`
}

type Auth struct {
	SecretKey string `mapstructure:"SECRET_KEY"`
	Expire    int64  `mapstructure:"EXPIRE"`
}
