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
|[~c](/https://www.hexstreamsoft.com/articles/common-lisp-format-reference/format-directives/#~C_character) |Char     |                  |Pretty  |Escape   |          |:@ not yet implemented|
|[~%](https://www.hexstreamsoft.com/articles/common-lisp-format-reference/format-directives/#~percent_newline)|Newline  |# newline         |
|[~&](https://www.hexstreamsoft.com/articles/common-lisp-format-reference/format-directives/#~ampersand_fresh-line)|Freshline|# lines           |
|[~`\|`](https://www.hexstreamsoft.com/articles/common-lisp-format-reference/format-directives/#~vertical-line_page)|Page  |# pages           |
|[~~](http://www.lispworks.com/documentation/HyperSpec/Body/22_cae.htm)|Tilde    |# ~               |