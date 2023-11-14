# Code Generation

## About

This project is based on [Riven](https://github.com/MingweiSamuel/Riven/)'s code generation.

## Changes

Changes had to be made to translate the rust output to golang (I miss Option<>), mainly with how the clients are organized (I miss impl).

Removed promises, removed `glob-promise`, `glob`, `request`, `request-promise-native`.

Added `node-fetch-commonjs`.

Added a prettier and eslint.

Separated the download and compilation process different files.

## Usage

After installing the dependencies with `npm i`, run `npm run update` and then `npm run compile`.
