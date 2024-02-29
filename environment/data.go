package environment

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jackjohn7/smllnk/utils"
)

type (
	dbEnv struct {
		DATABASE_URL string
		REDIS_URL    string
		REDIS_PW     string
	}

	authEnv struct {
		SessionSecret []byte
	}

	security struct {
		CSRF_KEY string
	}

	env struct {
		IsProd  bool
		Port    string
		DbEnv   dbEnv
		AuthEnv authEnv
		Sec     security
	}
)

var Env *env

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

	Env = &env{
		DbEnv: dbEnv{
			DATABASE_URL: os.Getenv("DATABASE_URL"),
			REDIS_URL:    os.Getenv("REDIS_URL"),
			REDIS_PW:     os.Getenv("REDIS_PASSWORD"),
		},
		IsProd: buildEnv == "PROD",
		Port:   port,
		AuthEnv: authEnv{
			SessionSecret: sessionSecret,
		},
		Sec: security{
			CSRF_KEY: os.Getenv("CSRF_KEY"),
		},
	}

	if Env.Sec.CSRF_KEY == "" {
		panic("Missing 32-byte long CSRF_KEY. Add it to your environment.")
	}
}
