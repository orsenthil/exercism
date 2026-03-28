package cipher

type Shift struct {
	encodeKey, decodeKey []rune
}

type Caesar Shift
type Vigenere Shift

func NewCaesar() Cipher {
	return Caesar{[]rune{3}, []rune{-3}}

}

func (c Caesar) Encode(input string) string {
	return shift(input, c.encodeKey)
}

func (c Caesar) Decode(input string) string {
	return shift(input, c.decodeKey)
}

func shift(s string, off []rune) string {
	var cipher []rune
	offIndex, l := 0, len(off)
	for _, r := range s {
		// change to lower case
		if r >= 'A' && r <= 'Z' {
			r = 'a' + (r - 'A')
		}

		var c rune = -1
		if r >= 'a' && r <= 'z' {
			c = 'a' + ((r - 'a' + off[offIndex] + 26) % 26)

			if offIndex++; offIndex >= l {
				offIndex = 0
			}

			cipher = append(cipher, c)
		}
	}
	return string(cipher)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	offset := rune(distance)
	return Shift{[]rune{offset}, []rune{-offset}}
}

func (c Shift) Encode(input string) string {
	return shift(input, c.encodeKey)
}

func (c Shift) Decode(input string) string {
	return shift(input, c.decodeKey)
}

func NewVigenere(key string) Cipher {
	if len(key) < 3 {
		return nil
	}

	var encodeKey, decodeKey []rune
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		offset := r - 'a'
		encodeKey = append(encodeKey, offset)
		decodeKey = append(decodeKey, -offset)
	}
	return Vigenere{encodeKey, decodeKey}
}

func (v Vigenere) Encode(input string) string {
	// implement encode for vigenere
	return shift(input, v.encodeKey)
}

func (v Vigenere) Decode(input string) string {
	// implement decode for vigenere
	return shift(input, v.decodeKey)
}
