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
		got, _ := GetAllBlocks()
		assert.Equal(t, []Block{testBlock}, got)
		FlushDBTest()
	})
}

func TestGetBlockById(t *testing.T) {

	t.Run("get existing block by id", func(t *testing.T) {
		database.ConnectRedis()
		defer database.DB.Close()
		AddTestBlockToDB()

		got, _ := GetBlockByID("C0")
		assert.Equal(t, testBlock, got)

		FlushDBTest()
	})
}

func TestDeleteBlock(t *testing.T) {
	t.Run("Delete block by id", func(t *testing.T) {
		database.ConnectRedis()
		defer database.DB.Close()
		AddTestBlockToDB()

		got := DeleteBlock("C0")
		assert.Equal(t, nil, got)

		FlushDBTest()
	})
}

func TestPutBlock(t *testing.T) {
	t.Run("Updating block by id", func(t *testing.T) {
		database.ConnectRedis()
		defer database.DB.Close()
		AddTestBlockToDB()
		block := Block{ID: "C0", Name: "CLIENTE 4", ParentID: "0",
			Centroid: *geojson.NewPointGeometry([]float64{-48.289546966552734, -18.931050694554795}),
			Value:    0}
		got, _ := PutBlock(&block)
		assert.Equal(t, block, got)
	})
}

func TestNewBlock(t *testing.T) {

	t.Run("Creating New Block", func(t *testing.T) {
		database.ConnectRedis()
		defer database.DB.Close()
		block := Block{ID: "C1", Name: "CLIENTE 4", ParentID: "0",
			Centroid: *geojson.NewPointGeometry([]float64{-48.289546966552734, -18.931050694554795}),
			Value:    0}

		got, _ := NewBlock(&block)
		assert.Equal(t, block, got)
	})
}
