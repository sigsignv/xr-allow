# xr-allow

A CLI utility for allowing SSH connection IPs using the XREA API.

## Install

```sh
$ go install github.com/sigsignv/xr-allow
```

## Usage

```
Usage: xr-allow [options...] 192.0.2.0
  -c, --config string   Specify config file path (default "./config.toml")
  -h, --help            Show help
  -q, --quiet           Suppress output
  -v, --version         Show version
```

## Configuration

```toml
[[servers]]
account = "USERNAME"
server_name = "SERVERNAME.xrea.com"
api_secret_key = "YOUR_API_KEY"
```

## Author

Sigsign <<sig@signote.cc>>

## License

Apache-2.0
