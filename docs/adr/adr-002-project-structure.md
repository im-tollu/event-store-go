# Pick project structure

## Decision

We will use the following rules to structure the project's source code:

* All main functions should be placed under `${PROJECT_ROOT}/src/cmd/` directory. Each individual
directory below `src/cmd/` should correspond to a single build executable, hence contain a single 
main package and main function.
* Every package exporting something should contain `${PACKAGE_NAME}.go` file that serves as
an entry point for a human reading the code.