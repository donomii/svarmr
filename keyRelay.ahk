#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
#SingleInstance force
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.
#n::
Run Notepad
return
Volume_Mute::
Run sendMessage.exe localhost 4816 mute 0
return
Volume_Up::
Run sendMessage.exe localhost 4816 unmute 0
Run sendMessage.exe localhost 4816 volume-up 2
return
Volume_Down::
Run sendMessage.exe localhost 4816 unmute 0
Run sendMessage.exe localhost 4816 volume-down 2
return