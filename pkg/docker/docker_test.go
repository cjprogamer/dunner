package docker

import (
	"fmt"

	"github.com/leopardslab/dunner/internal/settings"
	"github.com/spf13/viper"
)

func ExampleStep_Exec() {
	settings.Init()
	var testNodeVersion = "10.15.0"
	step := &Step{
		Task:     "test",
		Name:     "node",
		Image:    "node:" + testNodeVersion,
		Commands: [][]string{{"node", "--version"}},
		Env:      nil,
		Volumes:  nil,
	}

	err := step.Exec()
	if err != nil {
		panic(err)
	}
	// Output: OUT: v10.15.0
}

func ExampleStep_workingDirAbs() {
	var testNodeVersion = "10.15.0"
	var absPath = "/go"
	err := runCommand([]string{"pwd"}, absPath, testNodeVersion)

	if err != nil {
		panic(err)
	}
	// Output: OUT: /go
}

func Example_workingDirRel() {
	var testNodeVersion = "10.15.0"
	var relPath = "./"
	err := runCommand([]string{"pwd"}, relPath, testNodeVersion)
	if err != nil {
		panic(err)
	}
	// Output: OUT: /dunner
}

func runCommand(command []string, dir string, nodeVer string) error {
	settings.Init()
	step := &Step{
		Task:    "test",
		Name:    "node",
		Image:   "node:" + nodeVer,
		Command: command,
		Env:     nil,
		Volumes: nil,
		WorkDir: dir,
	}

	return step.Exec()
}

func ExampleStep_execWithErr() {
	var testNodeVersion = "10.15.0"
	var relPath = "./"
	err := runCommand([]string{"ls", "/invalid_dir" +
		""}, relPath, testNodeVersion)
	if err == nil {
		panic(err)
	}
	expectedErr := "Command execution failed with exit code 2"
	if err.Error() != expectedErr {
		panic(fmt.Errorf("expected error: %s, got: %s", expectedErr, err.Error()))
	}
	// Output: OUT: ls: cannot access '/invalid_dir': No such file or directory
}

func ExampleStep_execDryRun() {
	dryRun := viper.GetBool("Dry-run")
	viper.Set("Dry-run", true)

	defer viper.Set("Async", dryRun)
	var testNodeVersion = "10.15.0"
	var relPath = "./"
	err := runCommand([]string{"ls", "/invalid_dir" +
		""}, relPath, testNodeVersion)
	if err != nil {
		panic(err)
	}
	// Output:
}
