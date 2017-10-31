# Kleister: CLI client

[![Build Status](http://github.dronehippie.de/api/badges/kleister/kleister-cli/status.svg)](http://github.dronehippie.de/kleister/kleister-cli)
[![Go Doc](https://godoc.org/github.com/kleister/kleister-cli?status.svg)](http://godoc.org/github.com/kleister/kleister-cli)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/kleister-cli)](http://goreportcard.com/report/github.com/kleister/kleister-cli)
[![](https://images.microbadger.com/badges/image/kleister/kleister-cli.svg)](http://microbadger.com/images/kleister/kleister-cli "Get your own image badge on microbadger.com")
[![Join the chat at https://gitter.im/kleister/kleister](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kleister/kleister)
[![Stories in Ready](https://badge.waffle.io/kleister/kleister-api.svg?label=ready&title=Ready)](http://waffle.io/kleister/kleister-api)

**This project is under heavy development, it's not in a working state yet!**

Where does this name come from or what does it mean? It's quite simple, it's one german word for paste/glue, I thought it's a good match as it glues together the modpacks for Minecraft.

This project acts as a CLI client implementation to interact with an alternative Kleister API implementation. You can find the sources of the Kleister API at https://github.com/kleister/kleister-api.

The structure of the code base is heavily inspired by Drone, so those credits are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for this awesome project!


## Install

You can download prebuilt binaries from the GitHub releases or from our [download site](http://dl.webhippie.de/kleister-cli). You are a Mac user? Just take a look at our [homebrew formula](https://github.com/kleister/homebrew-kleister). If you are missing an architecture just write us on our nice [Gitter](https://gitter.im/kleister/kleister-api) chat. If you find a security issue please contact thomas@webhippie.de first.


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies and we are not exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get github.com/kleister/kleister-cli/cmd/kleister-cli` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/kleister/kleister-cli
cd $GOPATH/src/github.com/kleister/kleister-cli
make clean build

./kleister-cli -h
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2016 Thomas Boerger <http://www.webhippie.de>
```
