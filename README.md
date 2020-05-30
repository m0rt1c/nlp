# chromium-netlog-parser
Command line parser for chromium netlog

## NetLog
[NetLog](https://www.chromium.org/developers/design-documents/network-stack/netlog) is chrome network logging system

## Obtaining log

### From browser
Visit [chrome://net-export](chrome://net-export)

### From command line
You can collect the netlog with either google-chrome or chromium

```bash
chromium --log-net-log=/tmp/netlog.json --net-log-capture-mode=Everything
```
```bash
google-chrome-stable --log-net-log=/tmp/netlog.json --net-log-capture-mode=Everything
# <any-google-chrome-binary> --log-net-log=/tmp/netlog.json --net-log-capture-mode=Everything
```
The netlog capture mode flag is necessary to extract the sources.

## Usage Examples

### Show the event with id 1
```bash
nlp -f /path/to/netlog.json
> show id 1
```

```bash
nlp -f /path/to/netlog.json -c 'show id 1'
```

### Extract all the sources from the netowrk
Will extract all the website sources that were downloaded into the ./out folder. The folder can be changed.
```bash
nlp -f /path/to/netlog.json
> e src out
```
```bash
nlp -f /path/to/netlog.json -c 'e src out'
```

### Find all dns requests
```bash
nlp -f /path/to/netlog.json
> extract dns
```

```bash
nlp -f /path/to/netlog.json -c 'extract dns'
```

### Find all URLs accessed
```bash
nlp -f /path/to/netlog.json
> e url
```

```bash
nlp -f /path/to/netlog.json -c 'e url'
```

### Find all the TCP and UDP connections opened
```bash
nlp -f /path/to/netlog.json
> e con
```

```bash
nlp -f /path/to/netlog.json -c 'e con'
```

### Show help
```bash
nlp
> help
```

## Building and Installing

### Obtaing the code
```bash
go get github.com/AndreaJegher/nlp
```
The repo with the code will be created in your default go source folder

### Installing
```bash
go install github.com/AndreaJegher/nlp
```
If you have the default go binary folder in your path, you can start the program by using the `nlp` command

### Building
```bash
go get github.com/AndreaJegher/nlp
cd $GOPATH/src/github.com/AndreaJegher/nlp
make
```
