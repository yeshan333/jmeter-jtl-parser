# Apache-Jmeter JTL Parser

> Stream decode JMeter JTL files and output them as XML, CSV, HTTP JSON.

Parse JMeter JTL files, supporting:

- Nested samples
- JMeter custom variables
- Responses
- Assertions
- Cookies
- More...

The parser is a **stream decoder**, meaning it's safe to use for very large files. Also you can use it to merge jtl files and post jtl's json data to a http server for display your test report.

## Usage

require Go 1.12+.

```bash
# build the binary
$ make
# merge jtl files. merge a_create_data.jtl and b_platform_ci.jtl to data.jtl, the final file is test/merged.jtl.
$ ./jtl-parser data.jtl -o xml -f a_create_data.jtl b_platform_ci.jtl > test/merged.jtl

$ go run parser.go -h
Usage: parser [--output OUTPUT] [--mergedfiles MERGEDFILES] [--httphost HTTPHOST] FILENAME

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         specify the output type, valid options: csv,xml,json,http [default: http]
  --mergedfiles MERGEDFILES, -f MERGEDFILES
                         merge jtl files
  --httphost HTTPHOST, -h HTTPHOST
                         post jtl json data to http server [default: http://localhost:7167/jmeter]
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

- HTTP JSON

## Thanks

- [Gilad Peleg](https://www.giladpeleg.com)
