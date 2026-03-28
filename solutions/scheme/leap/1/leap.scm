(import (rnrs))

(define (leap-year? year)
  (cond
    ((= (remainder year 400) 0) #t)
    ((= (remainder year 100) 0) #f)
    ((= (remainder year 4) 0) #t)
    (else #f)))