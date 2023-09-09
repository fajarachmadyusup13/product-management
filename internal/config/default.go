package config

import "time"

const (
	// RetryAttempts :nodoc:
	RetryAttempts = 5
	// DefaultCockroachMaxIdleConns :nodoc:
	DefaultCockroachMaxIdleConns = 2
	// DefaultCockroachMaxOpenConns max connection pool
	DefaultCockroachMaxOpenConns = 5
	// DefaultCockroachConnMaxLifetime :nodoc:
	DefaultCockroachConnMaxLifetime = 1 * time.Hour
	// DefaultCockroachPingInterval :nodoc:
	DefaultCockroachPingInterval = 1 * time.Second
)
