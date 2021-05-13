# myhttp

## Description

A service dedicated to hash link using md5 hash

## Usage
Example:
```
cd myhttp
./myhttp --parallel 3 google.com reddit.com/r/funny
```
--parallel stands for the number of worker, the default is 10

## Development Guide
### Pre-requisite
- [Go 1.16](https://golang.org/doc/install)

### Installation
```
git clone https://github.com/danielbintar/myhttp.git
cd myhttp
```

### Contribute
- Create New Pull Request
- Run command below
```
make pretty
make test
make compile
```

