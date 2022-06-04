# Step 1: Modules caching
FROM golang:1.16-alpine3.15 as modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# Step 2: Builder
FROM golang:1.16-alpine3.15 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN rm -rf go.mod \
	&& go mod init order-app \
    && go mod tidy 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/app ./cmd/app/main.go

# Step 3: Final
FROM scratch
COPY config.yml /config.yml
COPY --from=builder /bin/app /app
CMD ["/app"]