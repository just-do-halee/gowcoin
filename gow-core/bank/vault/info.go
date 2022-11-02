package vault

// Info is the information of a vault.
type Info struct {
	Top        TopLevelInfo
	State      StateLevelInfo
	Supervisor SupervisorLevelInfo
}
