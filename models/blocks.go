package models

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/PedroALeo/crudRedis/database"
	geojson "github.com/paulmach/go.geojson"
)

var (
	ErrStatusInternalServerError = errors.New("InternalServerError")
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
func GetAllBlocks() ([]Block, error) {
	keys, err := database.DB.Keys(database.CTX, "*:*").Result()
	err = errors.New("fgfgsgf")
	if err != nil {
		fmt.Println(err)
		return []Block{}, ErrStatusInternalServerError
	}

	result, _ := database.DB.MGet(database.CTX, keys...).Result()
	if err != nil {
		fmt.Println(err)
		return []Block{}, ErrStatusInternalServerError
	}

	var blocks []Block

	for _, item := range result {
		var block Block
		err := block.UnmarshalBinary([]byte(fmt.Sprint(item)))
		if err != nil {
			return []Block{}, ErrStatusInternalServerError
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}
