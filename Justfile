linters-run:
  @gofumpt -l -w .
  @golangci-lint run --fix -j 3 ./...
  @nilaway -include-pkgs="github.com/tompaz3/fungo" ./...
#  golangci-lint run --fix -j --new-from-rev=HEAD~1 3 ./...

go-install:
  @go install go.uber.org/nilaway/cmd/nilaway@latest
  @go get github.com/onsi/ginkgo/v2
  @go install github.com/onsi/ginkgo/v2/ginkgo
  @go install mvdan.cc/gofumpt@latest

go-test:
  @ginkgo -r -p ./...
