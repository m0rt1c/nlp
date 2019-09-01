# chromium-netlog-parser
Command line parser for chromium netlog

## NetLog
[NetLog](https://www.chromium.org/developers/design-documents/network-stack/netlog) is chrome network logging system

## Obtaining log

### From browser
Visit [chrome://net-export](chrome://net-export)

### From command line
```bash
chromium --log-net-log=/tmp/netlog.json --net-log-capture-mode=Everything
```

## Examples

### Show help
```bash
npl
> help
```

### Show the event with id 1
```bash
nlp -file /path/to/netlog.json
> show id 1
```

### Find all dns requests
```bash
nlp -file /path/to/netlog.json
> extract dns
```
