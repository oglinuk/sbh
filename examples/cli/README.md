# CLI Implementation

## How to Use

### Normal

`go build` && `./sbh`

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

`go build` && `./sbh -a sha512_256`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 6a19b273eb219d0617b5e81aec263b84186aab22764b0d38890eda3868a4ba1f
Elapsed time: 2.096366ms
```

---

### W/ Uppercase Letter(s)

`go build` && `./sbh -u`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196A7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 3.053907ms
```

---

### W/ Specified Number of Uppercase Letters

`go build` && `./sbh -u -ut 3`

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

`go build` && `./sbh -s "!"`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2!
Elapsed time: 1.429811ms
```