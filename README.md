# SecurityBaeHash

## Why

Passwords are the most vital part of everyday life. Any individual can
have numerous accounts across various applications, and each of those
accounts has an associated password. Unfortunately as the number of
accounts/passwords grows, people tend to re-use passwords. There are existing
solutions in the form of password managers, like keypassxc or lastpass, but
even those require an account/password. The issue that I have with password
managers is if the master key is compromised, then the attackers have the keys
to the kingdom.

Security Before anything else Hash (SBH) is a ***stateless password
manager***, meaning it doesnt retain any information. It doesnt require
an account, doesnt store any passwords, and it never will. A password
only exists when its retrieved, and the same password can only be
retrieved with the exact same inputs. No keys and no kingdom.

---

## How it Works

SBH is a combination of two things: [caesar
cipher](https://en.wikipedia.org/wiki/Caesar_cipher) and a [hash
function](https://en.wikipedia.org/wiki/Cryptographic_hash_function). It
takes three inputs: plaintext, a number of rotations, and a seed. SBH
then generates a hash based on the result of applying X psuedo-randomly
generated rotations (0 < n < 9223372036854775807) to the plaintext, where
X is the specified number of rotations. The resulting hash is then used
as the password.

---

## Examples

See the [examples](examples) directory.

For a publicly available web version please visit
<https://sbhpass.herokuapp.com>.

---

## Todo
* [X] Add different hashing algorithm options
* [X] Add ability to add symbols and/or capitalize letters
* [ ] Add different cipher options
* [ ] Change letter capitalization
* [ ] Change where symbols are appended based on seed
