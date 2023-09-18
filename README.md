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
write clipboard contents to clipboard
```
$ clipf -w server.go main.go
```
