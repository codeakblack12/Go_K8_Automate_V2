package config

import (
	"fmt"
	"strings"
)

func (c *Config) Validate() error {

	c.NodeRole = strings.ToLower(strings.TrimSpace(c.NodeRole))
	c.PodNetworkPlugin = strings.ToLower(strings.TrimSpace(c.PodNetworkPlugin))
	c.JoinServiceBaseURL = strings.TrimRight(strings.TrimSpace(c.JoinServiceBaseURL), "/")
	c.JoinCode = strings.TrimSpace(c.JoinCode)
	c.JoinCommand = strings.TrimSpace(c.JoinCommand)
	c.ControlPlaneJoinCommand = strings.TrimSpace(c.ControlPlaneJoinCommand)
	c.CertificateKey = strings.TrimSpace(c.CertificateKey)
	c.ControlPlaneEndpoint = strings.TrimSpace(c.ControlPlaneEndpoint)
	c.APIServerAddress = strings.TrimSpace(c.APIServerAddress)
	c.PodNetworkCIDR = strings.TrimSpace(c.PodNetworkCIDR)
	c.KubernetesVersion = strings.TrimSpace(c.KubernetesVersion)

	switch c.OSFamily {
	case "ubuntu":
	default:
		return fmt.Errorf("unsupported OS family: %s", c.OSFamily)
	}

	switch c.NodeRole {
	case "master", "worker", "control-plane":
	default:
		return fmt.Errorf("unsupported node role: %s", c.NodeRole)
	}

	switch c.PodNetworkPlugin {
	case "calico", "cilium":
	default:
		return fmt.Errorf("unsupported pod network plugin: %s", c.PodNetworkPlugin)
	}

	if c.KubernetesRepoVersion == "" {
		return fmt.Errorf("kubernetes repo version cannot be empty")
	}

	if c.NodeRole == "master" && c.APIServerAddress == "" {
		return fmt.Errorf("API server address cannot be empty for master nodes")
	}

	if c.NodeRole == "master" && c.ControlPlaneEndpoint == "" {
		return fmt.Errorf("control plane endpoint cannot be empty for HA-capable master initialization")
	}

	if c.NodeRole == "worker" && c.JoinCode == "" && c.JoinCommand == "" {
		return fmt.Errorf("worker nodes require either join code or join command")
	}

	if c.NodeRole == "control-plane" && c.JoinCode == "" && c.ControlPlaneJoinCommand == "" {
		return fmt.Errorf("control-plane nodes require either shared join code or control-plane join command")
	}

	if (c.NodeRole == "worker" || c.NodeRole == "control-plane") && c.JoinServiceBaseURL == "" {
		return fmt.Errorf("join service base URL cannot be empty")
	}

	return nil
}
