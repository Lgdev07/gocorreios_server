FROM golang:alpine

WORKDIR /gocorreios_server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Install CompileDaemon which is used for hot reload each time a file is changed
RUN go get github.com/githubnemo/CompileDaemon

# The ENTRYPOINT defines the command that will be ran when the container starts up
# The "go build" command here build from the current directory
# We will also execute the binary so that the server starts up. CompileDaemon handles the rest - anytime any .go file changes in the directory
ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ." -command="./gocorreios_server"
