package config

import "github.com/kelseyhightower/envconfig"

// Env - 環境変数用の構造体
type Env struct {
	Port string `envconfig:"PORT" default:"8080"`
}

// LoadEnv - 環境変数の読み込み
func LoadEnv() (Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return env, err
	}

	return env, nil
}
