package cipher

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
	distance int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	// implement caesar cipher
	return shift{3}

}

func NewShift(distance int) Cipher {
	// implement shift cipher
	return shift{distance}
}

func (c shift) Encode(input string) string {
	// implement encode for shift
	return input
}

func (c shift) Decode(input string) string {
	// implement decode
	return input
}

func NewVigenere(key string) Cipher {
	// implement vigenere cipher
	return vigenere{key}
}

func (v vigenere) Encode(input string) string {
	// implement encode for vigenere
	return input
}

func (v vigenere) Decode(input string) string {
	// implement decode for vigenere
	return input
}
