FROM public.ecr.aws/docker/library/golang:1.18 as builder

RUN mkdir /app
WORKDIR /app

COPY . .
RUN go mod tidy && go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -o /app/main cmd/server/main.go

# FROM public.ecr.aws/docker/library/bash
FROM scratch
COPY --from=builder /app/main /main
COPY assets assets
COPY web web

ENTRYPOINT [ "/main" ]

