# Golang SaltBaeHash

## Example 

### W/ CLI

```
./main -pt test -nr 1729 -s 42

SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 1.540469ms
```

### W/ Web

```
./main -w

http://localhost:9001/sbh?plaintext=test&nrots=1729&seed=42

SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 1.593898ms
```

## Todo
* Create a gRPC microservice
* Add different hashing options
* Dockerize
* Optimize