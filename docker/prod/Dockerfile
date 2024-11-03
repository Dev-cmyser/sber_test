# Step 1: Vendor Setup
FROM golang:1.23.2-alpine as builder
COPY . /app
WORKDIR /app

# Step 2: Build
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /bin/app ./cmd/app/main.go

# Step 3: Final Stage
FROM scratch
COPY --from=builder /bin/app /
COPY --from=builder /app/config /config
CMD ["/app"]
