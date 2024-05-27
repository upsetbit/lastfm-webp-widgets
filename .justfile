set shell := ["/bin/bash", "-c"]

# --------------------------------------------------------------------------------------------------

_help:
    @just --list

_copy-assets target:
    @cp -r assets/ cmd/{{ target }}/assets

# --------------------------------------------------------------------------------------------------

# list all available targets
list:
    #!/bin/bash
    for d in cmd/*; do
        echo $(basename $d)
    done

# build a specific target
build target:
    @just _copy-assets {{ target }}
    @cd cmd/{{ target }} && go build -tags exec_local,save_s3 -o ../../bin/{{ target }}
    @echo "built {{ target }}"

# build all targets
build-all:
    #!/bin/bash
    for d in cmd/*; do
        just build $(basename $d)
    done

run target:
    @just build {{ target }}
    @./bin/{{ target }}
