package pkg

type HealthStatus struct {
	CPU           float64 `json:"cpu"`
	CPUPercentage float64 `json:"cpuPercentage"`
	RAM           float64 `json:"ram"`
	RAMPercentage float64 `json:"ramPercentage"`
	RAMMax        float64 `json:"ramMax"`
}
