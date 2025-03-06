SHELL := /bin/bash

SERVER_DIR := server
CLIENT_DIR := client

SERVER_SRC := $(SERVER_DIR)/server.go
CLIENT_SRC := $(CLIENT_DIR)/client.go

GO := go

.PHONY: run clean

run:
	$(GO) run $(SERVER_SRC) & sleep 1; $(GO) run $(CLIENT_SRC)