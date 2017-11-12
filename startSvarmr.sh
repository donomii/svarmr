#!/bin/bash
cd "$(dirname ${BASH_SOURCE[0]})"
echo "Starting svarmr in " `pwd`
svarmr/server svarmr/example svarmr/usernotify systray/tray menu/ModuleMenu
