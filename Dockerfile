FROM alpine/k8s:1.14.9
ENV PATH=$PATH:/root
WORKDIR /root
RUN apk add go
COPY main.go .
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go build main.go
CMD mkdir .kube
COPY config .kube/config
ENTRYPOINT [ "main" ] 
EXPOSE 8080
