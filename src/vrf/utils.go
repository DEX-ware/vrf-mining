package vrf

import (
	"github.com/SebastianElvis/vrf-mining/src/vrf/ed25519/edwards25519"
	"golang.org/x/crypto/sha3"
)

func h1(m []byte) *edwards25519.ExtendedGroupElement {
	return hashToCurve(m)
}

func h2(gammaBytes []byte, m []byte) []byte {
	hash := sha3.NewShake256()
	hash.Write(gammaBytes) // const length: Size
	hash.Write(m)

	var out []byte
	hash.Read(out)

	return out
}
