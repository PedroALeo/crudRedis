package models

import (
	"testing"

	"github.com/PedroALeo/crudRedis/database"
	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/assert"
)

var (
	c0 = Block{
		ID:       "C0:0",
		Name:     "Cliente A",
		ParentID: "0",
		Centroid: *geojson.NewPointGeometry([]float64{-48.289546966552734, -18.931050694554795}),
		Value:    10000,
	}
	f1 = Block{
		ID:       "F1:C0",
		Name:     "FAZENDA 1",
		ParentID: "C0",
		Centroid: *geojson.NewPointGeometry([]float64{-52.9046630859375, -18.132801356084773}),
		Value:    1000,
	}
)

var tree = Tree{
	Block: c0,
	Children: []Tree{
		{
			Block:    f1,
			Children: nil,
		},
	},
}

func MockTree(t *testing.T) {
	FlushDBTest()
	database.ConnectRedis()
	blocks := []Block{c0, f1}
	for _, block := range blocks {
		err := database.DB.Set(database.CTX, block.ID, block, 0).Err()
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestGetTreeById(t *testing.T) {
	t.Run("mocked tree", func(t *testing.T) {
		MockTree(t)
		defer FlushDBTest()

		got := GetTreeByID("C0")

		assert.Equal(t, tree, got)
	})
}
