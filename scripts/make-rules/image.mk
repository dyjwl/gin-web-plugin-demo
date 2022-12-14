# ==============================================================================
# Makefile helper functions for docker image
#

DOCKER := docker
DOCKER_SUPPORTED_VERSIONS ?= 19|20

REGISTRY_PREFIX ?= zhaosir1993

# Determine image files by looking into hack/docker/*.Dockerfile
IMAGE_FILES=$(wildcard ${ROOT_DIR}/Dockerfile)


.PHONY: image.build.verify
image.build.verify:
ifneq ($(shell $(DOCKER) -v | grep -q -E '\bversion ($(DOCKER_SUPPORTED_VERSIONS))\b' && echo 0 || echo 1), 0)
	$(error unsupported docker version. Please make install one of the following supported version: '$(DOCKER_SUPPORTED_VERSIONS)')
endif
	@echo "===========> Docker version verification passed"

.PHONY: image.build
image.build: image.build.verify go.build.verify $(addprefix image.push., gin-demo)

.PHONY: image.push
image.push: image.build.verify go.build.verify $(addprefix image.push., gin-demo)

.PHONY: image.build.%
image.build.%: go.build.linux_amd64.%
	@echo "===========> Building $* $(VERSION) docker image"
	@cat $(ROOT_DIR)/Dockerfile\
		| sed "s#{{REGISTRY_PREFIX}}#$(REGISTRY_PREFIX)#g" >tmp_$*.Dockerfile
	@$(DOCKER) build --pull -t $(REGISTRY_PREFIX)/$*:$(VERSION) -f tmp_$*.Dockerfile .
	@rm tmp_$*.Dockerfile

.PHONY: image.push.%
image.push.%: image.build.%
	@echo "===========> Pushing $* $(VERSION) image to $(REGISTRY_PREFIX)"
	@$(DOCKER) push $(REGISTRY_PREFIX)/$*:$(VERSION)