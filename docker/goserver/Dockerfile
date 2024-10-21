FROM golang:1.22.8-alpine3.20 AS build-stage

# set working directory
WORKDIR /app

RUN apk update \
  && apk add --no-cache \ 
  build-base  

COPY go.mod .
COPY go.sum .
RUN go mod download 

COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/goqlk ./cmd/http/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

# copy binary
COPY --from=build-stage /app/bin/goqlk ./

EXPOSE 8080

CMD ["/goqlk"]