package require Tk 
package require json::write

proc dict2json {dictToEncode} {
	::json::write indented no
    ::json::write object {*}[dict map {k v} $dictToEncode {
        set v [::json::write string $v]
    }]
}

source theme.tcl

 set shape {0 -47 -25 -69 -65 -57 -80 -2 0 68 0 68 80 -2 65 -57 25 -69 0  -47 0 -47}
 set throb {1.0 1.05 1.10 1.05}
 pack [canvas .c -width 200 -height 200 -bg $color3 ]
 set i [eval .c create polygon $shape]

 .c itemconfigure $i -smooth true -fill $color2 -tag heart

 set i 0
 while {1} {
         if {!([incr i] % [llength $throb])} {set i 0}
         eval .c coords heart $shape
         set factor [lindex $throb $i]
         .c scale heart 0 0 $factor $factor
         .c move heart 100 100
         update
         after 100
 }

puts stdout  [ dict2json [ dict create Selector ModuleStart Arg Heart ] ]