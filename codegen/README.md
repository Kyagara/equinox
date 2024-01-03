# Code Generation

## About

This is a rewrite of [Riven](https://github.com/MingweiSamuel/Riven/)'s code generation, the "dot-gen" project is licensed under the GPLv2 license and so is this rewrite.

You can find the source code for Riven [here](https://github.com/MingweiSamuel/Riven/) and the code for the "dot-gen" project [here](https://github.com/MingweiSamuel/Riven/tree/v/2.x.x/riven/srcgen).

In addition, the specs that are downloaded and used are from [riotapi-schema](https://github.com/MingweiSamuel/riotapi-schema) project.

### Changes

Changes had to be made to translate the Rust output to Golang (I miss Option<>), mainly with how the clients are organized (I miss impl).

Changes includes all libraries used, such as pongo2 for templating, strcase for case conversion and gjson for JSON.

### Todo

- Add checks for required fields (queries and headers), returning errors if they are missing, for now only Authorization headers are required so it's not a big deal.
- Generating code may allow for some performance improvements.

## Usage

First, install `betteralign` and `goimports`:

```bash
go install github.com/dkorunic/betteralign/cmd/betteralign@latest && go install golang.org/x/tools/cmd/goimports@latest`
```

To generate code, run from the root of the `equinox` project:

```bash
go generate ./...
```

To update the specs, you can do either:

```bash
# from the root of equinox
UPDATE_SPECS=1 go generate ./...
# or from inside the codegen folder
go run . -update
```
