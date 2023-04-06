package cmd

import (
	"fmt"

	"github.com/kli/kubernetes"
	"github.com/kli/utils"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	k8s "k8s.io/client-go/kubernetes"
)

var diagnoseCmd = &cobra.Command{
	Use:     "diagnose",
	Aliases: []string{"health", "check"},
	Short:   "Diagnoses related operations",
	RunE:    diagnoseAll,
}

func init() {
	rootCmd.AddCommand(diagnoseCmd)
}

func diagnoseNamespaces(client *k8s.Clientset) {

	logrus.Infof("Diagnosing namespaces")
	var criticalNamespaces = []string{"kube-system"}

	namespaces, err := kubernetes.GetNamespaces(client)

	if err != nil {
		logrus.Fatalf("Failed to retrieve namespaces: %v", err)
	}

	fmt.Printf("Found %v namespaces\n", len(namespaces))

	for _, ns := range criticalNamespaces {
		if utils.ArrayContains(namespaces, ns) {
			fmt.Printf(utils.BoldString("  * %s namespace exists\n", color.New(color.FgGreen).Sprintf(ns)))
		} else {
			fmt.Printf("  * Unable to find %v namespace\n", ns)
		}
	}
}

func diagnoseAll(cmd *cobra.Command, args []string) error {
	logrus.Infof("Running diagnostics")

	client := kubernetes.GetKubernetesClient()
	diagnoseNamespaces(client)

	return nil
}
