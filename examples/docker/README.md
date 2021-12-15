# Docker Implementation

## Compiling

`docker build . -t sbh:0.1.0` || `./init build`

## How to Use

### CLI

##### Normal

`./init cli`

##### W/ Specific Hashing Algorithm

`./init cli -a md5`

##### W/ Uppercase

`./init cli -u`

or for multiple uppercase letters

`./init cli -u -ut 3`

##### W/ Symbol(s)

`./init cli -s !`

### Web

`./init web`

goto <http://localhost:9001>

or

```
curl -X POST http://localhost:9001 \
	-H "Content-Type: application/json" \
	-d '{"plaintext": "test", "nrots": "1729", "seed": "42"}'
```
