This npm package provides binary golangci-lint-helper which can be used with
[lint-staged](https://github.com/okonet/lint-staged) to lint staged Go files.

[golangci-lint](https://golangci-lint.run/) expects to have package names as arguments, 
but lint-staged provides file names. This package converts file names to package names and
passes them to golangci-lint.
