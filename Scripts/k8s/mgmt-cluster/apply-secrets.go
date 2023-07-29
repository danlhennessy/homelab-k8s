package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateKubernetesSecret(namespace, secretName string, data map[string][]byte) error {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return fmt.Errorf("failed to build config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes clientset: %v", err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
		},
		Data: data,
		Type: corev1.SecretTypeOpaque,
	}

	_, err = clientset.CoreV1().Secrets(namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Kubernetes secret: %v", err)
	}

	fmt.Printf("Secret '%s' created in namespace '%s'\n", secretName, namespace)
	return nil
}

func Apply_String_Secret(namespace string, secretName string, key string, value string) {
	secretData := map[string][]byte{
		key: []byte(value),
	}

	err := CreateKubernetesSecret(namespace, secretName, secretData)
	if err != nil {
		fmt.Println("Error creating Kubernetes secret:", err)
		return
	}

	fmt.Println("Kubernetes secret created successfully.")
}

func Apply_Multi_String_Secret(namespace, secretName string, secretData map[string]string) {
	data := make(map[string][]byte)
	for key, value := range secretData {
		data[key] = []byte(value)
	}

	err := CreateKubernetesSecret(namespace, secretName, data)
	if err != nil {
		fmt.Println("Error creating Kubernetes secret:", err)
		return
	}

	fmt.Println("Kubernetes secret created successfully.")
}

func Apply_JSON_Secret(namespace string, secretName string, value string) {

	// Convert the JSON string to a map[string]string
	var data map[string]string
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Convert the string values to byte arrays
	secretData := make(map[string][]byte)
	for key, value := range data {
		secretData[key] = []byte(value)
	}

	CreateKubernetesSecret(namespace, secretName, secretData)
}
