package sealpir

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) GenerateQuery(index uint64) *PirQuery {
	return &PirQuery{}
}

func (c *Client) DecodeReply(r *PirReply) []byte {
	return []byte{}
}

func (c *Client) GenerateGaloisKeys()   {}
func (c *Client) ComposeToCiphertext()  {}
func (c *Client) ComputeInverseScales() {}

func computeIndices(index uint64, nvec []uint64) []uint64 {
	result := []uint64{}
	//num := len(nvec)
	product := uint64(1)

	for _, n := range nvec {
		product = product * n
	}

	return result
}
