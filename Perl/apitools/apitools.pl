#!/usr/bin/env perl
use strict;
use warnings;

use Net::DNS::Resolver;
use Mojolicious::Lite;

get '/' => sub {
  my $c = shift;
  $c->render(text => 'Hello World!');
};

# /Adnsrecord?domain=domain.tld
get '/Adnsrecord' => sub {
  my $c    = shift;
  my $domain = $c->param('domain');
  
  my $res = Net::DNS::Resolver->new(
    nameservers => [qw(8.8.8.8)],
  );
  
  my $query = $res->search($domain);
  
  if ($query) {
  foreach my $rr ($query->answer) {
    next unless $rr->type eq "A";
    $c->render(text => "$domain A record : ".$rr->address);
    }
  }
  else {
    $c->render(text => "$domain no A DNS record found.");
    }
};

app->start;