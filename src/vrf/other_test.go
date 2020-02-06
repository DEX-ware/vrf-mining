package vrf

import (
	"testing"

	"ekyu.moe/cryptonight"
	scrypt "github.com/elithrar/simple-scrypt"
	"github.com/seehuhn/sha256d"
)

func sha256df(b []byte) []byte {
	hash := sha256d.New()
	hash.Write(b)
	return hash.Sum(nil)
}

func scryptf(b []byte) []byte {
	hash, _ := scrypt.GenerateFromPassword(b, scrypt.DefaultParams)
	return hash
}

func cryptonightf(b []byte) []byte {
	return cryptonight.Sum(b, 0)
}

func BenchmarkSHA256D(b *testing.B) {
	alice := []byte("alice")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sha256df(alice)
	}
}

func BenchmarkScrypt(b *testing.B) {
	alice := []byte("alice")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		scryptf(alice)
	}
}

func BenchmarkCryptoNight(b *testing.B) {
	alice := []byte("alice")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cryptonightf(alice)
	}
}
