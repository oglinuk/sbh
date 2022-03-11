# REST API Implementation

## How to Use

`go build && ./rest`

send a POST request using `cURL`

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "length": 12}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "sha512", "length": 12}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "sha512", "uppercasetimes": 2, "length": 12}'
```

```BASH
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": 1729, "seed": 42, "algorithm": "sha512", "uppercasetimes": 2, "symbols": "!@#$", "length": 12}'
```
