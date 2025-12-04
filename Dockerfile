FROM golang:1.25-alpine AS builder

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY ./pkg ./pkg
COPY ./cmd ./cmd

RUN CGO_ENABLED=0 go build -o ard-jellyfin ./cmd/ard-jellyfin

FROM gcr.io/distroless/static-debian13

COPY --from=builder /work/ard-jellyfin /

ENTRYPOINT [ "/ard-jellyfin" ]
