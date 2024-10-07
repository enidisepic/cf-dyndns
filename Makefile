# ========================
# Variables and Meta
# ========================

OUTPUT_DIRECTORY ?= out
OUTPUT_FILE_NAME ?= cf-dyndns
OUTPUT_PATH      := $(OUTPUT_DIRECTORY)/$(OUTPUT_FILE_NAME)
CMD_DIRECTORY    := $(CURDIR)/cmd/cf-dyndns

ifeq ($(GOOS), windows)
ifneq ($(lastword $(subst ., , $(OUTPUT_FILE_NAME))), exe)
	OUTPUT_PATH := $(OUTPUT_PATH).exe
endif
endif

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
<<<<<<< HEAD
=======
	@echo

# ========================
# Linting
# ========================

.PHONY: lint
lint:
	@echo "-> Running revive <-"
	revive -config revive.toml -exclude vendor/... -formatter friendly ./...
>>>>>>> nullishamy
	@echo
