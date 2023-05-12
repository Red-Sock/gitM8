RSCLI_VERSION=V0.0.21-alpha
rscli-version:
	@echo $(RSCLI_VERSION)
#==============
# migrations
#==============

GOOSE_VERSION=$(shell goose -version)
MIG_DIR="migrations/"
goose-dep:
ifeq ("$(GOOSE_VERSION)", "")
	@echo "installing goose..."
	@go install github.com/pressly/goose/v3/cmd/goose@latest
else
	@echo "goose is installed!"
endif

mig-up:

	@echo "applying migration on postgres"
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql://postgres:password@0.0.0.0:5432/postgres goose up

#==============
