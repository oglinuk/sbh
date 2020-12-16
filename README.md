# SecurityBaeHash


## Why

Passwords are the most vital part of everyday life. Any individual can
have numerous accounts across various applications, and each of those
accounts has an associated password. Unfortunately as the number of
accounts/passwords grows, people tend to become lazy and re-use
passwords. There are existing solutions in the form of password managers,
like keypassxc or lastpass, but even those require an account/password.
The issue that I have with password managers is if the master key is
compromised, then the attackers have the keys to the kingdom.

SecurityBaeHash (SBH) is a stateless password manager, meaning it doesnt
retain any information. It doesnt require an account, doesnt store any
passwords, and it never will. A password only exists when its retrieved,
and the same password can only be retrieved with the exact same inputs.
No keys and no kingdom.

## How it Works

SBH is a combination of two things: [caesar
cipher](https://en.wikipedia.org/wiki/Caesar_cipher) and a [hash
function](https://en.wikipedia.org/wiki/Cryptographic_hash_function). It
takes three inputs: plaintext, a number of rotations, and a seed. SBH
then generates a hash based on the result of applying X psuedo-randomly
generated rotations (0 < n < 9223372036854775807) to the plaintext, where
X is the specified number of rotations. The resulting hash is then used
as the password.

## Examples

### Normal

`make` && `./sbh`

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

`make` && `./sbh -a sha512_256`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 6a19b273eb219d0617b5e81aec263b84186aab22764b0d38890eda3868a4ba1f
Elapsed time: 2.096366ms
```

---

### W/ Uppercase Letter(s)

`make` && `./sbh -u`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196A7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2
Elapsed time: 3.053907ms
```

---

### W/ Specified Number of Uppercase Letters

`make` && `./sbh -u -ut 3`

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

`make` && `./sbh -s "!"`

```
Plaintext: test
Number of Rotations: 1729
Seed: 42

SBH: 196a7f528702e5ca85cd0ac664843cfb4bdd615ce5bc384d60db65ee20a30fb2!
Elapsed time: 1.429811ms
```

---

### W/ Web

`make` && `./sbh -w`

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
