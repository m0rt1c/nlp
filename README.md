# chromium-netlog-parser
Command line parser for chromium netlog

## NetLog
[NetLog](https://www.chromium.org/developers/design-documents/network-stack/netlog) is chrome network logging system

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
