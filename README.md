# Ripplet

Ripplet is a dynamically typed script language, it internally compile its source code to opcodes, and then a internal opcode interpreter is responsible for executing that opcodes.

Ripplet is aiming to become a efficient and lightweight script language by keeping its
grammar as simple as possible and easy to be embedded in golang application.

## Development

### Prerequisite

- [ANTLR](https://www.antlr.org/) is required, macOS users can follow below command to install it:

	```bash
	brew install antlr
	```

  for user on other platforms please consult [The getting started guide](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)

- go 1.16+

### Project structure 

Project structure is complied with [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
