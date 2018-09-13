### How to use
First build the image :
`$ docker build -t perl-ssh .`

Then run :
```
$ docker run -it --rm --name running-perl-ssh \
    -e HOST='xx.xx.xx.xx' \
    -e USERNAME='user' \
    -e PASSWORD='pass' \
    -e CMD='command' \
    perl-ssh
```

Test with :

`docker run -it --rm perl-ssh /bin/bash`

You will be able to type :

`HOST='xx.xx.xx.xx' ; USERNAME='user' ; PASSWORD='pass' ; CMD='command' ; perl ssh-client.pl`
