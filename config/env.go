package config

import "github.com/kelseyhightower/envconfig"

// Env - 環境変数用の構造体
type Env struct {
	Port       string `envconfig:"PORT" default:"8080"`
	DBName     string `envconfig:"DB_NAME" default:"examin"`
	DBUser     string `envconfig:"DB_USERNAME" default:"examin"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"examin"`
	DBHost     string `envconfig:"DB_HOST" default:"127.0.0.1"`
	DBPort     string `envconfig:"DB_PORT" default:"3306"`
	Secret     string `envconfig:"SECRET_KEY" default:"4CHtx3AAnOPrxWYvqsHou6w8b8aO3BF7Ey93/D4clbBsgMDZf9Zt+Q=="`
	Iss        string `envconfig:"ISS_USER" default:"16france"`
}

// LoadEnv - 環境変数の読み込み
func LoadEnv() (Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}
