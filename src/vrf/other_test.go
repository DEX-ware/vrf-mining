package vrf

import (
	"testing"
)

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

func BenchmarkMyCompute(b *testing.B) {
	_, sk, err := GenerateKey(nil)
	if err != nil {
		b.Fatal(err)
	}
	alice := []byte("alice")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		MyCompute(alice, sk)
	}
}

func BenchmarkH1(b *testing.B) {
	alice := []byte("alice")
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		h1(alice)
	}
}

func BenchmarkScalarMult(b *testing.B) {
	_, sk, err := GenerateKey(nil)
	if err != nil {
		b.Fatal(err)
	}
	alice := []byte("alice")
	h := h1(alice)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		scalarMult(sk, h)
	}
}

func BenchmarkH2(b *testing.B) {
	_, sk, err := GenerateKey(nil)
	if err != nil {
		b.Fatal(err)
	}
	alice := []byte("alice")
	h := h1(alice)
	gamma := scalarMult(sk, h)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		h2(gamma, alice)
	}
}
