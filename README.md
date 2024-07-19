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
./vin-search -d example
```

Run the following command to see all available flags and their explanations:

```sh
./vin-search -h
```
## Output

If the output file (-o) is not provided, results are printed to the console.

The output format is:

`<file path>-><results> <space> <separated>`

Example Output

```
example/outter-text/inner-text/text.json->WAT4C5FWC8S014234
example/outter-text/inner-text/text3.txt->WAT4C5FWC8S014234 VF34C5FWC8S014287 WAT4C5FWC8S014884 WAT4C5FWC8S099234
example/outter-text/text2.txt->VF34C5FWC8S014234 VF34C5FWC8S014287
example/text.txt->
```

## Install binary
It possible to install the binary to a directory in your PATH.

1. Generate the binary
2. Move the binary to a directory in your PATH (e.g. /usr/local/bin)

```sh
sudo mv vin-search /usr/local/bin
```

3. Make the binary executable

```sh
sudo chmod +x /usr/local/bin/vin-search
```

If you're on a Mac, you may need to use the following command:

```sh
xattr -d com.apple.quarantine vin-search
```


## Uninstall binary
1. Remove the binary from the directory in your PATH

```sh
sudo rm /usr/local/bin/vin-search
```
