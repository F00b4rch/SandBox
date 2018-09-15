#F00b4rch
FROM perl

##################
# ENV HOST=''
# ENV USERNAME=''
# ENV PASSWORD=''
# ENV CMD=''
##################

RUN cpanm Net::SSH::Perl 
COPY . /usr/src/perl/
WORKDIR /usr/src/perl/
CMD [ "perl", "./ssh-client.pl" ]
