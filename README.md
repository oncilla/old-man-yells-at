# Old man yells at stuff!

Use this tool to make Abe Simpson yell at stuff!

This is the perfect toolkit to improve your slack emoji game.

## Installation

```
go install github.com/oncilla/old-man-yells-at/cmd/old-man-yells-at@latest
```

## Usage

```
Enjoy Abe yelling at stuff!

Provide an target image and Abe Simpson will yell at.

By default, the resulting image is created in the current working directory
as 'old-man-yells-at-<target-basename>.png'. If Abe should redirect his yelling,
you have the following options:

  - <filename>.png: Create image at the specified filename.
  - png: Create image at 'old-man-yells-at-<target-basename>.png'.
  - hex: Write image hex-encoded to stdout.
  - b64: Write image b64-encoded to stdout.

Usage:
  old-man-yells-at <target-file> [flags]
  old-man-yells-at [command]

Available Commands:
  completion  Generates shell completion scripts
  help        Help about any command
  version     Show the version information

Flags:
  -h, --help            help for old-man-yells-at
  -o, --output string   [png, b64, hex, <filename>.png] (default "png")

Use "old-man-yells-at [command] --help" for more information about a command.
```

## Examples

<img src="./testdata/old-man-yells-at-bazel.png" width=50 >

<img src="./testdata/old-man-yells-at-vscode.png" width=50 >

## Origin

[The Simpsons](https://youtu.be/tJ-LivK4-78)
