FROM golang:1.21.6-alpine3.19

# installing git & make
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    apk add --no-cache make


# setup folders
RUN mkdir /app
WORKDIR /app

# copy the source from the current directory to the working directory inside the container
COPY . .

# get dependencies
RUN go get -d -v ./...

# install packages
RUN go install -v ./...

# build app
RUN go build -o /server ./cmd/server/.

EXPOSE 8080

# Run the executable
CMD [ "/server" ]