package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// PostgresConfig menyimpan konfigurasi untuk primary dan standby PostgreSQL
type PostgresConfig struct {
	DBName      string `json:"dbname"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	MaxIdleConn int    `json:"max-idle-connection"`
	MaxIdleTime int    `json:"max-idle-time"`
	MaxOpenConn int    `json:"max-open-connection"`
}

// ScyllaConfig stores ScyllaDB settings
type ScyllaConfig struct {
	Hosts       []string `json:"hosts"`
	Keyspace    string   `json:"keyspace"`
	Consistency string   `json:"consistency"`
}

// DBConfig menyimpan konfigurasi database
type DBConfig struct {
	PostgreSQL struct {
		Primary PostgresConfig `json:"primary"`
		Standby PostgresConfig `json:"standby"`
	} `json:"postgresql"`
	ScyllaDB ScyllaConfig `json:"scylladb"`
}

// LoadConfig membaca konfigurasi dari file `config.json`
func LoadConfig(configPath string) (*DBConfig, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &DBConfig{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// GenerateDSN membuat DSN (Data Source Name) berdasarkan konfigurasi PostgreSQL
func (p PostgresConfig) GenerateDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.Host, p.Port, p.Username, p.Password, p.DBName,
	)
}
