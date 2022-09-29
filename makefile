export BINARY_NAME=hello-go
all: clean build

clean:
   cd chap7/7.1 && go clean
   cd chap7/bin && rm -f ${BINARY_NAME}
build:
 cd chap7 && go build