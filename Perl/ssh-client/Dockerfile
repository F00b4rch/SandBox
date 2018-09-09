#F00b4rch
FROM perl
RUN cpanm Net::SSH::Perl 
COPY . /usr/src/perl/
WORKDIR /usr/src/perl/
CMD [ "perl", "./ssh-client.pl" ]
