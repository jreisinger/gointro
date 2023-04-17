FROM golang:1.20 as build
WORKDIR /shop
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/shop shop.go

FROM gcr.io/distroless/static-debian11
COPY --from=build /bin/shop /
CMD ["/shop"]