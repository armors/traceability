FROM golang:1.14.4 as build
RUN mkdir -p /app/building
WORKDIR /app/building
ADD . /app/building
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io
RUN make build

FROM alpine:3.12.0
# Copy from docker build
COPY --from=build /app/building/target/admin /app/bin/
COPY --from=build /app/building/target/config.json /app/
# Copy from local build
#ADD
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /app/
CMD /app/bin/admin 2>&1 > /app/admin.log