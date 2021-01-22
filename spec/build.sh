#!/bin/sh

HERE=$(realpath $(dirname $0))
SED="sed"
if [ "$(uname)" = "Darwin" ]; then
    if ! which gsed >/dev/null; then
        echo "'gsed' not found" && exit 1
    fi
    SED="gsed"
fi

npx @apidevtools/swagger-cli bundle -r $HERE/openapi.yml -o $HERE/out/openapi.yml -t yaml
$SED -i -E 's/#components/#\/components/g' $HERE/out/openapi.yml
