FROM golang:1.10-alpine AS build-stage

RUN apk add --update make git
RUN mkdir -p /go/src/github.com/gemsorg/eligibility
WORKDIR /go/src/github.com/gemsorg/eligibility

COPY . /go/src/github.com/gemsorg/eligibility

ARG GIT_COMMIT
ARG VERSION
ARG BUILD_DATE

RUN make build-service

# Final Stage
FROM alpine

RUN apk --update add ca-certificates
RUN mkdir /app
WORKDIR /app

COPY --from=build-stage  /go/src/github.com/gemsorg/eligibility/bin/eligibility .

EXPOSE 3000

CMD ["./eligibility"]