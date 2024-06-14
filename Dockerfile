FROM golang:1.22.4-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash make gcc gettext musl-dev

# dependencies
COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download && go mod verify

# build
COPY app ./
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/app /

ARG USER_UID=1001
ARG GROUP_GID=1001
ARG UGNAME=www-data

RUN adduser --system --disabled-password --home /home/${UGNAME} \
    --uid ${USER_UID} --ingroup ${UGNAME} ${UGNAME}

USER ${UGNAME}

CMD ["/app"]
