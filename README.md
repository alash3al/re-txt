re-txt
==========
> reformates a text file from a structure to another, i.e: convert from json to yaml, toml to json, ... etc

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
$ re-txt --src example.json json2yaml

# json to yaml using stdin pipe
$ cat example.json | re-txt json2yaml

# csv to yaml by first converting to json
$ cat example.csv | re-txt csv2json | re-txt json2yaml

# csv to yaml by first converting to json
# also merge multiple json files and convert the piped result & them into yaml
$ cat example.csv | re-txt csv2json | re-txt --dest ./result.yaml --src=another1.json --src=another2.json json2yaml
```

Installations
==============
- from source: `go get github.com/alash3al/re-txt`
- binary download: go to [there](https://github.com/alash3al/re-txt/releases) and download the binary which support your env