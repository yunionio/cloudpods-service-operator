#!/bin/bash

set -o errexit
set -o pipefail

readlink_mac() {
  cd `dirname $1`
  TARGET_FILE=`basename $1`

  # Iterate down a (possible) chain of symlinks
  while [ -L "$TARGET_FILE" ]
  do
    TARGET_FILE=`readlink $TARGET_FILE`
    cd `dirname $TARGET_FILE`
    TARGET_FILE=`basename $TARGET_FILE`
  done

  # Compute the canonicalized name by finding the physical path
  # for the directory we're in and appending the target file.
  PHYS_DIR=`pwd -P`
  REAL_PATH=$PHYS_DIR/$TARGET_FILE
}

pushd $(cd "$(dirname "$0")"; pwd) > /dev/null
readlink_mac $(basename "$0")
cd "$(dirname "$REAL_PATH")"

genDoc=$GOPATH/bin/gen-crd-api-reference-docs
if [ ! -f "$genDoc" ]; then
  # make tmp dir
  tmpDir=$(mktemp -d)
  docDir=$tmpDir/gen-crd-api-reference-docs
  # install gen-crd-api-reference-docs
  git clone https://github.com/rainzm/gen-crd-api-reference-docs $docDir && cd $docDir
  go install
  cd -
  rm -rf $tmpDir
fi

$genDoc -config ./gen-doc/example-config.json -api-dir ../api/v1 -out-file ../docs/api/docs.md -template-dir ./gen-doc/template

popd > /dev/null
