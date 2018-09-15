#!/usr/bin/env perl

package Emp;
sub new {
   my $class = shift;
   my $self = {
      name => shift,
      dnsrecord  => shift,
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

# GET /Adnsrecord?domain=domain.tld
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
        my $res = $rr->address;
        my $e = new Emp( "$domain", "$res");
        $c->render(json =>$e);
    }

  }
  else {
    $c->render(json => "Error: no valid domain name found");
    }
};

app->start;