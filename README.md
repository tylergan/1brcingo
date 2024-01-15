# 1 Billion Row Challenge
Tried the 1 BRC challenge for fun to explore and better understand Go and how to setup Bazel projects in tandem with Go.

# File and Function Descriptions
## .vscode/extensions.json
- Lists the recommended extensions for the project.

## .vscode/settings.json
- Contains settings for GitHub Copilot and Git authentication.

## go.sum
- Contains the dependency checksums for the Go modules used in the project.

## pkg/iommap.go
- Contains the function [MapFileIntoMem](file:///Users/tylergan/Desktop/VisualStudioCode/Personal/billion_rows/main.go#11%2C25-11%2C25) which maps a file into memory for direct access.

## pkg/processing.go
- Contains functions for parsing and processing data, including [parseLine](file:///Users/tylergan/Desktop/VisualStudioCode/Personal/billion_rows/pkg/processing.go#14%2C6-14%2C6) and `processChunk`.

## main.go
- Contains the `main` function which loads data into memory and processes it.

## WORKSPACE
- Contains Bazel workspace configuration, including external dependencies.

## BUILD
- Defines Bazel build rules for the project.

## pkg/BUILD.bazel
- Contains Bazel build rules for the `pkg` package.

## go.mod
- Defines the Go module and its dependencies.

## pkg/datastructs.go
- Contains data structures used in the `pkg` package.
