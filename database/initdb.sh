#!/usr/bin/env bash

# /docker-entrypoint-initdb.d/02_dev_setup
if [ -f /.dockerenv ]; then
    echo "I'm inside docker";
    dirs='/docker-entrypoint-initdb.d/01_migrate'
    for d in $dirs; do
        for f in $d/*; do
            echo "$0: running $f"
            psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -d tournaments --no-password -f "$f"
            # psql postgresql://iamtoolazytotip:iamtoolazytotip@database/tournaments -a -f  $f
        done
    done
else
    echo "I'm living in real world!";
    dirs='./migrate ./dev_setup'
    for d in $dirs; do
        for f in $d/*; do
            echo "$f"
        done
    done
fi
