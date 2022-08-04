<img align="right" height="200" src="https://user-images.githubusercontent.com/5120554/182931905-09cf1084-5158-4a89-a580-10fd44743d57.jpg">

# ProtosEye
Bird's-view of your proto RPCs in JSON.

## Supports
✅ Supports `google.protobuf.*` types.

✅ Supports all standard protobuf types.


## Installation
Here are two ways that you can install this tool.

1. Go install:
```sh
go install github.com/AmirSoleimani/protoseye/cmd/...
```

2. From source code:
```sh
git clone git@github.com:AmirSoleimani/protoseye.git
cd protoseye
go install ./cmd/...
protoc-gen-protoseye version
```

## How to use!
Once you install it, You can easily use it with `protoc`
```sh
find . -name '*.proto' -exec protoc -I=. \
    --protoseye_out=./outputs {} \;
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
