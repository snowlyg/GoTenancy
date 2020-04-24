# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.

FROM golang:1.14.1
# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/snowlyg/go-tenancy

#build the application
RUN cd /go/src/github.com/snowlyg/go-tenancy go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct && go build
# RUN cd /go/src/github.com/snowlyg/IrisAdminApi/backend/config && cp application.yml.example application.yml

# Run the command by default when the container starts.
ENTRYPOINT /go/src/github.com/snowlyg/go-tenancy/go-tenancy

# Document that the service listens on port 8081
EXPOSE 8085
