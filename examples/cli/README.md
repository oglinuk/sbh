# CLI Implementation

## Compile

`go build -o sbh` || `./init build`

## How to Use

### Normal

`./sbh`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
No algorithm specified with -a, defaulting to sha256 ...
SBH: 505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb
Elapsed time: 2.353524ms
```

---

### W/ Specified Hashing Algorithm

`./sbh -a sha1`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
No algorithm specified with -a, defaulting to sha256 ...
SBH: 505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb
Elapsed time: 2.296912ms
```

---

### W/ Uppercase Letter(s)

`./sbh -u`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
No algorithm specified with -a, defaulting to sha256 ...
SBH: 505C3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb
Elapsed time: 2.166241ms
```

---

### W/ Specified Number of Uppercase Letters

`./sbh -u -ut 3`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
No algorithm specified with -a, defaulting to sha256 ...
SBH: 505C3C3fab111175642a3073756472621d0312388deC72ee9268c05f96463ecb
Elapsed time: 2.208992ms
```

---

### W/ Symbol(s)

Note: Some combinations (like "!@") cause an error. Need to fix.

`./sbh -s "!"`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42
No algorithm specified with -a, defaulting to sha256 ...
SBH: 505c3c3fab111175642a3073756472621d0312388dec72ee9268c05f96463ecb!
Elapsed time: 2.232295ms
```
