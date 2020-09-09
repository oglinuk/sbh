# Golang SecurityBaeHash

## Example 

### Normal

`./go-sbh`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
```

### W/ Web

`./go-sbh -w`

goto http://localhost:9001/sbh?plaintext=test&nrots=1729&seed=42

---

### W/ Docker

`docker build -t go-sbh .`

then

`docker run -it go-sbh`

or

`docker run go-sbh /bin/bash -c "./go-sbh -w"`

---

## Todo
* [X] Dockerize
* [ ] Create a gRPC microservice
* [ ] Add different hashing options
* [ ] Add ability to add uppercase letters and/or symbols
* [ ] Change from `"math/rand"` to `"crypto/rand"`
