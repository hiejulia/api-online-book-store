package utils

import (
	"os"
	"strconv"
)

// GetEnvInt will return an integer value from the env.
func GetEnvInt(key string) int {
	v := os.ExpandEnv(os.Getenv(key))
	if v == "" {
		v = "0"
	}

	i, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return i
}

// GetEnvStr will return a string value from the env.
func GetEnvStr(key string) string {
	return os.ExpandEnv(os.Getenv(key))
}

// FileExists will return true if the file exists, false if not.
func FileExists(path string) bool {
	path = os.ExpandEnv(path)
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
