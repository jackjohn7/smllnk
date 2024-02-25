package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jackjohn7/smllnk/utils"
)

type (
	DbEnv struct {
		DATABASE_URL string
		REDIS_URL    string
		REDIS_PW string
	}

	AuthEnv struct {
		SessionSecret []byte
	}

	Environment struct {
		IsProd  bool
		Port    string
		DbEnv   DbEnv
		AuthEnv AuthEnv
	}
)

var Env *Environment

func init() {
	// parse environment into struct

	err := godotenv.Load()
	if err != nil {
		if os.Getenv("BUILD_ENV") != "PROD" {
			// fail in DEV but PROD should have proper environment variables not .env
			log.Fatalln("Please configure your environment with a .env file")
		}
	}

	buildEnv := os.Getenv("BUILD_ENV")
	if buildEnv == "" {
		buildEnv = "DEV"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3005"
	}

	sessionSecret := []byte(os.Getenv("SESSION_SECRET"))
	if len(sessionSecret) < 32 {
		sec, err := utils.GenerateRandomBytes(32)
		if err != nil {
			log.Fatalln(err)
		}
		sessionSecret = sec
	}

	Env = &Environment{
		DbEnv: DbEnv{
			DATABASE_URL: os.Getenv("DATABASE_URL"),
			REDIS_URL:    os.Getenv("REDIS_URL"),
			REDIS_PW:     os.Getenv("REDIS_PASSWORD"),
		},
		IsProd: buildEnv == "PROD",
		Port:   port,
		AuthEnv: AuthEnv{
			SessionSecret: sessionSecret,
		},
	}
}
