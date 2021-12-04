package require Tk 
package require json::write

proc dict2json {dictToEncode} {
	::json::write indented no
    ::json::write object {*}[dict map {k v} $dictToEncode {
        set v [::json::write string $v]
    }]
}

source theme.tcl


frame .launchpad 
 entry .launchpad.entry1 -background $color2 -foreground white -relief ridge -highlightthickness 2 -highlightcolor $color3 -borderwidth 8 -font {Helvetica -18 bold} -width 35 -textvariable myvariable -justify right
 pack .launchpad.entry1

 button .launchpad.button2 -text "Launch Module" -background $textBackgroundColor -foreground $textColor -font $font -command {puts stdout  [ dict2json [ dict create Selector start-module Arg $myvariable ] ] }
 
 pack .launchpad.button2 -side top -fill x

 pack .launchpad -side top

bind .launchpad.entry1 <Return> {puts stdout  [ dict2json [ dict create Selector start-module Arg $myvariable ] ] }
focus .launchpad.entry1
puts stdout  [ dict2json [ dict create Selector ModuleStart Arg ModuleLoader ] ]