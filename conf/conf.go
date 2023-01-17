package conf

import (
	"log"
	"os"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
)

const (
	EnvConfigPath   = "CONFIG_PATH"
	EnvVarNamespace = "TEAPOT_"
	EnvDelimiter    = "_"
	PropDelimiter   = "."
	ConfFile        = ".env"
	cwd             = "."
)

// NewConf instantiates a new dotenv config with environment variables for context
// !!Note!! environment variables must be configured BEFORE calling NewConf
func NewConf(confName string) (*koanf.Koanf, error) {
	confName = defaultConfName(confName)
	koan := koanf.New(cwd)
	if err := koan.Load(file.Provider(confName), dotenv.Parser()); err != nil {
		log.Printf("error loading config: %s/%s: %v", cwd, confName, err)
		return koan, err
	}
	// load env variables under EnvVarNamespace namespace`
	err := koan.Load(env.Provider(EnvVarNamespace, EnvDelimiter, processEnvVar), nil)
	return koan, err
}

func defaultConfName(confName string) string {
	if len(strings.TrimSpace(confName)) == 0 {
		return ConfFile
	}
	return confName
}

func processEnvVar(s string) string {
	return strings.TrimPrefix(strings.Replace(strings.ToLower(
		strings.TrimPrefix(s, EnvVarNamespace)), EnvDelimiter, PropDelimiter, -1), PropDelimiter)
}

// GetEnv lookup an environment variable or fallback
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
