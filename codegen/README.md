# Code Generation

## About

This is a rewrite of [Riven](https://github.com/MingweiSamuel/Riven/)'s code generation, the "dot-gen" project is licensed under the GPLv2 license and so is this rewrite.

You can find the source code for Riven [here](https://github.com/MingweiSamuel/Riven/) and the code for the "dot-gen" project [here](https://github.com/MingweiSamuel/Riven/tree/v/2.x.x/riven/srcgen).

In addition, the specs that are downloaded and used are from [riotapi-schema](https://github.com/MingweiSamuel/riotapi-schema) project.

## Changes

Changes had to be made to translate the Rust output to Golang (I miss Option<>), mainly with how the clients are organized (I miss impl).

Changes includes all libraries used, such as pongo2 for templating, strcase for case conversion and gjson for navigating through the JSON files.

## Todo

- Improve flow, currently a lot of functions are being reused and edge cases that might pop up can be really annoying to fix.
- Generating code may allow for some performance improvements.

## Usage

First, install `betteralign` and `goimports`:

```bash
go install github.com/dkorunic/betteralign/cmd/betteralign@latest && go install golang.org/x/tools/cmd/goimports@latest`
```

Initialize a `go.work` file in the root of the repository:
```bash
go work init
go work use .
go work use codegen
```

To update and generate code, you can either:

```bash
# from the root of equinox
UPDATE_SPECS=1 go generate
# or from inside the codegen folder
UPDATE_SPECS=1 go run .
```
