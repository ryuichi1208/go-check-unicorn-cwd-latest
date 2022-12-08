# go-check-unicorn-cwd-latest

This tool uses a symbolic link deployment method to make sure the process has the latest directory.

## Usage


Checks whether the cwd of the pid that matches the process name specified with -p matches the latest directory among the directories specified with -d.

```
Usage:
 check-unicorn-cwd-latest [OPTIONS]

Application Options:
  -p=         process name
  -d=         release directory

Help Options:
  -h, --help  Show this help message
```
