package main

import "fmt"

// Config holds configuration options for a service.
type Config struct {
	Timeout   int
	Verbose   bool
	LogOutput string
}

// Option is a functional option for configuring a service.
type Option func(*Config)

// WithTimeout sets the timeout option.
func WithTimeout(timeout int) Option {
	return func(c *Config) {
		c.Timeout = timeout
	}
}

// WithVerbose sets the verbose option.
func WithVerbose(verbose bool) Option {
	return func(c *Config) {
		c.Verbose = verbose
	}
}

// WithLogOutput sets the log output option.
func WithLogOutput(logOutput string) Option {
	return func(c *Config) {
		c.LogOutput = logOutput
	}
}

// NewService creates a new service with the provided options.
func NewService(options ...Option) *Config {
	config := &Config{
		Timeout:   30,
		Verbose:   false,
		LogOutput: "stdout",
	}

	for _, option := range options {
		option(config)
	}

	return config
}

func main() {
	// Create a new service with custom options.
	service := NewService(
		WithTimeout(60),
		WithVerbose(true),
		WithLogOutput("file"),
	)

	// Print the service configuration.
	fmt.Printf("Timeout: %d, Verbose: %t, LogOutput: %s\n", service.Timeout, service.Verbose, service.LogOutput)
}
