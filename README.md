[![Build Status](https://travis-ci.org/Ragnaroek/format.svg?branch=master)](https://travis-ci.org/Ragnaroek/format)
[![codecov](https://codecov.io/gh/Ragnaroek/format/branch/master/graph/badge.svg)](https://codecov.io/gh/Ragnaroek/format)
[![dependencies](https://img.shields.io/badge/dependencies-0-green)]()

[![playground](https://img.shields.io/badge/playground-ready-blue)](https://ragnaroek.github.io/format/)

# format

This library has the goal to bring the Common Lisp format directive to Go. This is work-in-progress, see the summary implementation table below for an overview on what is working and what not.

For a nice introduction to the Common Lisp format see https://en.wikipedia.org/wiki/Format_(Common_Lisp).

<TODO: Example how it looks in Go>

<TODO Add Playground ref here with image>

<TODO: Summary table implementation status with link to original table>

Implemented:

|~ |Name     |Prefix args       |:       |@        |:@        |Note                  |
|--|---------|------------------|--------|---------|----------|----------------------|
|[~c](http://www.lispworks.com/documentation/HyperSpec/Body/22_caa.htm) |Char     |                  |Pretty  |Escape   |          |:@ not yet implemented|
|[~%](http://www.lispworks.com/documentation/HyperSpec/Body/22_cab.htm)|Newline  |# newline         |
|[~&](http://www.lispworks.com/documentation/HyperSpec/Body/22_cac.htm)|Freshline|# lines           |
|[~`\|`](http://www.lispworks.com/documentation/HyperSpec/Body/22_cad.htm)|Page  |# pages           |
|[~~](http://www.lispworks.com/documentation/HyperSpec/Body/22_cae.htm)|Tilde    |# ~               |