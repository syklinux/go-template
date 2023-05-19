FROM golang:1.20 as build
ARG projName=GoTemplate

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct
#ENV GONOPROXY "your private code base"
#ENV GONOSUMDB "your private code base"
WORKDIR /go/release
ADD . .
#RUN git config --global --add url."https://${username}:${password}@${your private code base}/".insteadOf "${your private code base}"
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go mod tidy && go build -ldflags="-s -w" -installsuffix app -o ${projName} src/main.go

FROM centos:7.2.1511 as prod
ARG projName=GoTemplate

ENV projName ${projName}
RUN mkdir /opt/keys
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build /go/release/config/*.json /opt/
COPY --from=build /go/release/${projName} /

CMD ["/GoTemplate"]
