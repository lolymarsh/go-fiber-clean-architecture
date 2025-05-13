package config

type App struct {
	Port         string `mapstructure:"app_port"`
	BodyLimit    int    `mapstructure:"app_body_limit"`
	IdleTimeout  int    `mapstructure:"app_idle_timeout"`
	AllowOrigins string `mapstructure:"app_allow_origins"`
}

type Database struct {
	Host            string `mapstructure:"db_host"`
	Port            string `mapstructure:"db_port"`
	User            string `mapstructure:"db_user"`
	Password        string `mapstructure:"db_password"`
	Name            string `mapstructure:"db_name"`
	MaxOpenConns    int    `mapstructure:"db_max_open_conns"`
	MaxIdleConns    int    `mapstructure:"db_max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"db_conn_max_lifetime"`
}

type Redis struct {
	Host     string `mapstructure:"redis_host"`
	Port     string `mapstructure:"redis_port"`
	Password string `mapstructure:"redis_password"`
	DB       int    `mapstructure:"redis_db"`
}

type Auth struct {
	SecretKey string `mapstructure:"auth_secret_key"`
	Expire    int    `mapstructure:"auth_expire"`
}
