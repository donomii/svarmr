package require Tk 
package require Thread
package require json::write

proc dict2json {dictToEncode} {
	::json::write indented no
    ::json::write object {*}[dict map {k v} $dictToEncode {
        set v [::json::write string $v]
    }]
}

source theme.tcl


frame .launchpad 
 label .launchpad.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "message goes here" -textvariable displaytextvariable -justify right
 pack .launchpad.label1

 pack .launchpad -side top


puts stdout  [ dict2json [ dict create Selector ModuleStart Arg ModuleLoader ] ]

set displaytextvariable hello

tsv::set foo bar "A shared string"
set string [tsv::object foo bar]

$string set [thread::id]

set t1 [thread::create {
		set string [tsv::object foo bar]
		while (1) {
			set mainThread [ $string get ]
			gets stdin message
			thread::send $mainThread [list set displaytextvariable $message]
		}
	}
]

