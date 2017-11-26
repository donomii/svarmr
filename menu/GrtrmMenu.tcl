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
 label .f.label1 -background $textBackgroundColor -foreground $textColor -font $font -text "Grtrm Controls"
 pack .f.label1

 button .f.button1 -text "Forwards" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-forward 1}
 pack .f.button1  {*}$menuPack
 
 button .f.button2 -text "Backwards" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-backwards 1}
 pack .f.button2 {*}$menuPack
  
 button .f.button3 -text "Left" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-left 1}
 pack .f.button3 {*}$menuPack

 button .f.button4 -text "Right" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-right 1}
 pack .f.button4 {*}$menuPack
 
 button .f.button5 -text "Up" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-up 1}
 pack .f.button5 {*}$menuPack
 
 button .f.button6 -text "Down" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage move-down 1}
 pack .f.button6 {*}$menuPack
 
 button .f.button7 -text "Reset" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage view-reset 1}
 pack .f.button7 {*}$menuPack
 
 button .f.button8 -text "Report Position" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage report-position 1}
 pack .f.button8 {*}$menuPack
 
 button .f.button9 -text "Skin Mode" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage skin-mode 0}
 pack .f.button9 {*}$menuPack
 
 
 button .f.button10 -text "Increment p0" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage parameter0-add 0.1}
 pack .f.button10 {*}$menuPack
 
 button .f.button11 -text "Decrement p0" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage parameter0-add -0.1}
 pack .f.button11 {*}$menuPack
 
  button .f.button12 -text "Increment p1" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage parameter1-add 0.1}
 pack .f.button12 {*}$menuPack
 
 button .f.button13 -text "Decrement p1" -background $textBackgroundColor -foreground $textColor -font $font -command {emitMessage parameter1-add -0.1}
 pack .f.button13 {*}$menuPack
 

set p0 0.0
set p1 0.0

 
proc set_p0 {val} {
    set p0 $val
	emitMessage parameter0-set $val
}

proc set_p1 {val} {
    set p1 $val
	emitMessage parameter1-set $val
}

scale .f.slide1  -from -5.0  -to 5.0  -label p0 -orient horizontal -resolution 0.01  -variable gvar(p0)   -orient horizontal -command set_p0
 pack .f.slide1 {*}$menuPack

 scale .f.slide2  -from -5.0   -to 5.0  -label p1 -orient horizontal -resolution 0.01  -variable gvar(p1)   -orient horizontal -command set_p1
 pack .f.slide2 {*}$menuPack
 
 tk::text .f.text -width 40 -height 10
 pack .f.text {*}$menuPack
 
 button .f.button14 -text "Recompile" -background $textBackgroundColor -foreground $textColor -font $font -command {
	emitMessage shader-code	 [.f.text get 1.0 end]
	}
	
 pack .f.button14 {*}$menuPack
 
 
 
 pack .f -side top

puts stdout  [ dict2json [ dict create Selector ModuleStart Arg MessageSender ] ]