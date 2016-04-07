#!/bin/bash

if [ $# -ne 1 ]
then
  echo "Usage: $0 [tag]"
  exit 1
fi

TAG=$1
git tag | grep $TAG > /dev/null 2>&1
if [ $? -eq 0 ]
then
  echo "Git tag $TAG already exists"
  exit 1
fi

git checkout master
git merge develop --commit

VERSION=( ${TAG//./ } )

sed -i .bak -e "s/Major: [0-9]*,/Major: ${VERSION[0]},/" -e "s/Minor: [0-9]*,/Minor: ${VERSION[1]},/" -e "s/Build: [0-9]*,/Build: ${VERSION[2]},/" predix.go
rm -f predix.go.bak

./build.sh

git add --all
git commit -m "Release $TAG"
git tag $TAG

DATE=`date "+%Y-%m-%d"`
OSX_SHA1=`cat bin/osx/predix | openssl sha1`
LINUX64_SHA1=`cat bin/linux64/predix | openssl sha1`
WIN64_SHA1=`cat bin/win64/predix.exe | openssl sha1`

echo "-------------PLUGIN INFO---------------"
sed -e "s/__TAG/$TAG/" -e "s/__DATE/$DATE/" -e "s/__OSX_SHA1/$OSX_SHA1/" -e "s/__LINUX64_SHA1/$LINUX64_SHA1/" -e "s/__WIN64_SHA1/$WIN64_SHA1/" repo-index.yml
echo "-------------PLUGIN INFO---------------"

echo ""
echo "Release done. Use 'git push --tags', to push the changes to github."
echo "Use the output between 'PLUGIN INFO' to publish to http://plugins.cloudfoundry.org/."
