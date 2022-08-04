# ProtosEye
Bird's-view of your proto RPCs.

## Motivation
It's not easy to design a proto file and have a good insight into the output (It gets worse when you have a big payload). 
This tool helps you to generate JSON output of your RPC's Request/Response with sample values by considering the relation between messages.

### Example
E.g. Input:
```protobuf
// TODO
```

Output:
```json
// TODO
```

## Installation
Here are two ways that you can install this tool.

1. Go install:
```sh
go install https://github.com/AmirSoleimani/protoseye@v1.0.0
```

2. From source code:
```sh
git clone git@github.com:AmirSoleimani/protoseye.git
go install ./cmd/...
protoc-gen-protoseye version
```

## Supports
^ TODO type table

## TODOs
- [ ] Enhance test coverage.
- [ ] More creative and intelligent while generating random values.