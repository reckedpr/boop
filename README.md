# boop

is a more featured cli tool to replace them quick `python3 -m http.server`'s

..for everytime you need to boop something across the lan

#### examples

```bash
# serve the current directory
boop

# serve a differant directory
boop /home/reckedpr

# serve from stdin
echo hai | boop

# serve contents for 5 minutes then stop
cat docker-compose.yml | boop -t 5
```

#### install

```bash
go install github.com/reckedpr/boop/cmd/boop@latest
```

#### args

```
Usage of boop:
  -p, --port int   port to serve (default 8080)
  -t, --time int   time in minutes to serve for
```