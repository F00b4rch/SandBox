### How to use
First build the image :
`$ docker build -t perlapitools .`

Then run :
```
$ docker run -it --rm --name running-perlapitools perlapitools
```

Debug with :

`$ docker run -it --rm perlapitools /bin/bash`