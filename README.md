# Golang SecurityBaeHash

## Example 

### Normal

`make` && `./go-sbh`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
```

### W/ Web

`make` && `./go-sbh -w`

goto http://localhost:9001/sbh?plaintext=test&nrots=1729&seed=42

---

### W/ Docker

`make docker`

or

`make docker-web`

---

## Todo
* [X] Dockerize
* [ ] Create a gRPC microservice
* [ ] Add different hashing options
* [ ] Add ability to add uppercase letters and/or symbols
* [ ] Change from `"math/rand"` to `"crypto/rand"`
