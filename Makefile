TEST_OUTPUT=test-output.log

.PHONY: test
test: gotestanalyse
	gotestsum \
		--post-run-command ./gotestanalyse \
		--jsonfile $(TEST_OUTPUT) \
		--rerun-fails \
		--packages . \
		-- \
		-count=1

gotestanalyse: ./gotestanalyse.go
	go build $^
