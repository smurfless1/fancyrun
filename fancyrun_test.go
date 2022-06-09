package lib

import (
	"fmt"
	"github.com/smurfless1/pathlib"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFancyRun(t *testing.T) {
	result, out, err := FancyRun("date", pathlib.New("~").ExpandUser(), false)
	read_out := strings.TrimRight(string(out), "\n")
	fmt.Println(read_out)
	fmt.Println(result.ProcessState.ExitCode())
	assert.Nil(t, err)
}

