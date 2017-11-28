#svarmrlib

#A library to connect to the svarmr message bus

#To use this, you must create a function to handle svarmr messages.  It must be called "svarmrMessageHandler".  svarmrMessageHandler will be called for each new message, and the data will be in the form of a dict.

#The dict will contain several fields, but most importantly, a "Selector" key, and an "Arg" key.  The Selector is the "procedure name", that the message is trying to call, and the Arg is the data for that function.

#Be prepared to ignore Selectors that you don't know how to handle, svarmr will sometimes route broadcast messages to your module.

#An example of including svarmr in your program


#proc svarmrMessageHandler {$message} {
#	SendSimple [dict create Selector GotMessage Arg $message]
#}
#
#source lib/svarmrlib.tcl

#This will respond to every svarmr message with another message - making this an "echo server"

#The input processing runs in its own thread, and calls svarmrMessageHandler in the main thread.

#Note that svarmr usually takes control of your STDIN and STDOUT, to pass messages.

package require Tk 
package require Thread
package require json::write
package require json

proc dict2json {dictToEncode} {
	::json::write indented no
    ::json::write object {*}[dict map {k v} $dictToEncode {
        set v [::json::write string $v]
    }]
}

proc jsonget {json args} {
    foreach key $args {
        if {[dict exists $json $key]} {
            set json [dict get $json $key]
        } elseif {[string is integer $key]} {
            if {$key >= 0 && $key < [llength $json]} {
                set json [lindex $json $key]
            } else {
                error "can't get item number $key from {$json}"
            }
        } else {
            error "can't get \"$key\": no such key in {$json}"
        }
    }
    return $json
}

proc SimpleSend {message} {
	puts stdout  [ dict2json $message ]
}

set displaytextvariable hello

tsv::set foo bar "A shared string"
set string [tsv::object foo bar]

$string set [thread::id]

#set svarmrMessageHandler [SimpleSend [dict create "Selector" "gotMessage" "Arg" $message]]

proc process_message {message} {

$svarmrMessageHandler [::json::json2dict $message]

}

set t1 [thread::create {
		package require json

		proc jsonget {json args} {
			foreach key $args {
				if {[dict exists $json $key]} {
					set json [dict get $json $key]
				} elseif {[string is integer $key]} {
					if {$key >= 0 && $key < [llength $json]} {
						set json [lindex $json $key]
					} else {
						error "can't get item number $key from {$json}"
					}
				} else {
					error "can't get \"$key\": no such key in {$json}"
				}
			}
			return $json
		}
		
		set string [tsv::object foo bar]
		while (1) {
			set mainThread [ $string get ]
			gets stdin message
			thread::send $mainThread [list set displaytextvariable [jsonget [::json::json2dict $message] Selector]]
			thread::send $mainThread [list process_message $message]
		}
	}
]

SimpleSend {"Selector" "module-started" "Arg" "tcl lib"}