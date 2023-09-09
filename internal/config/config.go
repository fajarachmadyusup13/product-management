package config

import (
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// env constants
const (
	EnvProduction  = "production"
	EnvStaging     = "staging"
	EnvDevelopment = "development"
)

// HTTPPort :nodoc:
func HTTPPort() string {
	return viper.GetString("http_port")
}

// Env :nodoc:
func Env() string {
	return viper.GetString("env")
}

// PostgreSQLHost :nodoc:
func PostgreSQLHost() string {
	return viper.GetString("postgresql.host")
}

// PostgreSQLDatabase :nodoc:
func PostgreSQLDatabase() string {
	return viper.GetString("postgresql.database")
}

// PostgreSQLUsername :nodoc:
func PostgreSQLUsername() string {
	return viper.GetString("postgresql.username")
}

// PostgreSQLPassword :nodoc:
func PostgreSQLPassword() string {
	return viper.GetString("postgresql.password")
}

// PostgreSQLSSLMode :nodoc:
func PostgreSQLSSLMode() string {
	if viper.IsSet("postgresql.sslmode") {
		return viper.GetString("postgresql.sslmode")
	}
	return "disable"
}

// PostgreSQLMaxIdleConns :nodoc:
func PostgreSQLMaxIdleConns() int {
	if viper.GetInt("postgresql.max_idle_conns") <= 0 {
		return DefaultCockroachMaxIdleConns
	}
	return viper.GetInt("postgresql.max_idle_conns")
}

// PostgreSQLMaxOpenConns :nodoc:
func PostgreSQLMaxOpenConns() int {
	if viper.GetInt("postgresql.max_open_conns") <= 0 {
		return DefaultCockroachMaxOpenConns
	}
	return viper.GetInt("postgresql.max_open_conns")
}

// PostgreSQLConnMaxLifetime :nodoc:
func PostgreSQLConnMaxLifetime() time.Duration {
	if !viper.IsSet("postgresql.conn_max_lifetime") {
		return DefaultCockroachConnMaxLifetime
	}
	return time.Duration(viper.GetInt("postgresql.conn_max_lifetime")) * time.Millisecond
}

// PostgreSQLPingInterval :nodoc:
func PostgreSQLPingInterval() time.Duration {
	if viper.GetInt("postgresql.ping_interval") <= 0 {
		return DefaultCockroachPingInterval
	}
	return time.Duration(viper.GetInt("postgresql.ping_interval")) * time.Millisecond
}

// DatabaseDSN :nodoc:
func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
		PostgreSQLUsername(),
		PostgreSQLPassword(),
		PostgreSQLHost(),
		PostgreSQLDatabase(),
		PostgreSQLSSLMode())
}

// LogLevel :nodoc:
func LogLevel() string {
	return viper.GetString("log_level")
}

// GetConf :nodoc:
func GetConf() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("svc")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Warningf("%v", err)
	}
}
