(import (rnrs))

(define (dna->rna dna)
  (string-map
    (lambda (c)
      (case c
        ((#\G) #\C)
        ((#\C) #\G)
        ((#\T) #\A)
        ((#\A) #\U)
        (else (error "Invalid DNA nucleotide")))
    )
  dna))