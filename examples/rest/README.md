# REST API Implementation

## Compile

`go build -o sbh` || `./init build`

## How to Use

`./sbh`

send a POST request using `cURL`

```
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": "1729", "seed": "42"}'
```
