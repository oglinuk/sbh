# Golang SecurityBaeHash

## Example 

### Normal

`make` && `./go-sbh`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
```

### W/ Uppercase Letter(s)

`make` && `./go-sbh -u`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
```

### W/ Specified Number of Uppercase Letters

`make` && `./go-sbh -u -ut 3`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
```

### W/ Symbol(s)

Note: Some combinations (like "!@") cause an error. Need to fix.

`make` && `./go-sbh -s "!"`

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
