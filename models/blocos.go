package models

import (
	geojson "github.com/paulmach/go.geojson"
)

type Bloco struct {
	ID       string
	Name     string
	ParentID string
	Centroid geojson.Geometry
	Value    float64
}
