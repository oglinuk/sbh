# Golang SecurityBaeHash

## Example 

### Normal

`make` && `./go-sbh`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

No algorithm specified with -a, defaulting to sha256 ...
SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 1.604004ms
```

---

### W/ Specified Hashing Algorithm

`make` && `./go-sbh -a sha512_256`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 6a19b273eb219d0617b5e81aec263b84186aab22764b0d38890eda3868a4ba1f
Elapsed time: 2.096366ms
```

---

### W/ Uppercase Letter(s)

`make` && `./go-sbh -u`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196A7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 3.053907ms
```

---

### W/ Specified Number of Uppercase Letters

`make` && `./go-sbh -u -ut 3`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196A7f528702e5cA85cd0Ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 3.468862ms
```

---

### W/ Symbol(s)

Note: Some combinations (like "!@") cause an error. Need to fix.

`make` && `./go-sbh -s "!"`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2!
Elapsed time: 1.429811ms
```

---

### W/ Web

`make` && `./go-sbh -w`

goto:

`http://localhost:9001/sbh?algorithm=&plaintext=test&nrots=1729&seed=42&uppercase=false&uptimes=0&symbols=`

or

`http://localhost:9001/sbh?algorithm=sha1&plaintext=test&nrots=1729&seed=42&uppercase=true&uptimes=1&symbols=!`

---

### W/ Docker

`make docker`

or

`make docker-web`

---

## Todo
* [X] Dockerize
* [X] Add different hashing algorithm options
* [X] Add ability to add symbols and/or capitalize letters
* [ ] Add different cipher options
