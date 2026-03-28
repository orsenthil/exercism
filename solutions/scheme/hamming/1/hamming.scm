(import (rnrs))

(define (hamming-distance strand-a strand-b)
  (apply +
    (map
      (lambda (a b)
        (if (eq? a b) 0 1))
      (string->list strand-a)
      (string->list strand-b)
    )
  )
)