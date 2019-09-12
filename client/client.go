package client

import (
	"github.com/gpestana/sealpir/server"
)

type Client struct {}
type Query struct {}

func New() *Client {
	return &Client{}
}

func (c *Client) GenerateQuery() *Query {
	return &Query{}
}

func (c *Client) DecodeReply(r *server.Reply) []byte {
	return []byte{}
}

func (c *Client) GenerateGaloisKeys() {}
func (c *Client) ComposeToCiphertext() {}
func (c *Client) ComputeInverseScales() {}
