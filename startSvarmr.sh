#!/bin/bash
cd "$(dirname ${BASH_SOURCE[0]})"
echo "Starting svarmr in " `pwd`
svarmr/server svarmr/usernotify systray/tray
