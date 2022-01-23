package config

import (
	"fmt"
	"os"
	"strconv"
)

func getEnvStr(k string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		return "", fmt.Errorf("env %s not found", k)
	}
	return v, nil
}

func getEnvStrWithDefault(k string, dv string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		return dv, nil
	}
	return v, nil
}

func getEnvInt(k string) (int, error) {
	v := os.Getenv(k)
	if v == "" {
		return 0, fmt.Errorf("env %s not found", k)
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s not integer", k)
	}
	return i, nil
}

func getEnvIntWithDefault(k string, dv int) (int, error) {
	v := os.Getenv(k)
	if v == "" {
		return dv, nil
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s not integer", k)
	}
	return i, nil
}
