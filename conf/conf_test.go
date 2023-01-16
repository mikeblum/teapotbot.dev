package conf

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	mockConfFile      = "mock.env"
	mockConfFilePerms = 0600
)

type ConfTestSuite struct {
	conf *os.File
}

func setupSuite(t *testing.T) (*ConfTestSuite, func(t *testing.T, conf *os.File)) {
	var conf *os.File
	var err error
	_, err = os.Stat(mockConfFile)
	if os.IsNotExist(err) {
		if conf, err = os.Create(mockConfFile); err != nil {
			log.Fatal(err)
		}
		defer conf.Close()
	} else {
		if conf, err = os.Open(mockConfFile); err != nil {
			log.Fatal(err)
		}
	}
	return &ConfTestSuite{
		conf: conf,
	}, teardownSuite
}

func teardownSuite(t *testing.T, conf *os.File) {
	err := os.Remove(mockConfFile)
	assert.Nil(t, err)
}

func TestConf(t *testing.T) {
	// <setup code>
	suite, teardown := setupSuite(t)
	// <teardown code>
	defer teardown(t, suite.conf)
	t.Run("conf=new", NewConfTest)
	t.Run("conf=dotenv", DotEnvConfTest)
	t.Run("conf=env-namespace", EnvConfTest)
	t.Run("conf=env-var", GetEnvVarTest)
	t.Run("conf=env-default", GetEnvDefaultTest)
}

func NewConfTest(t *testing.T) {
	conf, err := NewConf(mockConfFile)
	assert.Nil(t, err)
	assert.NotNil(t, conf)
}

// NOTE: conf file must be populated before calling `NewConf`
func DotEnvConfTest(t *testing.T) {
	expectedKey := "test"
	expectedValue := "test_file_value"
	// !!WARN!! `` injects \tabs
	cfg := fmt.Sprintf("%s=%s", expectedKey, expectedValue)
	err := os.WriteFile(mockConfFile, []byte(cfg), mockConfFilePerms)
	assert.Nil(t, err)
	conf, err := NewConf(mockConfFile)
	assert.Nil(t, err)
	for k, v := range conf.All() {
		key := k
		value := v
		log.Printf("%s -> %s\n", key, value)
	}
	assert.Equal(t, expectedValue, conf.Get(expectedKey).(string))
}

// NOTE: ENV_VARs must be declared before calling `NewConf`
func EnvConfTest(t *testing.T) {
	envVar := strings.Join([]string{EnvVarNamespace, "TEST"}, "_")
	expectedKey := "test"
	expectedValue := "test_env_value"
	os.Setenv(envVar, expectedValue)
	defer os.Unsetenv(envVar)
	conf, err := NewConf(mockConfFile)
	assert.Nil(t, err)
	assert.Equal(t, expectedValue, conf.Get(expectedKey).(string))
}

func GetEnvVarTest(t *testing.T) {
	envShell := "SHELL"
	assert.True(t, len(strings.TrimSpace(GetEnv(envShell, ""))) > 0)
}

func GetEnvDefaultTest(t *testing.T) {
	expected := "default env"
	assert.Equal(t, expected, strings.TrimSpace(GetEnv("", expected)))
}
