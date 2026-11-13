# Packages and Modules in Go

## Overview

This lesson covers how Go organizes code using packages and modules. Packages help organize reusable code, while modules manage a collection of packages and their dependencies.

## Concepts Covered

- Understanding Go packages
- Understanding Go modules
- `package main` and `func main()`
- Creating custom packages
- Importing packages
- Exported and unexported identifiers
- Uppercase and lowercase naming conventions
- Understanding `go.mod`


## Important Points

### 1. Package

- A package organizes related Go source files.
- Go files in the same directory normally belong to the same package.
- A package can contain functions, structs, interfaces, and variables.

### 2. Module

- A module is a collection of Go packages.
- The `go.mod` file defines the module path and dependencies.
- A project can contain multiple packages inside one module.

### 3. Exported and Unexported Identifiers

| Naming | Accessibility |
|---|---|
| `Add()` | Exported |
| `Subtract()` | Exported |
| `add()` | Unexported |
| `subtract()` | Unexported |

Identifiers beginning with an uppercase letter can be accessed from other packages. Lowercase identifiers are restricted to their package.

### 4. Package main

- `package main` is used for executable programs.
- `func main()` is the entry point of an executable.
- Reusable packages do not need a `main()` function.

## Key Takeaway

**Packages organize reusable code, while modules manage packages and dependencies within a Go project.**