PKG_NAME=dome9
TEST?=$$(go list ./${PKG_NAME} | grep -v -e 'vendor')
GOFMT_FILES?=$$(find . -name '*.go' | grep -v 'vendor')
WEBSITE_REPO=github.com/hashicorp/terraform-website


default: build

build: fmtcheck
	go install

tools:
	@echo "==> installing required tooling..."
	@sh "$(CURDIR)/scripts/gogetcookie.sh"
	GO111MODULE=off go get -u github.com/client9/misspell/cmd/misspell
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=off go get -u github.com/bflad/tfproviderlint/cmd/tfproviderlint

errcheck:
	@sh -c "'$(CURDIR)/scripts/errcheck.sh'"

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

lint:
	@echo "==> Checking source code against linters..."
	golangci-lint run  --timeout=5m ./$(PKG_NAME)/...

tflint:
	@echo "==> Checking source code against terraform provider linters..."
	@tfproviderlint \
		-c 1 \
		-AT001 \
		-AT002 \
		-S001 \
		-S002 \
		-S003 \
		-S004 \
		-S005 \
		-S007 \
		-S008 \
		-S009 \
		-S010 \
		-S011 \
		-S012 \
		-S013 \
		-S014 \
		-S015 \
		-S016 \
		-S017 \
		-S019 \
		./$(PKG_NAME)

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

test: fmtcheck
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test -v $(TESTARGS) -timeout=30s -parallel=4

testacc: fmtcheck
	TF_ACC=1 go test -v $(TEST) -timeout 240m

test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-lint:
	@echo "==> Checking website against linters..."
	@misspell -error -source=text website/

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

.PHONY: build test testacc vet fmt fmtcheck errcheck lint tools test-compile website website-lint website-test
