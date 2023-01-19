package models

import (
	"encoding/json"
	"fmt"

	"github.com/PedroALeo/crudRedis/database"
	geojson "github.com/paulmach/go.geojson"
)

type Block struct {
	ID       string           `json:"id,omitempty"`
	Name     string           `json:"name,omitempty"`
	ParentID string           `json:"parentID,omitempty"`
	Centroid geojson.Geometry `json:"centroid,omitempty"`
	Value    float64          `json:"value,omitempty"`
}

func (b Block) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Block) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, b)
}

// GetAllBlocks get all the blocks from the database
//
// if there's an error or nothing on the database returns am empty Block slice
//
// Returns a BlockSlice
func GetAllBlocks() []Block {
	keys, err := database.DB.Keys(database.CTX, "*:*").Result()
	if err != nil {
		fmt.Println(err)
		return []Block{}
	}

	result, _ := database.DB.MGet(database.CTX, keys...).Result()
	if err != nil {
		fmt.Println(err)
		return []Block{}
	}

	var blocks []Block

	for _, item := range result {
		var block Block
		err := block.UnmarshalBinary([]byte(fmt.Sprint(item)))
		if err != nil {
			return []Block{}
		}
		blocks = append(blocks, block)
	}

	return blocks
}
