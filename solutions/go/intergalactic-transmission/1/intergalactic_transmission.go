package intergalactictransmission

import "errors"

func Transmit(message []byte) []byte {
	var result []byte
	var current byte
	var count int
	var ones int
	var parity byte

	for _, b := range message {

		for n := 0; n < 8; n++ {
			bit := (b >> (7 - n)) & 1
			if bit == 1 {
				ones++
			}
			current = (current << 1) | bit
			count++
			if count == 7 {
				if ones%2 == 0 {
					parity = 0
				} else {
					parity = 1
				}

				result = append(result, (current<<1)|parity)
				current = 0
				count = 0
				ones = 0
			}
		}
	}

	if count > 0 {
		// pad current with 0s to make it 7 bits
		current = current << (7 - count)
		var ones int
		for n := 0; n < 7; n++ {
			bit := (current >> (7 - n)) & 1
			if bit == 1 {
				ones++
			}
		}
		parity := byte(ones % 2)
		result = append(result, (current<<1)|parity)
	}

	return result
}

func Decode(message []byte) ([]byte, error) {
	var result []byte
	var current byte
	var count int

	for _, b := range message {
		var ones int
		for n := 0; n < 8; n++ {
			bit := (b >> (7 - n)) & 1
			if bit == 1 {
				ones++
			}
		}
		if ones %2 != 0 {
			return nil, errors.New("wrong parity")
		}

		for n:= 0; n < 7; n++ {
			bit := (b >> (7 - n)) & 1
			current = (current << 1) | bit
			count++
			if count == 8 {
				result = append(result, current)
				current = 0
				count = 0
			}
		}
	}
	return result, nil
}
