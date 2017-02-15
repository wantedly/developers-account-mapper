# developers-account-mapper

[![Build Status](https://travis-ci.org/wantedly/developers-account-mapper.svg)](https://travis-ci.org/wantedly/developers-account-mapper)
[![GitHub release](https://img.shields.io/github/release/wantedly/developers-account-mapper.svg)](https://github.com/wantedly/developers-account-mapper/releases)

## Summary

Manage developers' accounts list.

```bash
$ developers-account-mapper to-slack-mention potsbo
<@U2XXXXXXX|shimpei>
```
## Description

## Usage

One DynamoDB tables named `DevelopersAccountMap` has to be created.

#### `DevelopersAccountMap` table

|Key|Type| |
|---|----|---|
|LoginName|String|Primary key|

In addition, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `AWS_REGION` must be set at your shell.
This IAM user/role must be allowed to read/write the DynamoDB table above.

## Install

To install, use `go get`:

```bash
$ go get -d github.com/wantedly/developers-account-mapper
```

## Contribution

1. Fork ([https://github.com/wantedly/developers-account-mapper/fork](https://github.com/wantedly/developers-account-mapper/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[wantedly](https://github.com/wantedly)
