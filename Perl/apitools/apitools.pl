#!/usr/bin/env perl
package main;
use strict;
use warnings;

use JSON::Create 'create_json';
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

  # Defining domain hash
  my %domainhash = (
        name => shift,
        dnsrecord => shift,
  );

  my $res = Net::DNS::Resolver->new(
    nameservers => [qw(8.8.8.8)],
  );
  my $query = $res->search($domain);

  if ($query) {
    foreach my $rr ($query->answer) {
        next unless $rr->type eq "A";
        my $res = $rr->address;
        $domainhash{name} = "$domain";
        $domainhash{dnsrecord} = "$res";
        my $result = create_json(\%domainhash);
        $c->render(json =>$result);
    }

  }
  else {
    $c->render(json => "Error: no valid domain name found");
    }
};

app->start;