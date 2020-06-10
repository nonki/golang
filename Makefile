mocks: mockgen
	$(MOCKGEN) -package mocks -destination mocks/foo.go github.com/nonki/golang/foo Issuer,Container

# Run go fmt against code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./...

test: fmt vet
	go test ./... -coverprofile cover.out

mockgen:
ifeq (, $(shell which mockgen))
	go get github.com/golang/mock/mockgen
MOCKGEN=$(GOBIN)/mockgen
else
MOCKGEN=$(shell which mockgen)
endif
