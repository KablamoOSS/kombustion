# Initialization

!> Kombustion requires Go version **1.10** or above

Initialization is the process of compiling the Kombustion binary and all of the dependant resources, such as the resource parsers and included plugins.

## Downloading Kombustion from source and initializing

```sh
go get github.com/KablamoOSS/kombustion/...
cd $GOPATH/src/github.com/KablamoOSS/kombustion/
./init.sh
```

If you would like the application to be available in your PATH, additionally execute the following:

```sh
go install
```