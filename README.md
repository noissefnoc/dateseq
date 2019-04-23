dateseq
=======

[![Build Status](https://travis-ci.org/noissefnoc/dateseq.svg?branch=master)][travis]
[![Coverage Status](https://coveralls.io/repos/noissefnoc/dateseq/badge.svg?branch=master)][coveralls]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![GoDoc](https://godoc.org/github.com/noissefnoc/dateseq?status.svg)][godoc]

[travis]: https://travis-ci.org/noissefnoc/dateseq
[coveralls]: https://coveralls.io/r/noissefnoc/dateseq?branch=master
[license]: https://github.com/noissefnoc/dateseq/blob/master/LICENSE
[godoc]: https://godoc.org/github.com/noissefnoc/dateseq

generate date sequence like Linux `seq` command.

## Synopsis

```console
# in case execution date is 20190101
% dateseq 3
20190101
20190102
20190103

# specify start date as 20190201
% dateseq 2 20190201
20190201
20190202
```

## Description

## Installation

```console
% go get github.com/noissefnoc/dateseq
```

## Author

[noissefnoc](https://github.com/noissefnoc)
