#lang racket/gui
;Change to using canvas% so user can resize window smaller than the bitmap!
(require net/base64)
(require json)
(define-values (in out) (tcp-connect "localhost" 4816))


[define scale 4.0]
(define bitmap #f)
[define do-resize #f]




[define viewer-frame% (class frame%
 
  (super-new)                ; superclass initialization

[define/override on-size [lambda [width height]
                         [let [[new-scale [/ [send  bitmap get-width] width 0.5]]]
                           [displayln new-scale]
                           [set! scale new-scale]
                           [handle-resize scale]
                           ]]]
                        )]

(define f #f)
[define mess #f]
[set! f (new viewer-frame% [parent f][label "Bitmap"])]
[set! bitmap [make-bitmap 5 5 #:backing-scale scale]] ;(read-bitmap "temp_pic.jpg" #:backing-scale 2.0)]
[set! mess (new message% [parent f] [label bitmap])]
(send mess auto-resize #t)
(send f show #t)

[define handle-resize
  [lambda [scale]
    [displayln do-resize]
    [when do-resize [displayln "Resizing"]
                      [set! bitmap (read-bitmap "temp_pic.jpg" )]
                       [let [[new-scale [/   [send  bitmap get-width] [send f get-width]]]]
                           
                           [set! scale new-scale]
                         [displayln scale]
                      [set! bitmap (read-bitmap "temp_pic.jpg" #:backing-scale scale)]
                      ;[send bitmap load-file "temp_pic.jpg"]
                      [send mess set-label bitmap]
                      (send f resize
                             [send  bitmap get-width]
                            [send  bitmap get-height])	]] 

(define process-port (lambda (a-port)
                       [letrec [[line (read-line a-port 'linefeed)]
                                [h [string->jsexpr line]]]
                      
                         [when [not [equal? line ""]]
                           [when [or [equal? "snapshot" [hash-ref h 'Selector]]
                                     [equal? "image" [hash-ref h 'Selector]]
                                     ]
                             [delete-file "temp_pic.jpg"]
                             (display-to-file (base64-decode [string->bytes/utf-8 [hash-ref h 'Arg]]) "temp_pic.jpg")
                             [new-picture]]]
                         ;[sleep 1]
                         [process-port a-port]]))


[define reset-resize [lambda[] [sleep 1][set! do-resize #t][reset-resize]]]
[define reset-thread [thread reset-resize]]
[define inthread [thread [lambda [] [process-port in]]]]


[define [mainloop]
  ;[sleep 1]
  [mainloop]]

;[mainloop]
