package internal

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Version: "1",
}

func Test_InitLogDebug(t *testing.T) {
	InitLog(testCmd, true, false, false)
	if log.DebugLevel != log.GetLevel() {
		t.Errorf("get log level = %s want %s", log.GetLevel(), log.DebugLevel)
	}
}

func Test_InitLogDevelop(t *testing.T) {
	InitLog(testCmd, false, true, false)
	if log.TraceLevel != log.GetLevel() {
		t.Errorf("get log level = %s want %s", log.GetLevel(), log.TraceLevel)
	}
}

func Test_InitLogVerbose(t *testing.T) {
	InitLog(testCmd, false, false, true)
	if log.InfoLevel != log.GetLevel() {
		t.Errorf("get log level = %s want %s", log.GetLevel(), log.InfoLevel)
	}
}

func Test_InitLogDefault(t *testing.T) {
	InitLog(testCmd, false, false, false)
	if log.WarnLevel != log.GetLevel() {
		t.Errorf("get log level = %s want %s", log.GetLevel(), log.WarnLevel)
	}
}
