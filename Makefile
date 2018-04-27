# git clone git@bitbucket.org:snippets/nakamkaz/8ea7Xj/golang-makefile.git
NAME := `cat appId`
# VERSION := 0.1.0
SRCS := $(shell find . -type f -name '*.go')
#  LDFLAGS := -ldflags="-s -w -X \"main.Version=$(VERSION)\" -extldflags \"-static\""

help: 
	@echo  $(NAME) Makefile has targets
	@echo  make all : go deploy the app
	@echo  make fmt : go fmt 
	@echo  meke help: default, show this help

all: $(SRCS) fmt
	gcloud app deploy --project=$(NAME) --version=1 --quiet	

fmt: $(SRCS)
	@go fmt $(SRCS)


