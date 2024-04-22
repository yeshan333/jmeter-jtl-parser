# Jmeter JTL Parser

> Stream decode JMeter JTL ([xml-format](https://jmeter.apache.org/usermanual/listeners.html#xmlformat2.1)) files and output them as XML, CSV, JSON, HTTP JSON.

Parse JMeter [xml-format](https://jmeter.apache.org/usermanual/listeners.html#xmlformat2.1) JTL files to "anything", supporting:

- Nested samples
- JMeter custom variables
- Responses
- Assertions
- Cookies
- More...

The parser is a **stream decoder**, meaning it's safe to use for very large jtl files.

## Usage

```bash
❯ ./bin/jtl-parser-linux -h
Usage: jtl-parser-linux [--output OUTPUT] [--mergedfiles MERGEDFILES] [--reportcenterapihost REPORTCENTERAPIHOST] FILENAME

Positional arguments:
  FILENAME

Options:
  --output OUTPUT, -o OUTPUT
                         specify the output type, valid options: csv,xml,json,http [default: http]
  --mergedfiles MERGEDFILES, -f MERGEDFILES
                         merge jtl files
  --reportcenterapihost REPORTCENTERAPIHOST, -h REPORTCENTERAPIHOST
                         transfer jtl data to JSON, and post it to report center. [default: http://localhost:8080/jmeter]
  --help, -h             display this help and exit
```

## Examples

```shell
git clone git@github.com:yeshan333/jmeter-jtl-parser.git
cd jmeter-jtl-parser
# compile bin file
make
# transfer to csv
./bin/jtl-parser-linux --output csv tests/github-etcd.jtl
# transfer to json
./bin/jtl-parser-linux --output json tests/github-etcd.jtl
```

## Output Formats

- XML

- JSON

- CSV

- HTTP JSON: for custom test report center show test data

## Parse Validating

Output/Input XML is not in canonical format, in order to diff them,
use the included `compare.sh` script:

```bash
$ ./compare.sh tests/github-etcd.jtl
```

## Thanks

- [Gilad Peleg](https://www.giladpeleg.com)
