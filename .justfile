set shell := ["/bin/bash", "-c"]

_help:
    @just --list

# list all available targets
list:
    #!/bin/bash
    for d in cmd/*; do
        echo $(basename $d)
    done

# build a specific target
build target:
    @cd cmd/{{ target }} && go build -o ../../bin/{{ target }}
    @echo "built {{ target }}"

# build all targets
build-all:
    #!/bin/bash
    for d in cmd/*; do
        just build $(basename $d)
    done
