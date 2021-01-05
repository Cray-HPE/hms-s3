#!/usr/bin/env bash

# When it's time to update the version information, use this script as a person

CHANGELOG=../CHANGELOG.md

if [ ! -f $CHANGELOG ]; then
    if [ -f CHANGELOG.md ]; then
        CHANGELOG=CHANGELOG.md
    else
        echo $0: CHANGELOG.md not found in . or .. 1>&2; exit 1
    fi
fi

echo "Detecting the version from the $CHANGELOG..."
VERSION=$(grep -m1 \#\# $CHANGELOG | sed -e "s/\].*$//" -e "s/^.*\[//")
echo "Version = $VERSION"

echo "Add the version to the VERSION file"
echo ${VERSION} > $(dirname $CHANGELOG)/version

# tag it
git commit -am "Version $VERSION"
git tag -a "v$VERSION" -m "Version $VERSION"
git push
git push --tags
