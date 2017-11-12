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



proc emitMessage {selector data} {
	puts stdout  [ dict2json [ dict create Selector $selector Arg $data ] ] 
}

frame .launchpad -background $containerBackground
 label .launchpad.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "Modules"
 pack .launchpad.label1

 button .launchpad.button1 -text "Module Loader" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module moduleLoader/moduleLoader}
 pack .launchpad.button1
 
 button .launchpad.button2 -text "Message Sender" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module messageSender/messageSender}
 pack .launchpad.button2
  
 button .launchpad.button3 -text "Canvas" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/heart}
 pack .launchpad.button3

 button .launchpad.button4 -text "Animated Canvas" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/heartthrob}
 pack .launchpad.button4
 
 button .launchpad.button5 -text "Message Monitor" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/monitor}
 pack .launchpad.button5
 
 button .launchpad.button6 -text "User Notifier" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module svarmr/usernotify}
 pack .launchpad.button6
 
 button .launchpad.button7 -text "Tray Icon" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module systray/tray}
 pack .launchpad.button7
 
 
 
 pack .launchpad -side top

puts stdout  [ dict2json [ dict create Selector ModuleStart Arg MessageSender ] ]