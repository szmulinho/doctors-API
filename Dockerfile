FROM golang:1.21.1-alpine AS build

WORKDIR /doctors
COPY . .

RUN apk add --no-cache git
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o doctors

FROM alpine:latest

WORKDIR /doctors
COPY --from=build /doctors/doctors /doctors/doctors

EXPOSE 8095

CMD ["./doctors"]
