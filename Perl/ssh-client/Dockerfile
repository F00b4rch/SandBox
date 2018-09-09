#F00b4rch
FROM perl

##################
# REPLACE ME
ENV HOST xx.xx.xx.xx
ENV USERNAME user
ENV PASSWORD pass
ENV CMD uptime
##################

RUN cpanm Net::SSH::Perl 
COPY . /usr/src/perl/
WORKDIR /usr/src/perl/
CMD [ "perl", "./ssh-client.pl" ]
