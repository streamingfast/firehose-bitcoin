# Change log

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this
project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). See [MAINTAINERS.md](./MAINTAINERS.md)
for instructions to keep up to date.

## Unreleased

### Added

* added flag `--headers` (repeatable) to allow setting HTTP headers, e.g. "Authorization: bearer aGVsbG8K"
* added support for addresses prefixed by `http://` or `https://`

## v1.0.0-rc.1

### StreamingFast Firehose Bitcoin implementation

* Bitcoin blocks are fully compatible with the RPC blocks
* Bitcoin Block contains all transaction data, including `vins` and `vouts`*
