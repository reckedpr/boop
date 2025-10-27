# boop

is a more featured cli tool to replace them `python3 -m http.server`'s

for when you need to quickly transfer or exfiltrate data

#### examples

```bash
# serve the current directory
boop

# serve a different directory
boop /home/reckedpr

# serve from stdin (as plaintext)
sudo journalctl -u ssh | boop

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

#### alternatives
\#humble
probably better alternatives out there; this was more of a personal project..

|lang|link|
|-:|-|
|JS| [http-party/http-server](https://github.com/http-party/http-server) |
|Go| [projectdiscovery/simplehttpserver](https://github.com/projectdiscovery/simplehttpserver) |
|Python| [sc0tfree/updog](https://github.com/sc0tfree/updog) |
|Go| [eliben/static-server](https://github.com/eliben/static-server)