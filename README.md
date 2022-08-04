# ProtosEye
Bird's-view of your proto RPCs in JSON.

## Supports
✅ Supports `google.protobuf.*` types.

✅ Supports all standard protobuf types.


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

### Example
E.g. Input:
```protobuf
message GetBirdRequest {
    string id = 1;
}

message GetBirdResponse {
    string id = 1;
    string name = 2;
    int age = 3;
    google.protobuf.Timestamp created_at = 4;
}

service BirdService {
    rpc GetBird(GetBirdRequest) GetBirdResponse;
}
```

Output:
```json
// BirdService_bird.GetBirdRequest.json
{
    "id": "mystring"
}

// BirdService_bird.GetBirdResponse.json
{
    "id": "mystring",
    "name": "name",
    "age": 13,
    "created_at": {
        "nanos": 32,
        "seconds": 64
    }
}
```

## Motivation
It's not easy to design a proto file and have a good insight into the output (It gets worse when you have an enormous payload). This tool helps you generate a JSON view of the RPC's Request and Response.

## TODOs
- [ ] Enhance test coverage.
- [ ] Generate smarter and more specific random values.
- [ ] Config file support for filling data using custom input.