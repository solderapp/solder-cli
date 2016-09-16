# Kleister: SDK for Golang

[![Build Status](http://github.dronehippie.de/api/badges/kleister/kleister-go/status.svg)](http://github.dronehippie.de/kleister/kleister-go)
[![Coverage Status](http://coverage.dronehippie.de/badges/kleister/kleister-go/coverage.svg)](http://coverage.dronehippie.de/kleister/kleister-go)
[![Go Doc](https://godoc.org/github.com/kleister/kleister-go?status.svg)](http://godoc.org/github.com/kleister/kleister-go)
[![Go Report](http://goreportcard.com/badge/github.com/kleister/kleister-go)](http://goreportcard.com/report/github.com/kleister/kleister-go)
[![Join the chat at https://gitter.im/kleister/kleister](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kleister/kleister)
[![Stories in Ready](https://badge.waffle.io/kleister/kleister-api.svg?label=ready&title=Ready)](http://waffle.io/kleister/kleister-api)

**This project is under heavy development, it's not in a working state yet!**

Where does this name come from or what does it mean? It's quite simple, it's one
german word for paste/glue, I thought it's a good match as it glues together the
modpacks for Minecraft.

This project acts as a client SDK implementation written in Go to interact with
the Kleister API implementation. You can find the sources of the Kleister API at
https://github.com/kleister/kleister-api.

The structure of the code base is heavily inspired by Drone, so those credits
are getting to [bradrydzewski](https://github.com/bradrydzewski), thank you for
this awesome project!


## Install

```
go get -d github.com/kleister/kleister-go
```


## Usage

Import the package:

```go
import (
  "github.com/kleister/kleister-go/kleister"
)
```

Create the client:

```go
const (
  host  = "http://kleister.example.com"
  token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0ZXh0IjoiYWRtaW4iLCJ0eXBlIjoidXNlciJ9.rm4cq4Jupb8BvvDdbwyVwC3rr_WDpdEbCTO0-DCYTWQ"
)

client := kleister.NewClientToken(host, token)
```

Get the current user:

```go
profile, err := client.ProfileGet()
fmt.Println(profile)
```

For a further reference please take a look at Godoc, you can find a link to it
above within the list of badges.


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
