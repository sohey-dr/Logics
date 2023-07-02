# dependencies: Mojolicious

use strict;
use warnings;
use Mojolicious::Lite;
use Mojo::JSON qw(decode_json encode_json);

get '/' => sub { shift->render(text => 'Hello World!') };
get '/json' => sub {
	my $self = shift;
	my $json = { name => 'mojo', version => '1.0' };
	$self->render(json => $json);
};

app->start;
