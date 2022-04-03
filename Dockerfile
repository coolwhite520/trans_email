FROM golang:1.17-alpine as stage-build
LABEL stage=stage-build
WORKDIR /opt/emailproject
ARG GOPROXY=https://goproxy.io
ARG VERSION=Unknown
ARG TARGETARCH
ENV GOPROXY=$GOPROXY
ENV VERSION=$VERSION
ENV TARGETARCH=$TARGETARCH
ENV GO111MODULE=on
ENV GOOS=linux
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux  GOARCH=amd64  go build

FROM busybox:glibc
COPY --from=stage-build /opt/emailproject/email /opt/email
COPY --from=stage-build /opt/emailproject/template.html /opt/template.html
VOLUME ["/opt/configdata"]
CMD ["/opt/email"]

