any-to-any
==========
> convert anything to anything

Supported Source Formats
=========================
- json
- yaml
- hcl
- toml
- csv

Supported Target Formats
=========================
- json
- yaml
- toml

Examples
=========

```bash
# json to yaml using flag
$ a2a --src example.json json2yaml

# json to yaml using stdin pipe
$ cat example.json | a2a json2yaml

# csv to yaml by first converting to json
$ cat example.csv | a2a csv2json | a2a json2yaml

# csv to yaml by first converting to json
# also merge multiple json files and convert the piped result & them into yaml
$ cat example.csv | a2a csv2json | a2a --dest ./result.yaml --src=another1.json --src=another2.json json2yaml
```

Installations
==============
- from source: `go get github.com/alash3al/a2a`
- binary download: go to [there](https://github.com/alash3al/a2a/releases) and download the binary which support your env