mocks:
	mockery --all --keeptree

unittest:
	clear && go test ./service/... -v -failfast

.PHONY: mocks