package config

// Config stores runtime settings for the cluster creator application.
type Config struct {
	OSFamily                string
	KubernetesRepoVersion   string
	APIServerAddress        string
	PodNetworkCIDR          string
	KubernetesVersion       string
	NodeRole                string // master, worker, control-plane
	JoinCommand             string
	JoinCode                string
	PodNetworkPlugin        string
	JoinServiceBaseURL      string
	ResetNode               bool
	ControlPlaneJoinCommand string
	ControlPlaneJoinCode    string
	CertificateKey          string
	ControlPlaneEndpoint    string
}

// New creates a Config populated with default values.
func New() *Config {
	return &Config{
		OSFamily:                defaultOSFamily,
		KubernetesRepoVersion:   defaultKubernetesRepoVersion,
		APIServerAddress:        defaultAPIServerAddress,
		PodNetworkCIDR:          defaultPodNetworkCIDR,
		KubernetesVersion:       defaultKubernetesVersion,
		NodeRole:                defaultNodeRole,
		JoinCommand:             "",
		JoinCode:                "",
		PodNetworkPlugin:        defaultPodNetworkPlugin,
		JoinServiceBaseURL:      defaultJoinServiceBaseURL,
		ResetNode:               defaultResetNode,
		ControlPlaneJoinCommand: "",
		ControlPlaneJoinCode:    "",
		CertificateKey:          "",
		ControlPlaneEndpoint:    "",
	}
}
