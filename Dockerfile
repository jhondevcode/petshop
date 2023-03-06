# We use golang alpine as base image
FROM golang:1.20-alpine3.17 

# We copy the source code into the image
COPY . /app

# We move to the application directory
WORKDIR /app

# We install the dependencies
RUN go mod download

# We compile the application
RUN go build -o petshop .

# We expose the port on which the application listens
ARG SERVER_PORT
EXPOSE ${SERVER_PORT}

# We specify the command that will be executed when the container starts
CMD ["./petshop"]
