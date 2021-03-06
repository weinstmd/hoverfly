dependencies: hoverfly-dependencies hoverfly-functional-test-dependencies hoverctl-dependencies hoverctl-functional-test-dependencies

hoverfly-dependencies:
	cd core && \
	glide --quiet install

hoverctl-dependencies:
	cd hoverctl && \
	glide --quiet install

hoverfly-functional-test-dependencies:
	cd functional-tests/core && \
	glide --quiet install

hoverctl-functional-test-dependencies:
	cd functional-tests/hoverctl && \
	glide --quiet install

hoverfly-test: hoverfly-dependencies
	cd core && \
	go test -v $$(go list ./... | grep -v -E 'vendor')

hoverctl-test: hoverctl-dependencies
	cd hoverctl && \
	go test -v $$(go list ./... | grep -v -E 'vendor')

hoverfly-build: hoverfly-test
	cd core/cmd/hoverfly && \
	go build -ldflags "-X main.hoverflyVersion=$(GIT_TAG_NAME)" -o ../../../target/hoverfly

hoverctl-build: hoverctl-test
	cd hoverctl && \
	go build -ldflags "-X main.hoverctlVersion=$(GIT_TAG_NAME)" -o ../target/hoverctl

hoverfly-functional-test: hoverfly-functional-test-dependencies hoverfly-build
	cp target/hoverfly functional-tests/core/bin/hoverfly
	cd functional-tests/core && \
	go test -v $(go list ./... | grep -v -E 'vendor')

hoverctl-functional-test: hoverctl-functional-test-dependencies hoverctl-build
	cp target/hoverctl functional-tests/hoverctl/bin/hoverctl
	cp target/hoverfly functional-tests/hoverctl/bin/hoverfly
	cd functional-tests/hoverctl && \
	go test -v $(go list ./... | grep -v -E 'vendor')

test: hoverfly-functional-test hoverctl-functional-test

build: test

fmt:
	go fmt $$(go list ./... | grep -v -E 'vendor')
