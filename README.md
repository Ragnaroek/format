[![Build Status](https://travis-ci.org/Ragnaroek/format.svg?branch=master)](https://travis-ci.org/Ragnaroek/format)
[![codecov](https://codecov.io/gh/Ragnaroek/format/branch/master/graph/badge.svg)](https://codecov.io/gh/Ragnaroek/format)
[![dependencies](https://img.shields.io/badge/dependencies-0-green)]()

[![playground](https://img.shields.io/badge/playground-ready-blue)](https://ragnaroek.github.io/format/)

# format

This library has the goal to bring the Common Lisp format directive to Go. This is work-in-progress, see the summary implementation table below for an overview on what is working and what not.

For a nice introduction to the Common Lisp format see https://en.wikipedia.org/wiki/Format_(Common_Lisp).

## Example Code

```go
import "github.com/ragnaroek/format/pkg"

ft.Sformat("~8r", 8) //returns "10"
ft.Sformat("~10,'#,'🥭,2:@X", 4099) //returns "####+10🥭03"

ft.Sformat("~r", 4343637058903381868) //returns "four quintillion three hundred forty-three quadrillion six hundred thirty-seven trillion fifty-eight billion nine hundred three million three hundred eighty-one thousand eight hundred sixty-eight"
ft.Sformat("~:@r", 2799) //returns "MMDCCLXXXXVIIII"
ft.Sformat("~@r", 2799) //returns "MMDCCXCIX"
```

## Playground

Since the format directives can get complicated and the best way to figure them output is to play around with them the `format` playground was created:

[![playground.png](https://i.postimg.cc/wx0qgqWG/playground.png)](https://postimg.cc/xqNDP2Vv)

The playground is hosted here: https://ragnaroek.github.io/format/
The library is compiled to WASM to make it run on the browser.

## Implementation Status

The directives listed below are already implemented:

|~ |Name     |Prefix args       |:       |@        |:@        |Note                  |
|--|---------|------------------|--------|---------|----------|----------------------|
|[~c](http://www.lispworks.com/documentation/HyperSpec/Body/22_caa.htm) |Char     |                  |Pretty  |Escape   |          |:@ not yet implemented|
|[~%](http://www.lispworks.com/documentation/HyperSpec/Body/22_cab.htm)|Newline  |# newline         |
|[~&](http://www.lispworks.com/documentation/HyperSpec/Body/22_cac.htm)|Freshline|# lines           |
|[~`\|`](http://www.lispworks.com/documentation/HyperSpec/Body/22_cad.htm)|Page  |# pages           |
|[~~](http://www.lispworks.com/documentation/HyperSpec/Body/22_cae.htm)|Tilde    |# ~               |
|[~r](http://www.lispworks.com/documentation/HyperSpec/Body/22_cba.htm)|Radix|mincol, padchar, comma-char, comma-interval|Ordinal|Roman|Old Roman||
|[~d](http://www.lispworks.com/documentation/HyperSpec/Body/22_cbb.htm)|Decimal|mincol, padchar, comma-char, comma-interval|
|[~b](http://www.lispworks.com/documentation/HyperSpec/Body/22_cbc.htm)|Binary|mincol, padchar, comma-char, comma-interval|
|[~o](http://www.lispworks.com/documentation/HyperSpec/Body/22_cbd.htm)|Octal|mincol, padchar, comma-char, comma-interval|
|[~x](http://www.lispworks.com/documentation/HyperSpec/Body/22_cbe.htm)|Hexadecimal|mincol, padchar, comma-char, comma-interval|
|[~f](http://www.lispworks.com/documentation/HyperSpec/Body/22_cca.htm)|Float|width, decimals, scale, overflow, pad||Sign|||


This table is derived from https://www.hexstreamsoft.com/articles/common-lisp-format-reference/clhs-summary/#subsections-summary-table, which was also a great help in the implementation of the directives so far. Many thanks to Jean-Philippe Paradis.

All other directives not mentioned in the table are not implemented yet.

## Current work-in-progress

Implementing more directives.

Optimisation! `format` is currently ~25% times slower than `fmt`:
```
BenchmarkFmtSimple-12       	18712560	        63.4 ns/op
BenchmarkFmtLong-12         	 4741552	       263 ns/op
BenchmarkFormatSimple-12    	14139633	        81.1 ns/op
BenchmarkFormatLong-12      	 3741561	       318 ns/op
```

You can run the benchmarks yourself with `make bench`.

# Thanks

This software is open source (LGPLv3) and was made while listening to a lot of [Rage against the Machine](https://www.last.fm/music/Rage+Against+the+Machine) ✊🏿