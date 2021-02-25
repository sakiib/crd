CONTROLLER_GEN=/home/sakib/go/bin/controller-gen

all: generate manifests

.PHONY: manifests generate

# Generate manifests for CRDs
manifests:
	$(CONTROLLER_GEN) crd:trivialVersions=true paths="./..." output:crd:artifacts:config=config/crd/bases

generate:
	go run ./cmd/codegen/main.go