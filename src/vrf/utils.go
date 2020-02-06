package vrf

import (
	"github.com/SebastianElvis/vrf-mining/src/vrf/ed25519/edwards25519"
	"golang.org/x/crypto/sha3"
)

func h1(m []byte) *edwards25519.ExtendedGroupElement {
	return hashToCurve(m)
}

func scalarMult(sk *[SecretKeySize]byte, h *edwards25519.ExtendedGroupElement) []byte {
	var gamma edwards25519.ExtendedGroupElement
	x, _ := expandSecret(sk)
	edwards25519.GeScalarMult(&gamma, x, h)

	var gammaBytes [32]byte
	gamma.ToBytes(&gammaBytes)
	return gammaBytes[:]
}

func h2(gammaBytes []byte, m []byte) []byte {
	hash := sha3.NewShake256()
	hash.Write(gammaBytes) // const length: Size
	hash.Write(m)

	// note that [Size]byte limits the length of out to Size
	var out [Size]byte
	hash.Read(out[:])
	return out[:]
}
