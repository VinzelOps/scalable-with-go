package utils

import (
	"crypto"
	"encoding/asn1"
	"fmt"
	"golang.org/x/crypto/cryptobyte"
	"golang.org/x/crypto/cryptobyte/asn1"
)

// Struktur sederhana untuk demonstrasi
type Person struct {
	Name string
	Age  int
}

// Fungsi untuk membangun struktur ASN.1 menggunakan cryptobyte
func BuildASN1Structure(p Person) ([]byte, error) {
	var b cryptobyte.Builder
	b.AddASN1(asn1.SEQUENCE, func(b *cryptobyte.Builder) {
		b.AddASN1UTF8String(p.Name)
		b.AddASN1Integer(big.NewInt(int64(p.Age)))
	})

	return b.Bytes()
}

// Fungsi untuk membaca struktur ASN.1 menggunakan cryptobyte
func ReadASN1Structure(data []byte) (Person, error) {
	var p Person
	input := cryptobyte.String(data)
	var nameStr, ageInt cryptobyte.String

	if !input.ReadASN1(&input, asn1.SEQUENCE) ||
		!input.ReadASN1UTF8String(&nameStr) ||
		!input.ReadASN1Integer(&ageInt) {
		return p, fmt.Errorf("error reading ASN.1 structure")
	}

	p.Name = string(nameStr)
	age, ok := ageInt.ReadASN1Integer()
	if !ok {
		return p, fmt.Errorf("error reading age")
	}
	p.Age = int(age)

	return p, nil
}
