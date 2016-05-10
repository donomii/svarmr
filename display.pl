use IO::Socket::INET;
 use JSON; # imports encode_json, decode_json, to_json and from_json.

 $sock = new IO::Socket::INET (
     PeerHost => '127.0.0.1',
     PeerPort => '4816',
     Proto => 'tcp',
     Reuse => 1
 );


while ( my $line = <$sock> ) {
    my $s = decode_json($line);
    warn "** ".$s->{Selector}." **\n\n".$s->{Arg}."\n\n\n";
}


