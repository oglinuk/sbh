# REST API Implementation

## Compile

`go build main.go handlers.go` || `./init build`

## How to Use

`./main`

send a POST request using `cURL`

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "md5"}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "md5", "uppercase": true}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "md5", "uppercase": true, "uppercasetimes": 3}'
```
