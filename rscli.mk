RSCLI_VERSION=V0.0.20-alpha
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

	@echo "applying migration on postgres_db"
	GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql://postgres:pwd@0.0.0.0:5432/postgres goose up

#==============
