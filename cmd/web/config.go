package main

type Config struct {
	Port       string `mapstructure:"port"`
	DBDSN      string `mapstructure:"db_dsn"`
	DBName     string `mapstructure:"db_name"`
	DBHost     string `mapstructure:"db_host"`
	DBPort     string `mapstructure:"db_port"`
	DBUser     string `mapstructure:"db_user"`
	DBPassword string `mapstructure:"db_password"`
	JWTSecret  string `mapstructure:"jwt_secret"`
	JWTSalt    string `mapstructure:"jwt_salt"`
}
