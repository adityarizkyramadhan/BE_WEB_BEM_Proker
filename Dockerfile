FROM golang:alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENV SUPABASE_USER=postgres
ENV SUPABASE_PASSWORD=adityarizky1020
ENV SUPABASE_HOST=db.jgjyjvyldoamqndazixl.supabase.co
ENV SUPABASE_PORT=5432
ENV SUPABASE_DB_NAME=postgres
ENV SECRET_KEY=PROJECTBEMPROKER
ENV BYCRIPT_COST=12

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

RUN mkdir -p ./assets/image/product

COPY --from=builder /app/main .

CMD ["./main"]
