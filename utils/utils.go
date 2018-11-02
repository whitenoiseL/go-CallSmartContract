package utils

import (
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"encoding/hex"
)

func keccack256(s string) string{
	h := sha3.NewKeccak256()
	h.Write([]byte(s))
	hash := h.Sum(nil)
	res := "0x" + hex.EncodeToString(hash)
	return res
}