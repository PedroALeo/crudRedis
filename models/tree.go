package models

import (
	"strings"

	"github.com/PedroALeo/crudRedis/database"
)

type Tree struct {
	Block    Block  `json:"block,omitempty"`
	Children []Tree `json:"children,omitempty"`
}

func GetTreeByID(id string) Tree {
	var tree Tree
	var keysChildren []string
	var blockId string
	var err error

	tree.Block, err = GetBlockByID(id)
	if err != nil {
		return Tree{}
	}

	idM := strings.Split(tree.Block.ID, ":")
	blockId = idM[0]

	path := "*:" + blockId
	keysChildren, _ = database.DB.Keys(database.CTX, path).Result()

	for _, child := range keysChildren {
		ID := strings.Split(child, ":")
		tree.Children = append(tree.Children, GetTreeByID(ID[0]))
	}

	return tree
}
