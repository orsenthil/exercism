(import (rnrs))

(define (square n)
  (if (<= 1 n 64)
      (expt 2 (- n 1))
      (error "square: n must be between 1 and 64")))

(define total
  (let rec ([n 64])
    (if (zero? n)
        0
        (+ (square n) (rec (- n 1))))))
