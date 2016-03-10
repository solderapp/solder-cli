# Solder: CLI client

[![Build Status](http://github.dronehippie.de/api/badges/solderapp/solder-cli/status.svg)](http://github.dronehippie.de/solderapp/solder-cli)
[![Coverage Status](https://aircover.co/badges/solderapp/solder-cli/coverage.svg)](https://aircover.co/solderapp/solder-cli)
[![Go Doc](https://godoc.org/github.com/solderapp/solder-cli?status.svg)](http://godoc.org/github.com/solderapp/solder-cli)
[![Go Report](http://goreportcard.com/badge/solderapp/solder-cli)](http://goreportcard.com/report/solderapp/solder-cli)
[![Join the chat at https://gitter.im/solderapp/solder](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/solderapp/solder)
![Release Status](https://img.shields.io/badge/status-beta-yellow.svg?style=flat)

This project acts as a CLI client implementation to interact with an
alternative Solder API implementation. You can find the sources of the Solder
API at https://github.com/solderapp/solder.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

Make sure you have a working Go environment, for further reference or a guide
take a look at the [install instructions](http://golang.org/doc/install.html).
As this project relies on vendoring of the dependencies and we are not
exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go
version `>= 1.6`

```bash
go get -d github.com/solderapp/solder-cli
cd $GOPATH/src/github.com/solderapp/solder-cli
make deps build

bin/drone-cli -h
```

Later on we will also provide a download of prebuilt binaries for various
platforms, but this will start if we get to an somehow working state or if we
are more or less on feature parity with the upstream project.


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
