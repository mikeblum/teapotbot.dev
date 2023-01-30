package conf

import (
	"fmt"
	"os"
	"testing"

	"github.com/mikeblum/teapotbot.dev/conftest"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	// <setup code>
	suite, teardown := setupSuite(t)
	// <teardown code>
	defer teardown(t, suite.conf)
	t.Run("log=stdout", LogStdoutTest)
	t.Run("log=info", LogLevelInfoTest)
	t.Run("log=dotenv", DotEnvLogTest)
}

func LogStdoutTest(t *testing.T) {
	log := NewLog(conftest.TestConfFile)
	assert.NotNil(t, log)
	assert.Equal(t, os.Stdout, log.Logger.Out)
}

func LogLevelInfoTest(t *testing.T) {
	log := NewLog(conftest.TestConfFile)
	assert.NotNil(t, log)
	assert.Equal(t, logrus.InfoLevel, log.Logger.Level)
}

// NOTE: conf file must be populated before calling `NewConf`
func DotEnvLogTest(t *testing.T) {
	expected := "*logrus.JSONFormatter"
	// !!WARN!! `` injects \tabs
	cfg := "LOG_LEVEL=WARN\nLOG_FORMAT=JSON"
	err := os.WriteFile(conftest.TestConfFile, []byte(cfg), conftest.TestConfFilePerms)
	assert.Nil(t, err)
	log := NewLog(conftest.TestConfFile)
	assert.NotNil(t, log)
	assert.Equal(t, logrus.WarnLevel, log.Logger.Level)
	assert.Equal(t, expected, fmt.Sprintf("%T", log.Logger.Formatter))
}
