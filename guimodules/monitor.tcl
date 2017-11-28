	package require Tk 
	package require Thread
	package require json::write
	package require json

source theme.tcl


frame .launchpad 
 label .launchpad.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "message goes here" -textvariable displaytextvariable -justify right
 pack .launchpad.label1

 pack .launchpad -side top

proc svarmrMessageHandler {$message} {
	SendSimple [dict create Selector GotMessage Arg $message]
}

source svarmrlib.tcl
