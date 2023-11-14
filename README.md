# Old man yells at stuff

Use this tool to make Abe Simpson yell at stuff!

This is the perfect toolkit to improve your slack emoji game.

## Repository

Check out the [repository](./repository/) for some pre-baked yelling!
There are over 500 yellings to choose from!

## Browser Version

Visit [oncilla.github.io/old-man-yells-at/](https://oncilla.github.io/old-man-yells-at/) and let Abe yell at stuff from the comfort of your browser!
All the code is executed locally in your browser and the image does not leave your machine.

## CLI Installation

We provide statically built binaries of the CLI on our [GitHub releases](https://github.com/oncilla/old-man-yells-at/releases).
Download the appropriate bundle for your platform from the assets and run it locally.

If you have the Go toolchain installed, you can also simply install the CLI by running the following command:

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

## Contribute

To regenerate the webassembly bundle, run the following:

```txt
GOOS=js GOARCH=wasm go build -o ./docs/yell-at.wasm ./cmd/wasm/
```
