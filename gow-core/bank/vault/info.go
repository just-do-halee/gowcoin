package vault

// Info is the information of a vault.
type Info struct {
	Top        TopLevelInfo        `json:"top"`
	State      StateLevelInfo      `json:"state"`
	Supervisor SupervisorLevelInfo `json:"supervisor"`
}
