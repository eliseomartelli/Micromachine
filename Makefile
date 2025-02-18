GO := go

PKG := micromachine

TEST_DIR := .

.PHONY: test
test:
	$(GO) test -v $(TEST_DIR)/... --cover

.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  test    - Run tests"
	@echo "  help    - Show this help message"

