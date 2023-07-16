package entity

type ETCDDataSource struct {
	Name      string   `json:"name"`
	Endpoints []string `json:"endpoints"`
	// Host      string   `json:"host"`
	//Port      string   `json:"port"`
}
