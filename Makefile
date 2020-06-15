SHELL=/bin/bash -o pipefail

export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on

.PHONY: open
open:
	 open ./html/*_api.html

.PHONY: external_gen
external_gen:
	 openapi-generator generate \
	       --git-user-id insolar \
	       --git-repo-id testdoc \
           --input-spec openapi_external.yaml \
           --generator-name go \
           --output external_generated \
           --package-name client \
           --skip-validate-spec

.PHONY: external_html
external_html:
	npx redoc-cli bundle --output html/external_api.html ./openapi_external.yaml

.PHONY: external_gen
internal_gen:
	 openapi-generator generate \
	       --git-user-id insolar \
	       --git-repo-id testdoc \
           --input-spec openapi_internal.yaml \
           --generator-name go \
           --output internal_generated \
           --package-name client \
           --skip-validate-spec

.PHONY: external_html
internal_html:
	npx redoc-cli bundle --output html/internal_api.html ./openapi_internal.yaml

external_server_gen:
	mkdir -p external_server && \
	oapi-codegen -package server -generate types,server ./openapi_external.yaml > external_server/generated.go

internal_server_gen:
	mkdir -p internal_server && \
	oapi-codegen -package server -generate types,server ./openapi_internal.yaml > internal_server/generated.go