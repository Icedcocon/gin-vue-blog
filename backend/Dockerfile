FROM golang:1.20 AS build
ENV WORK_PATH /project/backend
WORKDIR /
COPY . ${WORK_PATH}
RUN cd ${WORK_PATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o backend -mod vendor

FROM alpine:latest
ENV WORK_PATH /project/backend
WORKDIR /backend
COPY --from=build  ${WORK_PATH}/backend .
COPY --from=build  ${WORK_PATH}/config.yaml .
COPY --from=0 ${WORK_PATH}/assets/ip2region.xdb ./assets/ip2region.xdb
COPY --from=0 ${WORK_PATH}/assets/wait-for .
RUN chmod a+x ./wait-for
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' > /etc/timezone

#CMD ["./backend", "-c" ,"config.yaml"]