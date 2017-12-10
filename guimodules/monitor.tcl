	package require Tk 
	package require Thread
	package require json::write
	package require json

source theme.tcl
set menuPack [list .launchpad -fill both]

frame .launchpad 
 label .launchpad.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "message goes here" -textvariable displaytextvariable -justify right
 pack .launchpad.label1
 
 tk::text .launchpad.text -width 40 -height 10
 pack .launchpad.text {*}$menuPack

 pack .launchpad -side top

source lib/svarmrlib.tcl


proc svarmrMessageHandler {message} {
SimpleSend "debug" [dict get $message "Selector"]
	if {[dict get $message "Selector"] eq "mdns-service-found-summary"} {
	.launchpad.text insert 1.0 "\n"
	.launchpad.text insert 1.0  [dict get $message Arg]
	}
	
}

SimpleSend "debug" "started monitor"
