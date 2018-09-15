### How to use
First build the image :
`$ docker build -t perlapitools .`

Then run :
```
$ docker run --name running-perlapitools -p 8080:8080 perlapitools
```

Debug with :

`$ docker run -it --rm perlapitools /bin/bash`