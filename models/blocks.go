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
	ErrStatusNotFound            = errors.New("NotFound")
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
	if err != nil {
		fmt.Println(err)
		return []Block{}, ErrStatusInternalServerError
	}

	result, err := database.DB.MGet(database.CTX, keys...).Result()
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

// GetBlockByID get an Block from the database by id(key)
//
// if there's an error or no block with that id returns an empty block
//
// returns the Block related to the key
func GetBlockByID(id string) (Block, error) {
	pattern := id + ":" + "*"
	key, err := database.DB.Keys(database.CTX, pattern).Result()
	if err != nil {
		return Block{}, ErrStatusInternalServerError
	}
	if key == nil {
		return Block{}, ErrStatusNotFound
	}

	result, err := database.DB.Get(database.CTX, key[0]).Result()
	if err != nil {
		return Block{}, ErrStatusInternalServerError
	}
	var block Block
	errb := block.UnmarshalBinary([]byte(fmt.Sprint(result)))
	if errb != nil {
		return Block{}, errb
	}

	return block, nil
}

func DeleteBlock(id string) error {
	pat := id + ":*"
	key, err := database.DB.Keys(database.CTX, pat).Result()
	if err != nil {
		return ErrStatusNotFound
	}

	database.DB.Del(database.CTX, key[0])
	return nil
}

func PutBlock(block *Block) (Block, error) {
	blockR, err := GetBlockByID(block.ID)
	if err != nil {
		return Block{}, ErrStatusNotFound
	}
	database.DB.Set(database.CTX, blockR.ID, block, 0)
	return *block, nil
}
