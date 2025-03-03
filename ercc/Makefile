# Copyright 2019 Intel Corporation
# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

TOP = ..
include $(TOP)/build.mk

GOTAGS += -tags WITH_PDO_CRYPTO

build: ercc

ercc: ercc_dependencies
	# ERCC's binary is created here.
	# The binary is then referenced by the fpc-peer cli, inserted in the ERCC package,
	# and eventually run by the external launcher.
	$(GO) build $(GOTAGS) -o ercc main.go

ercc_dependencies:
	# hard to list explicitly, so just leave empty target,
	# which forces ecc to always be built

test: build
	$(GO) test $(GOTAGS) $(GOTESTFLAGS) ./...

clean: docker-clean
	$(GO) clean
	rm -f ercc

generate:
	$(GO) generate ./...

DOCKER_IMAGE?=fpc/ercc

docker: ercc
	$(DOCKER) build $(DOCKER_BUILD_OPTS) \
		--build-arg FPC_VERSION=$(FPC_VERSION) \
		-t $(DOCKER_IMAGE):$(FPC_VERSION) \
		.

docker-clean:
	$(DOCKER) rmi -f $(DOCKER_IMAGE)
