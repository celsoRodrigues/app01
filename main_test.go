package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

var appName = "gocityTest"

func TestMain(m *testing.M) {

	fmt.Println("-> Building...")
	// checking if Microsoft Windows OS
	// if so, we append .exe to make it excutable by Go
	if runtime.GOOS == "windows" {
		appName += ".exe"
	}
	// We run a command to build the code
	// we are using appName string as tthe name of the excutable
	build := exec.Command("go", "build", "-o", "./bin/"+appName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building %s: %s", appName, err)
		os.Exit(1)
	}
	fmt.Println("-> Running...")
	// Running the
	result := m.Run()
	fmt.Println("-> Getting done...")
	os.Exit(result)
}

func TestBadArgs(t *testing.T) {
	var err error
	cmd := exec.Command("./bin/" + appName)
	out, err := cmd.CombinedOutput()
	sout := string(out) // because out is []byte
	if err == nil && !strings.Contains(sout, "usage") {
		fmt.Println(sout) // so we can see the full output
		t.Errorf("%v", err)
	}

	os.Remove("./bin/" + appName)
}
