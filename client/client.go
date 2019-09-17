package client

import (
	"github.com/gpestana/sealpir/server"
)

type Client struct {}
type Query struct {}

func New() *Client {
	return &Client{}
}

func (c *Client) GenerateQuery(uint64 index) *Query {
	return &Query{}
}

func (c *Client) DecodeReply(r *server.Reply) []byte {
	return []byte{}
}

func (c *Client) GenerateGaloisKeys() {}
func (c *Client) ComposeToCiphertext() {}
func (c *Client) ComputeInverseScales() {}

func computeIndices(uint64 index, []uint64 nvec) {
	result := []uint64{}
	num := len(nvec)
	product := 1

	for _, n := range nvec {
		product *= n
	}


	return result
}
