# lexactivator-go
LexActivator - Go licensing library

Refer to following for documentation:

https://docs.cryptlex.com/node-locked-licenses/using-lexactivator

## To run the sample execute following command:

`go run examples/sample.go`

**NOTE:** In order to set a custom relative path for LexActivator shared libs you can pass *-ldflags* option:

`go run -ldflags="-r ./lexactivator_libs" examples/sample.go`
