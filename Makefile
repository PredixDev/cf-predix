# This is what the binary is called
BINARY=cf-predix

# The build information
BUILD_NUMBER?=DEV
VERSION="1.0.0+BUILD_${TRAVIS_BUILD_NUMBER}"
BUILD := $(shell git rev-parse HEAD)
BUILDS_DIR="builds"
BUILD_SUFFIX := $(shell git diff-index --quiet HEAD -- || echo "-dirty")
TAG=${BUILD}${BUILD_SUFFIX}

GO_PACKAGE=github.com/PredixDev

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.VersionString=${VERSION} \
							-X main.CommitSha=${BUILD}${BUILD_SUFFIX}"
.DEFAULT_GOAL: all

build = CGO_ENABLED=0 GOOS=$(1) GOARCH=$(2) go build -a -tags netgo -installsuffix netgo  ${LDFLAGS} -o ${BUILDS_DIR}/${BINARY}_$(1)_$(2)$(3) ${GO_PACKAGE}/${BINARY}
.PHONY: all clean darwin linux
all: clean cover linux darwin windows

#Cleans the project: Deletes the binaries
clean:
	rm -rf ${BUILDS_DIR}

# running ginkgo twice, sadly, the problem is that -cover modifies the source code with the effect
# that if there are errors the output of gingko refers to incorrect line numbers
# tip: if you don't like colors use gingkgo -r -noColor
test:
	@echo ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending

race:
	@echo ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending --race

cover:
	ginkgo -r --randomizeAllSpecs --randomizeSuites --failOnPending -cover
	@echo 'mode: atomic' >_total
	@for f in `find . -name \*.coverprofile`; do tail -n +2 $$f >>_total; done
	@mv _total total.coverprofile
	@COVERAGE=$$(go tool cover -func=total.coverprofile | grep "^total:" | awk '{print $$NF}'); \
	echo "*** Code Coverage is $$COVERAGE% ***"

##### DARWIN (MAC) BUILDS #####
darwin: build/darwin_386

build/darwin_386: $(sources)
	$(call build,darwin,386,)

##### LINUX BUILDS #####
linux: build/linux_386 build/linux_arm build/linux_arm64

build/linux_386: $(sources)
	$(call build,linux,386,)

build/linux_arm: $(sources)
	$(call build,linux,arm,)

build/linux_arm64: $(sources)
	$(call build,linux,arm64,)

windows: build/windows_386

build/windows_386: $(sources)
	$(call build,windows,386,.exe)