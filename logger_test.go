package log2

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_SetLogger(t *testing.T) {
	myLogger, _ := test.NewNullLogger()
	SetLogger(myLogger)

	assert.Equal(t, myLogger, GetLogger(), "Should properly set logger")
}

func Test_GetLogger(t *testing.T) {
	myLogger, _ := test.NewNullLogger()
	SetLogger(myLogger)

	assert.Equal(t, myLogger, GetLogger(), "Should properly get logger")
}

func Test_Debug(t *testing.T) {
	myLogger, hook := test.NewNullLogger()
	SetLogger(myLogger)
	SetLevel(DebugLevel)

	Debug("Debug")
	assert.Equal(t, "Debug\n", hook.LastEntry().Message)
}

func Test_Warn(t *testing.T) {
	myLogger, hook := test.NewNullLogger()
	SetLogger(myLogger)
	SetLevel(WarnLevel)

	Warn("Warn")
	assert.Equal(t, "Warn\n", hook.LastEntry().Message)
}

func Test_Error(t *testing.T) {
	myLogger, hook := test.NewNullLogger()
	SetLogger(myLogger)
	SetLevel(ErrorLevel)

	Error("Error")
	assert.Equal(t, "Error\n", hook.LastEntry().Message)
}

func Test_Info(t *testing.T) {
	myLogger, hook := test.NewNullLogger()
	SetLogger(myLogger)
	SetLevel(InfoLevel)

	Info("Info")
	assert.Equal(t, "Info\n", hook.LastEntry().Message)
}

func Test_Log(t *testing.T) {
	levels := []Level{InfoLevel, WarnLevel, ErrorLevel, DebugLevel, TraceLevel}

	testCases := []struct {
		desc   string
		msg    string
		fields Fields
	}{
		{
			desc: "Works in info Level",
			msg:  "Testing logs ...",
		},
	}

	myLogger, hook := test.NewNullLogger()
	SetLogger(myLogger)

	for _, tC := range testCases {
		for _, level := range levels {
			t.Run(tC.desc+" on level "+fmt.Sprint(level), func(t *testing.T) {
				SetLevel(level)
				Log(level, tC.msg)
				assert.Equal(t, tC.msg+"\n", hook.LastEntry().Message)

				hook.Reset()
			})
		}
	}
}
