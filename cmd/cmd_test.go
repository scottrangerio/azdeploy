package cmd

import (
	"testing"

	"github.com/scottrangerio/azdeploy/cmdtest"
)

func TestCmd(t *testing.T) {
	cmdtest.TestCommand(t, cmd)
}
