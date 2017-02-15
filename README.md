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

Convert a key username, which is recommended to be unique among the developers group, to other username.
Currently, we support Slack and GitHub.

### Slack

Especially abount Slack, developers-account-mapper is able to fetch Slack ID by Slack username and output it as a Slack mention format.

## Usage

One DynamoDB tables named `DevelopersAccountMap` has to be created.

#### `DevelopersAccountMap` table

|Key|Type| |
|---|----|---|
|LoginName|String|Primary key|

In terraform, you can create the table with below.

```tf
resource "aws_dynamodb_table" "developers-account-map" {
  name           = "DevelopersAccountMap"
  read_capacity  = 5 # or anything you want
  write_capacity = 5 # or anything you want
  hash_key       = "LoginName"
  attribute {
    name = "LoginName"
    type = "S"
  }
}
```

In addition, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_REGION` and `SLACK_API_TOKEN` must be set at your shell.
This IAM user/role must be allowed to read/write the DynamoDB table above.

### Command Usage

```
usage: developers-account-mapper [--version] [--help] <command> [<args>]

Available commands are:
    delete              Delete record with <login_name>
    exec                Set account information as env vars and exec commands
    list                List mapping of <login_name> and mapped accounts
    register            Register LoginName and other accounts mapping
    to-github-name      Get <github_username> from <login_name>
    to-slack-mention    Get <slack_mention> from <login_name>
    version             Print developers-account-mapper version and quit
```

### Use with Docker

#### Run

```
docker run --rm \
  -e SLACK_API_TOKEN=<slack token get by https://api.slack.com/docs/oauth-test-tokens>  \
  -e AWS_ACCESS_KEY_ID=yourawsaccesskeyid \
  -e AWS_SECRET_ACCESS_KEY=yourawssecretaccesskey \
  -e AWS_REGION=ap-northeast-1 \
  quay.io/wantedly/developers-account-mapper \
  <command>
```

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

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
