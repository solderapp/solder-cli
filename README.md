# Kleister: CLI client

[![Build Status](https://cloud.drone.io/api/badges/kleister/kleister-cli/status.svg)](https://cloud.drone.io/kleister/kleister-cli)
[![Join the Matrix chat at https://matrix.to/#/#kleister:matrix.org](https://img.shields.io/badge/matrix-%23kleister-7bc9a4.svg)](https://matrix.to/#/#kleister:matrix.org)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/9833f3cc84c146a2a13cb8fa5543c11e)](https://www.codacy.com/app/kleister/kleister-cli?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=kleister/kleister-cli&amp;utm_campaign=Badge_Grade)
[![Go Doc](https://godoc.org/github.com/kleister/kleister-cli?status.svg)](http://godoc.org/github.com/kleister/kleister-cli)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/kleister-cli)](http://goreportcard.com/report/github.com/kleister/kleister-cli)
[![](https://images.microbadger.com/badges/image/kleister/kleister-cli.svg)](http://microbadger.com/images/kleister/kleister-cli "Get your own image badge on microbadger.com")

**This project is under heavy development, it's not in a working state yet!**

Within this repository we are building the command-line client to interact with the [Kleister API](https://github.com/kleister/kleister-api) server, for further information take a look at our [documentation](https://kleister.tech).

*Where does this name come from or what does it mean? It's quite simple, it's one german word for paste/glue, I thought it's a good match as it glues together the modpacks for Minecraft.*


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.kleister.tech/cli). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/kleister/homebrew-kleister).


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). This project requires Go >= v1.11.

```bash
git clone https://github.com/kleister/kleister-cli.git
cd kleister-cli

make sync generate build

./bin/kleister-cli -h
```


## Security

If you find a security issue please contact kleister@webhippie.de first.


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2018 Thomas Boerger <thomas@webhippie.de>
```
