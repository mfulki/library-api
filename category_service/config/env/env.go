package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	commonErr  = "env %s is must be filled out"
	convIntErr = "env %s is must be integer"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	log.Fatal(fmt.Errorf(commonErr, key))
	return ""
}

func getIntEnv(key string) int {
	if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return val
	}

	log.Fatal(fmt.Errorf(convIntErr, key))
	return 0
}
