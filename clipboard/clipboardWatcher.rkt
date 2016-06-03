#lang racket/gui
(require file/md5)
 (require racket/date)
 (require framework/splash)
(require json)
[define last-data ""]
[define last-md5 [md5 last-data]]
(define-values (in out) (tcp-connect "localhost" 4816))

(define process-port (lambda (a-port)
                       [let [[line (read-line a-port 'linefeed)]]
                         [displayln [string->jsexpr line]]
                         [process-port a-port]]))

[define serialise-bitmap [λ [a-bitmap]
                            [letrec [[a-string [open-output-string]]]
                                  (send a-bitmap save-file a-string 'png)
                              [get-output-bytes a-string]]]]

[define inthread [thread [lambda [] [process-port in]]]]
(define ht (make-hash)) 
[define [mainloop]
  
;(send the-clipboard get-clipboard-data "FileName" 0)
 ;(send the-clipboard get-clipboard-data "HTML" 0)    
  [letrec [[clip-bitmap (send the-clipboard get-clipboard-bitmap 0)]]
    [when clip-bitmap
      [letrec [
           [serialised-bitmap [serialise-bitmap clip-bitmap]]
        [this-md5 [md5 serialised-bitmap]]]
      [when  [not [equal? last-md5 this-md5]]
        (start-splash "WKeVH.gif" "Saving bitmap" 1)
        [display [format "Saving bitmap ~a/~a" this-md5 last-md5]][newline]
      
        [set! last-data clip-bitmap]
      [set! last-md5 this-md5]
      [shutdown-splash]
        [close-splash]
        ]]]]

  [let [[clip-data (send the-clipboard get-clipboard-string 0)]]
    
      [when [and [< 0 [string-length clip-data]] [not [equal? last-md5 [md5 clip-data]]]]
        (start-splash "WKeVH.gif" "Saving text" 1)
        [write clip-data][newline]
        
        (hash-set! ht 'Selector "clipboard-change")
        (hash-set! ht 'Arg clip-data)
        [display [jsexpr->string ht] out]
        [display [format "~a" #\newline] out]
        [flush-output out]
      ;[write [string-ref clip-data 0]]
      
        [set! last-data clip-data]
      [set! last-md5 [md5 clip-data]]
        [shutdown-splash]
        [close-splash]]
      ]
  
  
      
  [sleep 1]
  [mainloop]
  ]

[mainloop]
