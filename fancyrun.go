package fancyrun

import (
	"fmt"
	"github.com/google/shlex"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/smurfless1/pathlib"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
)

func sanitizedLogFileName(value string) string {
	// todo return cleared up log file name with static incrementing serial, 32 chars long
	reg, err := regexp.Compile("[^a-zA-Z0-9-_]+")
	if err != nil {
		log.Fatal(err)
	}
	runes := []rune(reg.ReplaceAllString(value, ""))
	// first3 := string(s[0:3])
	return string(runes[Min(0, len(runes)-3):])
}

func FancyRunWithNamedLog(cmd string, cwd pathlib.Path, check bool, logFileName string) (*exec.Cmd, string, error) {
	parts, err := shlex.Split(cmd)
	if err != nil {
		logrus.Fatal(err)
	}
	cmdobj := exec.Command(parts[0], parts[1:]...)
	cmdobj.Dir = cwd.String()
	logrus.Info(fmt.Sprintf("pushd %s ; %s ; popd", cwd.String(), cmd))
	cmdobj.Env = os.Environ()

	outBytes, err := cmdobj.Output()
	// write output to log file TODO at default log location
	err = ioutil.WriteFile(fmt.Sprintf("/tmp/%s.log", logFileName), outBytes, 0644)
	if check {
		CheckInline(err)
		if cmdobj.ProcessState.ExitCode() != 0 {
			logrus.Error(cmdobj.ProcessState.ExitCode())
			logrus.Error(string(outBytes))
			return cmdobj, string(outBytes), errors.New("Exit code was not 0")
		}
	}
	return cmdobj, string(outBytes), err
}

func FancyRun(cmd string, cwd pathlib.Path, check bool) (*exec.Cmd, string, error) {
	logFileName := sanitizedLogFileName(cmd)
	return FancyRunWithNamedLog(cmd, cwd, check, logFileName)
}
