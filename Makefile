empty:=
space:=$(empty) $(empty)
define newline


endef

all: gen

ovsdb_gen:
	go build ./cmd/ovsdb_gen

.PHONY: ovsdb_gen

schemas := \
	ovn-nb.ovsschema \
	ovn-sb.ovsschema \
	vswitch.ovsschema \
	vtep.ovsschema
schemas := $(addprefix types/,$(schemas))

gen_schema = ./ovsdb_gen -gen schema -schema $(1) -outdir schema/$(notdir $(basename $(1)))
gen_schemas = $(foreach schema,$(schemas),$(call gen_schema,$(schema))$(newline))
gen: ovsdb_gen
	$(gen_schemas)
	./ovsdb_gen -gen types
	goimports -local yunion.io/x/ -w schema/
	goimports -local yunion.io/x/ -w types/*_zz_generated.go
	go build ./schema/...

.PHONY: gen

test:
	go test -v ./...

.PHONY: test
