#!/bin/sh
set -e

a=$(pwd)
for d in $(find . -name go.mod -execdir pwd \;); do
    cd $d
    if [ -f profile.out ]; then
        rm profile.out
    fi
    go test -count=1 -covermode=atomic -coverprofile=profile.out -coverpkg=./... ./...
done
cd $a

if [ -f coverage-temp.out ]; then
    rm coverage-temp.out
fi
for d in $(find . -name profile.out); do
    cat $d >> coverage-temp.out
    rm $d
done

echo mode: atomic > coverage.out
sed '/mode/d' ./coverage-temp.out >> coverage.out
rm coverage-temp.out
