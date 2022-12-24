# Jmeter JTL Parser

> Stream decode JMeter JTL files and output them as XML, CSV

Parse JMeter JTL files, supporting:

- Nested samples
- JMeter custom variables
- Responses
- Assertions
- Cookies
- More...

The parser is a **stream decoder**, meaning it's safe to use for very large files

## Usage

```bash
❯ ./jtl-parser -h
Usage: jtl-parser [--output OUTPUT] [--mergedfiles MERGEDFILES] [--reportapihost REPORTAPIHOST] FILENAME

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         specify the output type, valid options: csv,xml,json,http [default: http]
  --mergedfiles MERGEDFILES, -f MERGEDFILES
                         merge jtl files
  --reportapihost REPORTAPIHOST, -h REPORTAPIHOST [default: http://localhost:8080/jmeter]
  --help, -h             display this help and exit
```

## Parse Validating

Output/Input XML is not in canonical format, in order to diff them,
use the included `compare.sh` script:

```bash
$ ./compare.sh data.jtl
```

## Output Formats

- XML

- JSON

- CSV

- HTTP

## Thanks

- [Gilad Peleg](https://www.giladpeleg.com)
