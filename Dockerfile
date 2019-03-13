FROM golang:1.8.5-jessie
# create a working directory
WORKDIR /go/src/playlist-generator
# add source code
COPY . /go/src/playlist-generator/
RUN go get -u github.com/gorilla/mux
# run main.go
CMD ["go", "run", "/go/src/playlist-generator/application.go"]