package id

import (
	"github.com/bwmarrin/snowflake"
)

type Generator interface {
	Generate() string
}

type SnowflakeGenerator struct {
	node *snowflake.Node
}

func NewSnowflakeGenerator(nodeID int64) (*SnowflakeGenerator, error) {
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		return nil, err
	}

	return &SnowflakeGenerator{
		node: node,
	}, nil
}

func (g *SnowflakeGenerator) Generate() string {
	return ToBase64(g.node.Generate().Int64())
}
