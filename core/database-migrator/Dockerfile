FROM  golang:1.12.5-alpine3.9
LABEL maintainer="Nick Pocock <pocockn@hotmail.co.uk>"

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Enable Go Modules
ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context.
COPY ./ ./

RUN CGO_ENABLED=0 go build

ENTRYPOINT ["/app/database-migrator"]