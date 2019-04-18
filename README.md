# GORMock
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxorcare%2Fgormock.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxorcare%2Fgormock?ref=badge_shield)


The fantastic mock for the fantastic [GORM] library, aims to be developer
friendly.

## How it's work?

The package works on the basis of type substitution when compiling the
application. To use the **gormock** package, you need to replace the 
import `gorm.DB` on imports `linker.DB`, package **linker** implements change of
types that will allow to use **gormock** for testing. But for testing you will
need to use the build tag **gormock** to make the compiler understand what
types to use. In the end, the tests using **gormock** must be run with the
command `go test -tags gormock ./...`.

## How do I use it?

See here [GORMock-examples]

## License

© Vasiliy Vasilyuk, 2018-2019

Released under the [BSD 3-Clause License][LICENSE].

[GORMock]: https://git.io/fhHpT 'The fantastic mock for the fantastic GORM
library, aims to be developer friendly.'
[GORM]: https://git.io/fhHbK 'GORM The fantastic ORM library for Golang, aims
to be developer friendly.'
[LICENSE]: https://git.io/fhHbM 'BSD 3-Clause "New" or "Revised" License'
[GORMock-examples]: https://git.io/fhAxN 'Examples of using gormock and linker'


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fxorcare%2Fgormock.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fxorcare%2Fgormock?ref=badge_large)