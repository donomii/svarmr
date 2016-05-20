# svarmr

A networked control bus and messaging system

* Link the volume on all your devices
* Automatically start torrents on a different machine
* Trigger on custom events
* Easy access to platform libraries

## Easy to add modules

Svarmr is a simple message bus that is super easy to write new modules for.

Simple modules can be a few lines long, and there are examples in Go, Racket, Perl, and C.

## Useful modules

Svarmr already has some useful modules, to do things like monitor clipboard changes, report mDNS events and broadcast keys.

## Design

### Sources, sinks and processors

Each module can transmit and receive any events it pleases, but it is useful to break them into categories of modules that mainly generate events (sources), modules that wait for events and then do something (sinks), and modules that wait for events, process them and send new events (processors).

#### Sources

* clipboardWatcher  - Broadcasts clipboard changes
* mdnsWatcher       - Broadcasts mDNS results
* heartBeat         - Broadcasts a regular beat

#### Sinks

* torrentListener   - Starts torrents in Deluge
* monitor.c         - Prints events to stdout
* monitor.go        - Prints events to stdout
* volume            - Sets the volume (needs helpers)
* userNotify        - Pops up a message on the screen (needs helpers)
* moduleStarter     - launches other modules

#### Processors

* clipboardProcessor- Choose an action based on clipboard contents


#### Other

* server            - The message bus daemon
* svarmr.server     - Avahi service definition file
* relay             - Connects two computers (network bus)


