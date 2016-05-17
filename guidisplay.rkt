#lang racket/gui
[require srfi/1]
[require srfi/13]
(require file/md5)
 (require racket/date)
 (require framework/splash)
(require json)
[define last-data ""]
[define last-md5 [md5 last-data]]
(define-values (in out) (tcp-connect "localhost" 4816))

; Make a frame by instantiating the frame% class 
(define frame (new frame% [label "Example"])) 
  
; Make a static text message in the frame 
(define messages [map [lambda [x](new message% [parent frame] 
                          [label "No events so far..........."])] [iota 10]])
[define msgs [make-hash]]
  
; Show the frame by calling its show method 
(send frame show #t) 

[define update-messages [lambda [] [map [lambda [label key][let [[l [format "~a:~a" key [hash-ref msgs key]]]]
                                                             (send label set-label 
                          [if [> [string-length l] 80]
                              [string-take l 80]
                              l])]] messages [hash-keys msgs]]]]

(define process-port (lambda (a-port)
                       [letrec [[line (read-line a-port 'linefeed)]
                             [h [string->jsexpr line]]]
                         [when [not [equal? line ""]]
                           
                           [hash-set! msgs [hash-ref h 'Selector] [hash-ref h 'Arg]]
                           
                           ]
                         [process-port a-port]]))

[define update-messages-loop [lambda [][update-messages][sleep 2][update-messages-loop]]]
[define inthread [thread [lambda [] [process-port in]]]]
[define updatethread [thread [lambda [] [update-messages-loop]]]]
(define ht (make-hash)) 
[define [mainloop]
  
      
  [sleep 1]
  [mainloop]
  ]

[displayln "Starting"]
[mainloop]