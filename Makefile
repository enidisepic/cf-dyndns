# ========================
# Variables and Meta
# ========================

OUTPUT_DIRECTORY ?= out
OUTPUT_FILE      := $(OUTPUT_DIRECTORY)/cf-dyndns
CMD_DIRECTORY    := $(CURDIR)/cmd/cf-dyndns

# ========================
# Building and Running
# ========================

.PHONY: run
run:
	@echo "-> Running cf-dyndns <-"
	go run $(CMD_DIRECTORY)
	@echo

.PHONY: build
build:
	@echo "-> Building cf-dyndns <-"
	go build -o $(OUTPUT_FILE) $(CMD_DIRECTORY)
	@echo
