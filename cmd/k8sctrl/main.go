package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"k8s/config"
)

func promptInput(label string, defaultVal string) string {
	var input string

	if defaultVal != "" {
		fmt.Printf("%s [%s]: ", label, defaultVal)
	} else {
		fmt.Printf("%s: ", label)
	}

	fmt.Scanln(&input)

	if input == "" {
		return defaultVal
	}

	return input
}

func confirmConfig(cfg *config.Config) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n----- CONFIG SUMMARY -----")
	fmt.Printf("Role: %s\n", cfg.NodeRole)
	fmt.Printf("API Server Address: %s\n", cfg.APIServerAddress)
	fmt.Printf("Control Plane Endpoint: %s\n", cfg.ControlPlaneEndpoint)
	fmt.Printf("Join Command: %s\n", cfg.JoinCommand)
	fmt.Printf("Control Plane Join Command: %s\n", cfg.ControlPlaneJoinCommand)
	fmt.Printf("Join Code: %s\n", cfg.JoinCode)
	fmt.Printf("Join Service URL: %s\n", cfg.JoinServiceBaseURL)
	fmt.Printf("Pod Network CIDR: %s\n", cfg.PodNetworkCIDR)
	fmt.Printf("Pod Network Plugin: %s\n", cfg.PodNetworkPlugin)
	fmt.Printf("Kubernetes Version: %s\n", cfg.KubernetesVersion)
	fmt.Printf("Repo Version: %s\n", cfg.KubernetesRepoVersion)
	fmt.Printf("Reset Node: %t\n", cfg.ResetNode)
	fmt.Println("--------------------------")

	for {
		fmt.Print("Proceed? (y/n): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "y" || input == "yes" {
			return true
		}
		if input == "n" || input == "no" {
			return false
		}

		fmt.Println("Please enter 'y' or 'n'")
	}
}

func main() {
	cfg := config.New()
	role := flag.String("role", cfg.NodeRole, "Node role: master, worker, or control-plane")
	joinCommand := flag.String("join-command", cfg.JoinCommand, "Full kubeadm join command for worker nodes")
	controlPlaneJoinCommand := flag.String("control-plane-join-command", cfg.ControlPlaneJoinCommand, "Full kubeadm join command for control-plane nodes")
	joinCode := flag.String("join-code", cfg.JoinCode, "Shared join code for worker or control-plane nodes")
	joinServiceURL := flag.String("join-service-url", cfg.JoinServiceBaseURL, "Base URL for the join-code service")
	apiServerAddress := flag.String("apiserver-address", cfg.APIServerAddress, "API server advertise address for master node")
	controlPlaneEndpoint := flag.String("control-plane-endpoint", cfg.ControlPlaneEndpoint, "Stable control plane endpoint, e.g. 10.0.0.100:6443")
	podNetworkCIDR := flag.String("pod-network-cidr", cfg.PodNetworkCIDR, "Pod network CIDR for cluster initialization")
	podNetworkPlugin := flag.String("pod-network-plugin", cfg.PodNetworkPlugin, "Pod network plugin: calico or cilium")
	kubernetesVersion := flag.String("kubernetes-version", cfg.KubernetesVersion, "Optional Kubernetes version for kubeadm init")
	repoVersion := flag.String("repo-version", cfg.KubernetesRepoVersion, "Kubernetes apt repository version, e.g. v1.35")
	resetNode := flag.Bool("reset-node", cfg.ResetNode, "If true, reset the node with kubeadm reset before initializing")

	flag.Parse()

	cfg.NodeRole = *role
	if cfg.NodeRole == "" {
		cfg.NodeRole = promptInput("Enter node role (master/worker/control-plane)", "")
	}
	cfg.JoinCommand = *joinCommand
	cfg.ControlPlaneJoinCommand = *controlPlaneJoinCommand
	cfg.JoinCode = *joinCode

	switch cfg.NodeRole {
	case "master":
		if cfg.APIServerAddress == "" {
			cfg.APIServerAddress = promptInput("Enter API server address", "")
		}

	case "worker":
		if cfg.JoinCommand == "" && cfg.JoinCode == "" {
			cfg.JoinCode = promptInput("Enter join code", "")
		}

	case "control-plane":
		if cfg.ControlPlaneJoinCommand == "" && cfg.JoinCode == "" {
			cfg.JoinCode = promptInput("Enter join code", "")
		}
	}

	cfg.JoinServiceBaseURL = *joinServiceURL
	cfg.APIServerAddress = *apiServerAddress
	cfg.ControlPlaneEndpoint = *controlPlaneEndpoint
	cfg.PodNetworkCIDR = *podNetworkCIDR
	cfg.PodNetworkPlugin = *podNetworkPlugin
	cfg.KubernetesVersion = *kubernetesVersion
	cfg.KubernetesRepoVersion = *repoVersion
	cfg.ResetNode = *resetNode

	if err := cfg.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "configuration error: %v\n", err)
		os.Exit(1)
	}

	if !confirmConfig(cfg) {
		fmt.Println("Operation cancelled by user.")
		os.Exit(0)
	}

	fmt.Println("cluster setup completed successfully")
}
