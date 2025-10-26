# boop

is a more featured cli tool to replace them quick `python3 -m http.server`'s

..for everytime you need to boop something across the lan

#### examples

```bash
# serve the current directory
boop

# serve a different directory
boop /home/reckedpr

# serve from stdin
echo hai | boop

# serve file for 5 minutes then stop
boop docker-compose.yml -t 5

# bind to host and serve foo/
boop foo/ --host
```

#### install

```bash
go install github.com/reckedpr/boop/cmd/boop@latest
```

#### args

```
Usage of boop:
      --host       bind to host (0.0.0.0)
  -p, --port int   port to serve (default 8080)
  -t, --time int   time in minutes to serve for
```