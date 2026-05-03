(import (rnrs))

(define (leap-year? year)
  (and (= (modulo year 4) 0)
       (not (= (modulo year 100) 0))
       (or (= (modulo year 400) 0))))