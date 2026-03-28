package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func modExp(base, exp, mod *big.Int) *big.Int {
	result := new(big.Int).SetInt64(1)
	return result.Exp(base, exp, mod)
}

func PrivateKey(p *big.Int) *big.Int {
	maxnumber := new(big.Int).Sub(p, big.NewInt(2))
	n, err := rand.Int(rand.Reader, maxnumber)
	if err != nil {
		panic(err)
	}
	return n.Add(n, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return modExp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return modExp(public2, private1, p)
}
