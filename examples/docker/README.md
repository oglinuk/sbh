# Docker Implementation

## How to Use

`docker build . -t sbh:0.1.0` || `./init build`

### CLI

##### Normal

`./init cli`

##### W/ Specific Hashing Algorithm

`./init cli "-a sha512 -l 12"`

##### W/ Uppercase

`./init cli "-ut 3 -l 12"`

##### W/ Symbol(s)

`./init cli "-s ! -l 12"`

### Web

`./init web`

goto <http://localhost:9001>

or

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
