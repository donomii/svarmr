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

set menuPack [list .f -fill both]

frame .f -background $containerBackground
 label .f.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "Modules"
 pack .f.label1

 button .f.button1 -text "Module Loader" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module moduleLoader/moduleLoader}
 pack .f.button1  {*}$menuPack
 
 button .f.button2 -text "Message Sender" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module messageSender/messageSender}
 pack .f.button2 {*}$menuPack
  
 button .f.button3 -text "Canvas" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/heart}
 pack .f.button3 {*}$menuPack

 button .f.button4 -text "Animated Canvas" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/heartthrob}
 pack .f.button4 {*}$menuPack
 
 button .f.button5 -text "Message Monitor" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module guimodules/monitor}
 pack .f.button5 {*}$menuPack
 
 button .f.button6 -text "User Notifier" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module svarmr/usernotify}
 pack .f.button6 {*}$menuPack
 
 button .f.button7 -text "Tray Icon" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage start-module systray/tray}
 pack .f.button7 {*}$menuPack
 
 button .f.button8 -text "Grtrm" -background $textBackgroundColor -foreground $textColor -font $font -command {
 emitMessage start-module pocketbonsai
 emitMessage start-module menu/GrtrmMenu
 }
 pack .f.button8 {*}$menuPack
 
 
 
 pack .f -side top

puts stdout  [ dict2json [ dict create Selector ModuleStart Arg MessageSender ] ]