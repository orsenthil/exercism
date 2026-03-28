package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) []byte {
	result := make([]byte, 0)
	for _, value := range input {
		if value == 0 {
			result = append(result, 0)
			continue
		}

		temp := make([]byte, 0)

		for value > 0 {
			byte := byte(value & 0x7f)
			value >>= 7
			if value > 0 {
				byte |= 0x80
			}
			temp = append(temp, byte)
		}
		result = append(result, temp...)

	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	result := make([]uint32, 0)
	value := uint32(0)

	for i, b := range input {
		// check for overflow
		if value&0xfe000000 != 0 {
			return nil, errors.New("Overflow")
		}

		if b&0x80 == 0 {
			value = (value << 7) | uint32(b)
			result = append(result, value)
			value = 0
		} else {
			value = (value << 7) | uint32(b&0x7f)
			if i == len(input)-1 {
				return nil, errors.New("Incomplete sequence")
			}
		}
	}

	return result, nil
}
