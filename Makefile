# ========================
# Variables and Meta
# ========================

OUTPUT_DIRECTORY ?= out
OUTPUT_FILE_NAME      := cf-dyndns
OUTPUT_PATH      := $(OUTPUT_DIRECTORY)/$(OUTPUT_FILE_NAME)
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
	go build -o $(OUTPUT_PATH) $(CMD_DIRECTORY)
	@echo
