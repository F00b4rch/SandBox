#!/usr/bin/env perl

package Emp;
sub dns {
   my $class = shift;
   my $self = {
      domain => shift,
      A  => shift,
   };

   bless $self, $class;
   return $self;
}

sub TO_JSON { return { %{ shift() } }; }

package main;
use strict;
use warnings;

use Net::DNS::Resolver;
use Mojolicious::Lite;

# GET /
get '/' => sub {
  my $c = shift;
  $c->render(text => 'Hello World!');
};

# GET /dnsrecord?domain=domain.tld
get '/dnsrecord' => sub {
  my $c    = shift;
  my $domain = $c->param('domain');

  my $res = Net::DNS::Resolver->new(
    nameservers => [qw(8.8.8.8)],
  );
  my $query = $res->query($domain, 'A');

  if ($query) {
    foreach my $rr ($query->answer) {
        next unless $rr->type eq "A";
        my $resA = $rr->address;
        my $e = dns Emp("$domain", "$resA");
        $c->render(json =>$e);
    }

  }
  else {

    $c->render(json => "Error: no valid domain name found");
    }
};

app->start;