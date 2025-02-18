package db

type SimulationDomain struct {
	ChainId uint64 `json:"chainId" gorm:"column:chain_id"`
	From    string `json:"from" gorm:"column:from_address"`
}

type SimulationInputs struct {
	Inputs []map[string]any `json:"inputs" gorm:"serializer:json"`
}

type SimulationOptions struct {
	Simulate bool `json:"simulate" gorm:"column:simulate"`
	Submit   bool `json:"submit" gorm:"column:submit"`
}

type SimulationDefinition struct {
	Id               string `json:"id" gorm:"primaryKey"`
	SimulationDomain `gorm:"embedded"`
	SimulationInputs `gorm:"embedded"`
	Options          SimulationOptions `json:"options" gorm:"embedded"`
}
