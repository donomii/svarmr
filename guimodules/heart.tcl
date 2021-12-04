package require Tk 
package require json::write


proc dict2json {dictToEncode} {
	::json::write indented no
    ::json::write object {*}[dict map {k v} $dictToEncode {
        set v [::json::write string $v]
    }]
}

source theme.tcl

frame .f -background $containerBackground
canvas .f.c -width 200 -height 200 -bg $color3
 pack .f.c .f -fill both -expand true
 .f.c create polygon 100 55 75 33 35 45 20 100 100 170 100 170 180 100 165 45 125 33 100 55 100 55 -smooth true -fill $color2

puts stdout  [ dict2json [ dict create Selector ModuleStart Arg Heart ] ]