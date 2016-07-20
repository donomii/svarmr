# Writing a svarmr client

Svarmr clients are extremely simple.  The protocol is designed to be as easy as possible to work with.  In a major programming language, you can get a client working, from scratch, in under 10 minutes.  Here's how.

## The protocol

Svarmr clients send each other JSON messages.  Each message ends with a newline.  

### Listener client

So to make a client that just listens to messages, you will need to:


    connect to the svarmr server over TCP
    loop
        readline from network
        decode json from the string
        ... do something with the json ...

And that's all you need for a basic client.  This sort of client can do a lot, and there are several useful clients in the svarmr package that work like this.  Monitor is one, and so is imageDisplay.

But soon you will want to send messages as well

### Transmitter client


