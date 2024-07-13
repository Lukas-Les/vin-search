# vin-search

vin-search is a tool to search for Vehicle Identification Numbers (VINs) in files or directories.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
  - [Build](#build)
  - [Run](#run)
- [Output](#output)

## Installation

First, clone the repository:

```sh
git clone https://github.com/Lukas-Les/vin-search.git
cd vin-search
```

## Usage

### Build

To build the vin-search binary, run:

```sh
make build
```

### Run

Use flags to provide the target file or directory, but not both:

    -f : target file
    -d : target directory
    -o : [optional] output file
    -max : [optional] max results (per file)

Example:

```sh
vin-search -d example
```

Run the following command to see all available flags and their explanations:

```sh
cd src && go run . -h
```
## Output

If the output file (-o) is not provided, results are printed to the console.

The output format is:

`<file path> -> <results> <space> <separated>`

Example Output

```
example/outter-text/inner-text/text.json->WAT4C5FWC8S014234
example/outter-text/inner-text/text3.txt->WAT4C5FWC8S014234 VF34C5FWC8S014287 WAT4C5FWC8S014884 WAT4C5FWC8S099234
example/outter-text/text2.txt->VF34C5FWC8S014234 VF34C5FWC8S014287
example/text.txt->
```