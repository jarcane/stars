go build

$GOOS = "linux"
$GOARCH = "amd64"

go build -o stars.linux

$GOOS = "darwin"

go build -o stars.osx
