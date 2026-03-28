package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	var result []rune
	for idx, c := range phoneNumber {
		if idx == 0 && c == '+' {
			continue
		}
		if c == ' ' || c == '-' || c == '.' {
			continue
		}
		if c >= '0' && c <= '9' {
			result = append(result, c)
		}
	}
	if len(result) == 11 && result[0] == '1' {
		result = result[1:]
	}
	if len(result) != 10 {
		return "", errors.New("Invalid phone number")
	}
	if result[0] < '2' || result[3] < '2' {
		return "", errors.New("Invalid phone number")
	}
	return string(result), nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phoneNumber[:3], phoneNumber[3:6], phoneNumber[6:]), nil
}
