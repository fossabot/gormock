# [GORMock] / Linker

The linker is required to connect the required types depending on the type of
application assembly.

## How it work?

During a normal build of application using `go build`, the linker appends the
standard package [GORM] but if you run tests with the addition of a flag
`-tags gormock` linker replace [GORM] on [GORMock] in the tests will have the
opportunity to use a mock instead original types of [GORM].

```Shell
# Using for run tests
go test -tags gormock
```

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
