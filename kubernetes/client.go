package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"

	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetKubernetesClient() *k8s.Clientset {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")

	// Load kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Println("Failed to load kubeconfig:", err)
		os.Exit(1)
	}

	// Create Kubernetes clientset
	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to create Kubernetes client:", err)
		os.Exit(1)
	}

	return clientset
}
