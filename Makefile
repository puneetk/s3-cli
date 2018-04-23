SRC=*.go
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

s3-cli: $(SRC)
	go build -o $@ $(SRC)

test:
	go test

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"
