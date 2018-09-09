#!/bin/perl
# dev : f00b4rch
# description : First Perl program

# Use part
use warnings;
use strict;
use Net::SSH::Perl;

# Starting
print("[INFO] Initialize connexion...\n");

=begin
Please see full documentation here : 
- https://metacpan.org/pod/Net::SSH::Perl

Exemple vars :
$host       // IP-host
$username   // ssh username
$password   // ssh password
$port       // ssh port
=cut

my $host = $ENV{'HOST'};;
my $username = $ENV{'USERNAME'};
my $password = $ENV{'PASSWORD'};

my $ssh = Net::SSH::Perl->new($host);
print("[OK] Connexion success!\n[INFO]Begin authentification...\n");

$ssh->login($username,$password);

# Defining 
my $command = "w";
my($stdout, $stderr, $exit) = $ssh->cmd('$command');
# Checking error
if ($stderr ne "")
{
    print('Error while executing $command : $stderr');
}
print $stdout;
print("[INFO] Exit.")
exit;
