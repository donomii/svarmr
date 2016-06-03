#lang racket/gui
(require net/base64)
(require json)
(define-values (in out) (tcp-connect "localhost" 4816))


[define scale 2.0]
(define bitmap #f)

(define f #f)
[define mess #f]
[set! f (new frame% [parent f][label "Bitmap"])]
[set! bitmap [make-bitmap 5 5 #:backing-scale scale]] ;(read-bitmap "temp_pic.jpg" #:backing-scale 2.0)]
[set! mess (new message% [parent f] [label bitmap])]
(send mess auto-resize #t)
(send f show #t)



[define new-picture [lambda []
                      [set! bitmap (read-bitmap "temp_pic.jpg" #:backing-scale scale)]
                      ;[send bitmap load-file "temp_pic.jpg"]
                      [send mess set-label bitmap]
                      (send f resize
                            [inexact->exact [* [/ 1.0 1.0] [send  bitmap get-width]]]
                            [inexact->exact [* [/ 1.0 1.0] [send  bitmap get-height]]])	]] 

(define process-port (lambda (a-port)
                       [letrec [[line (read-line a-port 'linefeed)]
                                [h [string->jsexpr line]]]
                         [when [not [equal? line ""]]
                           [when [equal? "snapshot" [hash-ref h 'Selector]]
                             [delete-file "temp_pic.jpg"]
                             (display-to-file (base64-decode [string->bytes/utf-8 [hash-ref h 'Arg]]) "temp_pic.jpg")
                             [thread new-picture]]]
                         [sleep 0.01]
                         [process-port a-port]]))


[define inthread [thread [lambda [] [process-port in]]]]


[define [mainloop]
  [sleep 1]
  [mainloop]]

[mainloop]