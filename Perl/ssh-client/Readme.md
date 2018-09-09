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
