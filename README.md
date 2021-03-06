# svarmr

A networked control bus and messaging system, with many useful modules

* Link the volume on all your devices
* Automatically start torrents on a different machine
* Trigger on custom events
* Easy access to platform libraries

## Easy to add modules

Svarmr is a simple message bus that is super easy to connect to.

Simple modules can be a few lines long, and there are examples in Go, Racket, Perl, and C.

* [Go](https://github.com/donomii/svarmrgo)

## Features

Svarmr already has some useful modules, that monitor clipboard changes, respond to keys, speak text, and detect notes.

# Design

## Sources, sinks and processors

Each module can transmit and receive any events it pleases, but it is useful to break them into sources, sinks, and processors.  Sources mainly generate events, sinks wait for events and then do something, and processors wait for events, deal with them and then send new events.

### Sources

* clipboardWatcher  - Broadcasts clipboard changes
* mdnsWatcher       - Broadcasts mDNS results
* heartBeat         - Broadcasts a regular beat (more useful than it sounds)

### Sinks

* torrentListener   - Starts torrents in Deluge
* monitor.c         - Prints events to stdout
* monitor.go        - Prints events to stdout
* volume            - Sets the volume (needs helpers)
* userNotify        - Pops up a message on the screen (needs helpers)
* moduleStarter     - launches other modules
* imgdisplay        - Displays an image

### Processors

* clipboardProcessor - Choose an action based on clipboard contents
* relay              - Connects two computers (network bus)

## Other modules

Not all modules fall into these easy categories, and sometimes it is better to keep related modules together.

### Pitch

* detect/pitchDetect    - Listens on the microphone and outputs note pitch
* pitchWrapper      - Wraps pitchDetect
* pitchProcessor    - Filters pitch results and outputs notes
* noteKeyboard      - Turns notes into keypresses and sends them to the active window

### Misc

* server            - The message bus daemon
* svarmr.server     - Avahi service definition file

# Helpers

Svarmr relies on a lot of other projects to provide cross platform features.  As a result, I package and ship the following programs with svarmr.

* Notifu - Popup messages on windows
    * https://www.paralint.com/projects/notifu/download.html
* Autohotkey - Intercept keys, and insert keys
    * http://autohotkey.com
* Deluge - Download torrents (linux)
    * http://deluge-torrent.org/
* fswatch - Detect filesystem changes
    * https://github.com/emcrisostomo/fswatch
* imagesnap - Take a picture
    * http://iharder.sourceforge.net/current/macosx/imagesnap/
* SyntaxNet - Sentence tagging
    * https://github.com/tensorflow/models/tree/master/syntaxnet


