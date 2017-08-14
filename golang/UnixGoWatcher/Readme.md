## Installation

Pour build : 

```
env GOOS=linux go build -v
```

Envoyez :

```
scp '-P123' UnixGoWatch user@host:/path/to/put
```

Cron :

```
#Golang Unix Watcher
* * * * * /path/to/bin
```
