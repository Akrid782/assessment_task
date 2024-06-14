package util

import (
	"os"
)

func GetEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
