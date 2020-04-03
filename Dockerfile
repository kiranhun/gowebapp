# Multi-Stage Docker Build for fa19-go-webapp

# First Stage - Build
# Loading Base Golang Image
FROM golang:1.14.1 as builder
# Set as work directory
WORKDIR /go-webapp
ENV SRC_DIR=/go/src/github.com/gowebapp
ADD . $SRC_DIR
RUN cd $SRC_DIR; go get github.com/gorilla/mux;CGO_ENABLED=0 go build -o goapp


# Second Stage - Run
FROM alpine as final
RUN apk add --update tzdata
COPY --from=builder /go/src/github.com/gowebapp/goapp /go-webapp/.
RUN ls -la /go-webapp
ENTRYPOINT ["/go-webapp/goapp"]


# FROM golang:1.14.1 AS builder
# WORKDIR /go/src/github.com/gowebapp/
# COPY . .
# RUN cd main && go get -d -v
# RUN cd main && CGO_ENABLED=0 GOOS=linux go build -a 

# FROM alpine as final
# WORKDIR /
# RUN apk add --update tzdata
# COPY --from=builder /go/src/github.com/gowebapp/main/main .
# CMD ["./main"]


# Set an ENV var that matches your github repo name
#ENV SRC_DIR=/go/src/github.com/gowebapp/
# Add the source code:
#ADD . $SRC_DIR
# Build it:
#RUN cd $SRC_DIR; go get github.com/gorilla/mux;CGO_ENABLED=0 go build -o goapp

