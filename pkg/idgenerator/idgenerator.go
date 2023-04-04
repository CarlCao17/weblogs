package idgenerator

import (
	"sync/atomic"

	"github.com/bwmarrin/snowflake"
)

// IDGenerator 采用 Snowflake 算法，每次生成一个近乎唯一的ID
// 关于该算法和生成的 ID，参见 https://github.com/bwmarrin/snowflake
type IDGenerator interface {
	ID() int64
}

type idGenerator struct {
	*snowflake.Node
}

// New 生成一个新的 IDGenerator，但总数不允许超过 1024 个（snowflake 限制）
func New() (IDGenerator, error) {
	node, err := snowflake.NewNode(nextNode.Add(1))
	if err != nil {
		return nil, err
	}
	return &idGenerator{node}, nil
}

var g IDGenerator

func GetIDGeneratorInstance() IDGenerator {
	if g == nil {
		g, _ = New()
	}
	return g
}

func (g *idGenerator) ID() int64 {
	return g.Generate().Int64()
}

var nextNode atomic.Int64
