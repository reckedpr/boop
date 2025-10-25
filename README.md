# boop

small tool to replace them quick `python3 -m http.server`'s

#### examples:

```bash
# serve the current directory
boop

# serve a differant directory
boop /home/reckedpr

# serve output directly from stdin..
echo hai | boop

# ..or a files contents for example
cat docker-compose.yml | boop
```

#### install

```bash
go install https://github.com/reckedpr/boop/cmd/boop@latest
```