# Project Title

Coyote Test

## Requirements

* Go 1.8.3

### Prerequisites

Get Sources with Git:

```
git clone https://github.com/g0tzer0/coyote.git
```

Get Sources with Go:

```
go get github.com/g0tzer0/coyote
```

### Installing

Build packages and dependencies:

```
go build -o %GOPATH%\src\github.com\g0tzer0\coyote\bin\web.exe github.com/g0tzer0/coyote/web
```

Run the executable, this will require both the data and certs folder from the sources like this:

```
\bin\web.exe
\cert\server.key
\cert\server.crt
\data\cities.geojson
```

Also make sure that port 8443 is available.

```
cd %GOPATH%\src\github.com\g0tzer0\coyote\bin\
web.exe
```

## Running the tests

Be sure the Web Service is running through the web.exe then select your way to test:

* Non-Verbose tests - Test the REST API without any detail:

```
go test github.com/g0tzer0/coyote/test
```

* Verbose tests - Test the REST API including all methods testing and logging of each request headers, status code and body:

```
go test -v github.com/g0tzer0/coyote/test
```

* You can also Redirect the output feed to log file for easier readibility

```
go test -v github.com/g0tzer0/coyote/test > test.log
```

## Using VSCode

* Download and Install the Go Extension for VSCode by lukehoban

* Open VSCode in the downloaded directory

```
vscode .
```

* Press F5 to run the web service in debugger mode.

## Built With

* [Go](https://golang.org/) - The Go open source programming language
* [VSCode](https://code.visualstudio.com/) - Visual Studio Code - Best IDE ever used
* [Go Extension for VSCode](https://marketplace.visualstudio.com/items?itemName=lukehoban.Go) - Amazing Go Extension for VSCode by lukehoban

## Authors

* **Philippe Matte** - *Initial work* - [g0tzer0](https://github.com/g0tzer0)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Thank you