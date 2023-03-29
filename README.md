# BSIUII Go Modules



BSIUII is a bunch of method written in [Go](https://go.dev/) to facilitate and speed up 
the software development process at [Badan Sistem Informasi](https://bsi.uii.ac.id) - [Universitas Islam Indonesia](https://uii.ac.id)

**The key features of module are:**

- Customize error types
- Logging function by implementing [Logrus](https://github.com/sirupsen/logrus)


## Getting started

### Prerequisites

- **[Go](https://go.dev/)**: any one of the **three latest major** [releases](https://go.dev/doc/devel/release) (we test it with these).

### Getting BSIUII Module

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/arisatriop/bsiuii"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `bsiuii` package:

```sh
$ go get -u github.com/arisatriop/bsiuii
```


## Documentation

See [API documentation and descriptions](https://godoc.org/github.com/arisatriop/bsiuii) for package.


## Contributing

BSIUII is the work of hundreds of contributors. We appreciate your help!

Please see [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and the contribution workflow.