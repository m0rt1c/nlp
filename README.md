# chromium-netlog-parser
Command line parser for chromium netlog

## Feature: Extract all the sources from the netowrk
```bash
nlp -f /path/to/netlog.json
> e src out
```
Will extract all the sources passed on the network to the ./out folder.

## NetLog
[NetLog](https://www.chromium.org/developers/design-documents/network-stack/netlog) is chrome network logging system

## Obtaining log

### From browser
Visit [chrome://net-export](chrome://net-export)

### From command line
```bash
chromium --log-net-log=/tmp/netlog.json --net-log-capture-mode=Everything
```
The netlog capture mode flag is necessary to extract the sources.

## Other Examples

### Show help
```bash
nlp
> help
```

### Show the event with id 1
```bash
nlp -f /path/to/netlog.json
> show id 1
```

### Find all dns requests
```bash
nlp -f /path/to/netlog.json
> extract dns
```
