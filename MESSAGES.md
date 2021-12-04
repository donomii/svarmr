# Messages

The message is the important part of svarmr, and the only part that is formally documented.  Servers and nodes are free to do whatever they like with a message, but messages must in the following form.

## Transmission

Messages are sent as JSON objects, one per line.  Each line must end with a \n only, although most clients will helpfully ignore a \r.

## Format

Each message is a JSON object, that has some of the following fields:

{
    Selector: "name",               // The "method" to call.  "object" or "server" is specified elsewhere
    Arg: "any data in a string",    // Can be a stringified JSON, Base64 data, whatever... mainly needed for static-typed languages that have trouble parsing JSON that contains unpredictable keys (looking at Golang...)
    NamedArgs: { A JSON hash } // For language that can easily deal with free-form JSON
}

## Use

### Selector
    Chooses the code to run.  Selectors are strings like "record-audio", "displayMessage", or "load_file".  Messages are broadcast to each node, which then chooses whether or not to do something with the message.

### Arg

A JSON string.  It was the only option during bootstrapping, and is still useful when talking to nodes written in languages that suck at handling JSON structures that can change unpredictably (most typed languages have this problem).  You should use this field whenever possible, since it allows primtive clients to understand your messages.  e.g. a "Clock" node should broadcast the time like:

    {
        "Selector": "announce-time",
        "Arg":      "12:15"
    }

rather than

    {
        "Selector": "announce-time",
        "NamedArgs":   { "time": "12:15" }
    }

since the second one is unreasonably difficult to parse in C.

There are no restrictions on what you put into the string, and there are already nodes which send an entire PNG in the Arg field.

### NamedArgs

Mainly intended for making complex calls with lots of parameters.  e.g.

    {
        "Selector": "create-panel",
        "NamedArgs": {
                        "width": 100,
                        "height": 100,
                        "colour": "AA0011",
                        "label": "test panel",
                        ... etc ...
                    }
    }


## Proposed fields

In the future, the following fields will probably be used:

### Target

Chooses the object to handle the message, or possibly a "type" of object so multiples can answer?

### MessageId

A unique identifier for the message

### InResponseTo

Quotes a previous MessageID, allowing a node to make a call and identify the return value



