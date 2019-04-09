FROM golang:1.11.1
WORKDIR /go/src/github.com/joway/kuafu/
COPY . .
RUN make build

FROM 804775010343.dkr.ecr.cn-north-1.amazonaws.com.cn/golang:1.12-alpine

COPY --from=0 /go/src/github.com/joway/kuafu/kuafu /usr/bin/kuafu

ENTRYPOINT ["kuafu"]
