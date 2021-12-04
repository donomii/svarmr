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
 label .launchpad.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "Selector" -justify right
 pack .launchpad.label1
 
  entry .launchpad.entry1 -background $color2 -foreground white -relief ridge -highlightthickness 2 -highlightcolor $color3 -borderwidth 8 -font {Helvetica -18 bold} -width 35 -textvariable SelectorBound -justify right
 pack .launchpad.entry1

 label .launchpad.label2 -background $textBackgroundColor -foreground $textColor -font $font -text "Arg"  -justify right
 pack .launchpad.label2
 
 entry .launchpad.entry2 -background $color2 -foreground white -relief ridge -highlightthickness 2 -highlightcolor $color3 -borderwidth 8 -font {Helvetica -18 bold} -width 35 -textvariable ArgBound -justify right
 pack .launchpad.entry2

 
button .launchpad.button2 -text "Send Message" -background $textBackgroundColor -foreground $textColor -font $font -command {puts stdout  [ dict2json [ dict create Selector $SelectorBound Arg $ArgBound ] ] }
pack .launchpad.button2
 

 pack .launchpad -side top


bind .launchpad.entry1 <Return> {puts stdout  [ dict2json [ dict create Selector $SelectorBound Arg $ArgBound ] ] }
bind .launchpad.entry2 <Return> {puts stdout  [ dict2json [ dict create Selector $SelectorBound Arg $ArgBound ] ] }
focus .launchpad.entry2
 
puts stdout  [ dict2json [ dict create Selector ModuleStart Arg MessageSender ] ]

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
