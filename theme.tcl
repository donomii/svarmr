source themeColors.tcl
set windowConfig [list -background $containerBackground -borderwidth 20 -bd 10 -relief ridge]

wm minsize . 300 300
. configure {*}$windowConfig