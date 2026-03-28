package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	// implement a rotational cipher
	var output []rune

	for _, r := range plain {
		base := 0
		switch {
		case r >= 'A' && r <= 'Z':
			base = int('A')
		case r >= 'a' && r <= 'z':
			base = int('a')
		}
		if base != 0 {
			r = rune(((int(r) - base + shiftKey) % 26) + base)
		}

		output = append(output, r)
	}

	return string(output)

}
