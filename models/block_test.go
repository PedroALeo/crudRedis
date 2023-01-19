package models

import (
	"testing"

	"github.com/PedroALeo/crudRedis/database"
	geojson "github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/assert"
)

var testBlock = Block{
	ID:       "C0:0",
	Name:     "Test Block",
	ParentID: "0",
	Centroid: *geojson.NewPointGeometry([]float64{-48.289546966552734, -18.931050694554795}),
	Value:    10,
}

func AddTestBlockToDB() {
	database.DB.Set(database.CTX, testBlock.ID, testBlock, 0)
}

func FlushDBTest() {
	database.DB.FlushAll(database.CTX)
}

func TestGetAllBlocks(t *testing.T) {
	t.Run("get array of existing blocks", func(t *testing.T) {
		database.ConnectRedis()
		defer database.DB.Close()
		AddTestBlockToDB()
		defer FlushDBTest()
		got, _ := GetAllBlocks()
		assert.Equal(t, []Block{testBlock}, got)
	})
}
