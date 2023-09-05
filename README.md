# fsp
Frozen sprint planning helper: wrapper for gh cli

# Usage
Currently only bulk-create command is supported.

## bulk-create
bulk-creates issues from a csv file using an issue template
```
$ fsp bulk-create -h
NAME:
   fsp bulk-create - create issues from a csv file

USAGE:
   fsp bulk-create [command options] [arguments...]

OPTIONS:
   --file value, -f value                                 path to input csv file
   --labels value, -l value [ --labels value, -l value ]  labels to add
   --project-id value, -p value                           project id
   --help, -h                                             show help

```
- by supplying `--project-id` you can create and add issues to a GitHub Project
- `issue-template.md` will be used as the template, `%BODY%` will be replaced with body texts in csv file
- your csv file needs to follow the following format
```csv
title,body,repo,assignee
my issue title,my issue body,Masamerc/fsp,Masamerc
```


# Installation
Prerequisites:
- Docker

There is no need to have Go installed since the project builds the binary using Docker.
To install simply run:
```
$ make install GOOS=<your_os> GOARCH=<your_arch>
```
*default is `GOOS=darwin GOARCH=arm64`

The compiled binary will be put under `/usr/local/bin` which is in the `PATH` by default for Mac OS users. Linux users may need to adjust the `Makefile` so the binary will be in the `PATH`