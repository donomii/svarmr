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

proc sendMessage {message} {



	puts stdout  [ dict2json $message ]
}

set displaytextvariable hello

tsv::set foo bar "A shared string"
set string [tsv::object foo bar]

$string set [thread::id]

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
		}
	}
]

