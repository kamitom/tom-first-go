FROM golang:latest

# install glide
RUN go get net/http && go get syreclabs.com/go/faker

# create a working directory
WORKDIR /go/src/app
# add source code
ADD src src

# build main.go
RUN go build src/main.go

EXPOSE 8633/tcp

# run main.go
# CMD ["go", "run", "src/main.go"]
# run the binary
CMD ["./main"]