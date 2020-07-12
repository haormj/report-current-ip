# variable
binaryName=report-current-ip
versionPath=github.com/haormj/version
version=v1.0.0
outputPath=_output
GOARCH=amd64

all: build

build:
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	GOARCH=${GOARCH} \
	go build -ldflags "-X '${versionPath}.Version=${version}' \
	                   -X '${versionPath}.BuildTime=$$buildTime' \
	                   -X '${versionPath}.GoVersion=`go version`' \
	                   -X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/${binaryName};

run: build
	./${outputPath}/${binaryName}

clean:
	rm -rf ./${outputPath}
	rm -rf logs

.PHONY: all build run clean
