package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("not equal")
    }
    length := len(a)
    var count int = 0;
    for i := 0; i < length; i++ {
        if (a[i] != b[i]) {
            count += 1
        }
    }
    return count, nil
}
