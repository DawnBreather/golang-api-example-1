FROM --platform=amd64 golang:1.18 as build
WORKDIR /workbench
COPY . .
RUN GOARCH=amd64 GOOS=linux go build -o /output/api ./main.go
RUN chmod +x /output/api



FROM --platform=amd64 ubuntu as runtime

COPY --from=build /output/api /usr/bin/api

RUN apt update && apt install ca-certificates -y && rm -rf /var/lib/apt/lists/*

CMD ["/usr/bin/api"]