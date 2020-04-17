# Mackerel Project validator

This tool validates your Mackerel project definition files.

The schema is provided as [JSON Schema](https://json-schema.org/).

## How to use

The validator is a binary which is available for different package managers.

You'll find example project 

Find valid and invalid files in examples/ folder.
(Note that examples/invalid/docker_double-expose.yml currently validates... :/)

In case you're looking for the source code of the validator, it's in src.

### Standalone

```shell script
./bin/project-validator path/to/definition.yml
```

### [PyPI (Python)](https://pypi.org/project/mackerel-project-validator)

```shell script
pip install mackerel-project-validator
```

Binary will be installed into `...`.

### Planned packages

* pypi (Python)
* Composer (PHP)
* NPM (NodeJS)