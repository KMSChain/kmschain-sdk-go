#!/usr/bin/env bash
LIBNAME="cryptomagic"
LIBFILE="lib$LIBNAME.a"
SOURCE_DIR="`pwd`"

# Cloning and building C++ library
git clone https://github.com/kmschain/cryptomagic.git "$LIBNAME"
cd cryptomagic
rm -rf build
mkdir -p build && cd build

cmake ..
make -j4
cp "$LIBFILE" "$SOURCE_DIR"
cd "$SOURCE_DIR" && rm -rf "$LIBNAME"