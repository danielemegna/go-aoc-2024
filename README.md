# Go Advent Of Code 2024

Solutions for [2024 Advent Of Code challenges](https://adventofcode.com/2024) in Golang

Challenge solutions are in tests as assertions:

```go
func TestSolveFirstPartWithFile(t *testing.T) {
	var fileContent = readFileContent()
	var actual = TotalDistanceBetweenListsElements(fileContent)
	assert.Equal(t, 1530215, actual)
}
```

> following AOC authors guidelines, input.txt files are not included in the repo
----

### Docker-only dev setup

Use temporary container with current directory volume:

```
$ docker run --rm -it -v $PWD:/app -w /app golang:1.24.1 bash
# go install github.com/mfridman/tparse@latest
```

and run tests with
```
# go test ./... -json | tparse
# go test ./day01 -json | tparse --all
# go test ./day02 -run TestRunSingleTest -json | tparse --all
```

### VSCode Dev Container setup

VSCode can use devcontainers to be configured with the proper extensions without language specific utilities locally installed.

In order to proceed, install the [related extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) (`ms-vscode-remote.remote-containers`) and reopen vscode in Dev Container mode.

> the first run can take some minutes to install and setup properly container and vscode extensions

If you want to open an additional bash session in the vscode container:

```
$ docker exec -it -w /workspaces/$(basename $PWD) <container-name> bash
# go test .....
```
