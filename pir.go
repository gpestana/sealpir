package sealpir

import (
	"math/big"
)

type PirContext struct {
	N uint16
	D uint32
	ExpansionRatio uint32
	DecBitCount uint32
	NVec []uint64
}

//typedef std::vector<seal::Plaintext> Database;
type PirDatabase []big.Int

//typedef std::vector< std::vector< seal::Ciphertext >> PirQuery;
type PirQuery [][]big.Int

//typedef std::vector<seal::Ciphertext> PirReply;
type PirReply []big.Int
