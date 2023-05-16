# Janus

`janus` is a simple and easy-to-use command-line utility for comparing
files and finding differences between them. In ancient Roman religion
and myth, **Janus** is the god of beginnings, transitions, and endings,
hence the utility's name.

## Installation

### From source

First install the dependencies:

- Go 1.20 or above.
- make.
- [scdoc](https://git.sr.ht/~sircmpwn/scdoc).

Then compile and install:

```console
make
sudo make install
```

## Usage

To run `janus`, provide two files to compare in your terminal:

```console
janus <oldFile> <newFile>
```

See _janus(1)_ for more information.

## Contributing

Anyone can help make `janus` better. Check out [the contribution
guidelines](https://git.sr.ht/~jamesponddotco/janus/tree/trunk/item/CONTRIBUTING.md)
for more information.

## Resources

The following resources are available:

- [Support and general discussions](https://lists.sr.ht/~jamesponddotco/janus-discuss).
- [Patches and development related questions](https://lists.sr.ht/~jamesponddotco/janus-devel).
- [Instructions on how to prepare patches](https://git-send-email.io/).
- [Feature requests and bug reports](https://todo.sr.ht/~jamesponddotco/janus).

---

Released under the [EUPL license](LICENSE.md).
