# Web Implementation

## Compiling

`go build -o sbh` || `./init build`

## How to Use

### Web UI

`./sbh`

goto <https://localhost:9001>

### REST API

`./sbh`

then

`http://localhost:9001/sbh?algorithm=&plaintext=test&nrots=1729&seed=42&uppercase=false&uptimes=0&symbols=`

or

`http://localhost:9001/sbh?algorithm=sha1&plaintext=test&nrots=1729&seed=42&uppercase=true&uptimes=1&symbols=!`
