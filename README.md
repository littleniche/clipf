## Copy file contents to clipboard

## Installation

### Go
```
$ go install github.com/littleniche/clipf@latest
```
### Manual
```
$ make install
```

### Usage

```
$ clipf main.go
```
copy from multiple files
```
$ clipf /etc/fstab main.go
```
copy from stdin
```
$ cat Makefile | clipf
```
write clipboard contents to files
```
$ clipf -w server.go main.go
```
