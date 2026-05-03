(import (rnrs))

(define (leap-year? year)
  (and (= (% year 4) 0)
       (not (= (% year 100) 0))
       (or (= (% year 400) 0))))