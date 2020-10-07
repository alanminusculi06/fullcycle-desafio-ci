FROM golang:1.15-alpine as builder
 
WORKDIR /src/
COPY . .
RUN go build main.go \
    && chmod +x main

FROM scratch

COPY --from=builder ./src/hello_world ./hello_world

ENTRYPOINT [ "./main" ]