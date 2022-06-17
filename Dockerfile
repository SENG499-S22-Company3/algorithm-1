FROM golang:alpine AS builder 
RUN apk update && apk add --no-cache git 
WORKDIR /app
COPY go.mod ./

RUN go mod download
# copy source code
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch as runner
COPY --from=builder /app/main .
CMD ["./main"]