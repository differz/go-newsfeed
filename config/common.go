package config

import (
	"flag"
	"os"
	"strconv"
)

func Load() {
	flag.Parse()
}

func envString(env string, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func envInt64(env string, fallback int64) int64 {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	v, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fallback
	}
	return v
}

func envBool(env string, fallback bool) bool {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	v, err := strconv.ParseBool(e)
	if err != nil {
		return fallback
	}
	return v
}

func envFloat64(env string, fallback float64) float64 {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	v, err := strconv.ParseFloat(e, 64)
	if err != nil {
		return fallback
	}
	return v
}
