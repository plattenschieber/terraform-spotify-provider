#!/usr/bin/bash
set -ex

organization=plattenschieber
plugin=spotify
version=0.1.0
arch=windows_amd64
filename=terraform-provider-${plugin}_v${version}

function build_win64() {
  go build -o bin/${filename}.exe
}

function install() {
  mkdir -p ${APPDATA}/terraform.d/plugins/local/${organization}/${plugin}/${version}/${arch}/
  mv bin/${filename} ${APPDATA}/terraform.d/plugins/local/${organization}/${plugin}/${version}/${arch}/${filename}
}

build_win64
install
