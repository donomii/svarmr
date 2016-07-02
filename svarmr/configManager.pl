#!/usr/bin/perl
use IO::Socket::INET;
use JSON; # imports encode_json, decode_json, to_json and from_json.

 $sock = new IO::Socket::INET (
     PeerHost => '127.0.0.1',
     PeerPort => '4816',
     Proto => 'tcp',
     Reuse => 1
 );

sub readConfig {
    my $file = 'config.json';
    open my $fh, '<', $file or die;
    $/ = undef;
    my $data = <$fh>;
    close $fh;
    decode_json($data);
}

sub newMessage {
    my ($selector, $Arg, $NamedArgs) = @_;
    my $struct = { Selector => $selector, Arg => $Arg, NamedArgs => $NamedArgs};
    encode_json($struct);
}

while ( my $line = <$sock> ) {
    chomp $line;
    my $s = {Selector => "Decode failed"};
    eval { $s = decode_json($line)};
    #print "** ".$s->{Selector}." **\n\n".$s->{Arg}."\n\n\n";
    if ($s->{Selector} eq "get-all-config") {
        print newMessage("config", "", readConfig())."!\n";
        print $sock newMessage("config", "", readConfig())."\n";
    }
}


