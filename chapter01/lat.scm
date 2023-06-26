(define lat?
  (lambda(l)
    (cond
       ((null? l)#t)
       ((atom? (car L))(lat ?(cdr l)))
    (else #f ))))