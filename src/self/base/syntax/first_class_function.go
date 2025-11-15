package main

import (
	"time"
)

type Client struct {
	config *Config
}
type Config struct {
	Timeout time.Duration
	Retries int
}

type Option func(config *Config)

func withTimeout(d time.Duration) Option {
	return func(config *Config) {
		config.Timeout = d
	}
}

func withRetries(n int) Option {
	return func(config *Config) {
		config.Retries = n
	}
}

func NewClient(options ...Option) *Client {
	cfg := &Config{
		Timeout: time.Second * 5,
		Retries: 3,
	}
	for _, opt := range options {
		opt(cfg)
	}
	return &Client{config: cfg}
}
