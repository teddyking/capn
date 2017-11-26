package main

import (
	"fmt"
	"os"

	"github.com/syndtr/gocapability/capability"
)

func currentThreadCaps() (capability.Capabilities, error) {
	currentPid := os.Getpid()
	threadCaps, err := capability.NewPid(currentPid)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve capabilities for pid %d", currentPid)
	}

	return threadCaps, nil
}

func currentFileCaps() (capability.Capabilities, error) {
	currentFilepath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("unable to determine path of currently-running executable")
	}
	fileCaps, err := capability.NewFile(currentFilepath)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve capabilities for file %s", currentFilepath)
	}

	return fileCaps, nil
}

func main() {
	threadCaps, err := currentThreadCaps()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fileCaps, err := currentFileCaps()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("=== THREAD CAPS ===")
	fmt.Println(threadCaps.String())

	fmt.Println("=== FILE CAPS ===")
	fmt.Println(fileCaps.String())
}
