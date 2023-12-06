# Code Generation

## About

This project is based on [Riven](https://github.com/MingweiSamuel/Riven/)'s code generation, the "dot-gen" license appears to be GPL-2.0 so hopefully I am following it properly.

### Changes

Added `node-fetch-commonjs`, `prettier` and `eslint`.

Removed `glob-promise`, `glob`, `request`, `request-promise-native`.

Changes had to be made to translate the rust output to golang (I miss Option<>), mainly with how the clients are organized (I miss impl).

Separated the download and compilation process in separate scripts.

## Usage

> `npm run compile` uses `betteralign` and `goimports`, make sure to install both with `go install github.com/dkorunic/betteralign/cmd/betteralign@latest && go install golang.org/x/tools/cmd/goimports@latest`.

After installing the dependencies with `npm i`, run `npm run update` and then `npm run compile`.

If you are running multiple go and node commands, you can use `npm run compile --prefix srcgen` at the project root, so you don't need to keep changing directories or have multiple terminals open.

Sometimes you will need to run `npm run align` after `npm run compile` because some structs won't be aligned on `compile`.

A normal workflow would be:

```bash
npm run update
npm run compile
npm run align
```
