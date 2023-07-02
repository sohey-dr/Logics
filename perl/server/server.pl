# dependencies: Mojolicious

use strict;
use warnings;
use Mojolicious::Lite;
use Mojo::JSON qw(decode_json encode_json);

get '/' => sub { shift->render(text => 'Hello World!') };

app->start;
