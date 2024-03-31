# Step by step instructions

An open source project becomes meaningful when people collaborate to improve the code.

Feel free to look at the code, critique and make suggestions. Lets make `hubapp` better!

## Required version

```sh
go version 1.13
```

### Step 1 - Create workspace

*Skip this step if you already have a github go workspace*
Create a github workspace.

<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-workspace.gif" alt="drawing" style="width: 600px;"/>

### Step 2 - Set GOPATH

*Skip this step if you already have a github go workspace*
Export your GOPATH environment variable in your `go` directory.

```sh
export GOPATH=`pwd`
```

<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-gopath.gif" alt="drawing" style="width: 600px;"/>

### Step 3 - Clone repository

Git clone this repository.

```sh
git clone git@github.com:warrensbox/hubapp.git
```

<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-git-clone.gif" alt="drawing" style="width: 600px;"/>

### Step 4 - Get dependencies

Go get all the dependencies.

```sh
go mod download
```

```sh
go get -v -t -d ./...
```

Test the code (optional).

```sh
go vet -tests=false ./...
```

```sh
go test -v ./...
```

<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-go-get.gif" alt="drawing" style="width: 600px;"/>

### Step 5 - Build executable

Create a new branch.

```sh
git checkout -b feature/put-your-branch-name-here
```

Refactor and add new features to the code.
Go build the code.

```sh
go build -o test-hubapp
chmod +x test-hubapp
./test-hubapp --help
```

Test the code and create a new pull request!

<img src="https://s3.us-east-2.amazonaws.com/kepler-images/warrensbox/tfswitch/tfswitch-build.gif" alt="drawing" style="width: 600px;"/>
