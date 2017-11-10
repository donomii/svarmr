use strict;
use IO::Socket::INET;
 use JSON; # imports encode_json, decode_json, to_json and from_json.

my $method = shift;
die "First arg must be the connection method.  Either server:port or 'pipes'" unless $method;

if ($method eq 'pipes') {
	
while ( my $line = <> ) {
    my $s = decode_json($line);
    print "** ".$s->{Selector}." **\n\n".$s->{Arg}."\n\n\n";
}
	} else {
		my ($server, $port) = split/:/, $method;
$sock = new IO::Socket::INET (
     PeerHost => $server,
     PeerPort => $port,
     Proto => 'tcp',
     Reuse => 1
 );


while ( my $line = <$sock> ) {
    my $s = decode_json($line);
    print "** ".$s->{Selector}." **\n\n".$s->{Arg}."\n\n\n";
}


}