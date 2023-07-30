package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runKubectlCommand(args ...string) error {
	command := "kubectl"
	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error executing kubectl command: %s", err)
	}

	return nil
}

func Apply_Manifests(manifests []string) {
	for _, manifest := range manifests {
		if err := runKubectlCommand("apply", "-f", manifest); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s applied successfully!\n", manifest)
	}
}
