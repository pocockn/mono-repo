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

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN ls -l

RUN cd core/database-migrator && CGO_ENABLED=0 go build -mod=vendor -o database-migrator

ENTRYPOINT ["/app/database-migrator"]
