#!/bin/sh

# clear existing parsers
rm -rf parsers

# build generate and run
go generate

# build the plugins and save to the ~/.kombustion/plugins/ directory
mkdir -p ~/.kombustion/plugins
cd plugins
for plugin in *; do
    cd $plugin
    go build -buildmode plugin -tags plugin -o $plugin.so && cp -f $plugin.so ~/.kombustion/plugins/$plugin.so
    cd ..
done
cd ..

# build kombustion
go build
cp -f ./kombustion /usr/local/bin/kombustion