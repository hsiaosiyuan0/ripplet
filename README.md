# Ripplet

Ripplet is a dynamically typed script language, it internally compiles its source code to opcodes, and then a internal opcode interpreter is responsible for executing that opcodes.

Ripplet is aiming to become an efficient and lightweight script language by keeping its grammar as simple as possible and easy to be embedded in golang application.

## Preview

Getting the preview artifacts on [Release Page](https://github.com/hsiaosiyuan0/ripplet/releases) or follow below commands on MacOS

```bash
$ brew tap hsiaosiyuan0/ripplet
$ brew install ripplet
$ ripplet -f your_script.rip 

$ ripplet -h # print usage
```

## Development

Ripplet is currently in a prototyping stage, so it uses [ANTLR](https://www.antlr.org/) to generate its parser, however a handmade parser will be introduced once it enters a relatively stable stage.

### Prerequisite

- [ANTLR](https://www.antlr.org/) is required, macOS users can follow below command to install it:

	```bash
	$ brew install antlr
	```

  for user on other platforms please consult [The getting started guide](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)

- go 1.16+

### Build & Run

```bash
$ go get -v ./...
$ make
./ripplet -f your_script.rip
```

### Project structure 

Project structure is complied with [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
