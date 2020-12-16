FROM golang:latest as BUILD
RUN apt-get clean && \
    apt-get update && \
    apt-get install -y binutils upx
WORKDIR builddir
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o go-gmux-proper-unit-testing-api
# RUN upx --best --ultra-brute go-gmux-proper-unit-testing-api

FROM scratch
COPY --from=BUILD ./go/builddir/go-gmux-proper-unit-testing-api .
ENTRYPOINT ["./go-gmux-proper-unit-testing-api"]