#lang racket/gui
(require net/base64)
(require json)
(define-values (in out) (tcp-connect "localhost" 4816))


(define bitmap #f)

(define f #f)
[define mess #f]


  

[define new-window [lambda []
[set! f (new frame% [parent f][label "Bitmap"])]
                             [set! bitmap (read-bitmap "temp_pic.jpg")]
                               [set! mess (new message% [parent f] [label bitmap])]
                             (send f show #t)
                     ]]

(define process-port (lambda (a-port)
                       [letrec [[line (read-line a-port 'linefeed)]
                             [h [string->jsexpr line]]]
                         [when [not [equal? line ""]]
                           
                           [when [equal? "snapshot" [hash-ref h 'Selector]]
                             [delete-file "temp_pic.jpg"]
                             (display-to-file (base64-decode [string->bytes/utf-8 [hash-ref h 'Arg]]) "temp_pic.jpg")
                             [thread new-window]
                             ]
                           ]
                         [sleep 1]
                         [process-port a-port]]))


[define inthread [thread [lambda [] [process-port in]]]]


[define [mainloop]
  
      
  [sleep 1]
  [mainloop]
  ]

[mainloop]