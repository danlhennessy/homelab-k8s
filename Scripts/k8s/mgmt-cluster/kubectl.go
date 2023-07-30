package main

import (
	"fmt"
	"log"
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

func Apply_Manifests(namespace string, manifests []string) {
	for _, manifest := range manifests {
		if err := runKubectlCommand("apply", "-n", namespace, "-f", manifest); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s applied successfully!\n", manifest)
	}
}
func Create_Namespaces(namespaces []string) {
	for _, namespace := range namespaces {
		if err := runKubectlCommand("create", "ns", namespace); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s applied successfully!\n", namespace)
	}
}
func Helm_Install(chart string, release string, namespace string) {
	cmd := exec.Command("helm", "install", chart, release, "--namespace", namespace)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))
}
