# ls

Clone of the classic *NIX `ls` command, written in Go. Windows compatible, with untested Linux compatibility (I can't imagine that it wouldn't work but most distributions include `ls` anyway)


## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

You will need to have the Go toolchain installed on your machine:
- [Toolchain download](https://golang.org/dl/)
- [Installation guide](https://golang.org/doc/install)

### Installing

Installing is a simple process assuming `$GOPATH/bin` is included in your system path, if not you can simply copy the built executable to a folder that is in your system path

```
go install github.com/Melephas/ls
```

Then run `ls` in a command line

## Built With

* [fatih/color](https://github.com/fatih/color) - Output colourisation

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Melephas/ls/tags). 

## Authors

* **Sam Miller** - *Initial work* - [Melephas](https://github.com/Melephas)

See also the list of [contributors](https://github.com/Melephas/ls/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Huge thanks to `fatih` for solving a problem that would've taken me months

## Footnotes
* Windows Powershell already has a `ls` command, but this is little more than an alias for the default Windows `dir`
