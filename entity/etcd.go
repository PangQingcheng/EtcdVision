package entity

import (
	"encoding/json"
	"gorm.io/datatypes"
)

type ETCDDataSource struct {
	Name      string         `json:"name" gorm:"name;primaryKey"`
	Endpoints datatypes.JSON `json:"endpoints"`
}

func (s *ETCDDataSource) GetEndpoints() ([]string, error) {
	data, err := s.Endpoints.MarshalJSON()
	if err != nil {
		return nil, err
	}

	endpoints := []string{}
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
