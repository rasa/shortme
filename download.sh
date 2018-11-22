#!/usr/bin/env bash

dl() {
  file="$(basename "$1")"
  wget -O "${file}" "$1"
}

pushd static/css

dl https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/css/bootstrap.min.css
dl https://maxcdn.bootstrapcdn.com/css/ie10-viewport-bug-workaround.css

popd

pushd static/js

dl https://cdn.jsdelivr.net/npm/bootstrap-offcanvas@1.0.0/ie-emulation-modes-warning.js
dl https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js
dl https://cdnjs.cloudflare.com/ajax/libs/jquery.qrcode/1.0/jquery.qrcode.min.js
dl https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js
dl https://cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js
dl https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/js/bootstrap.min.js
dl https://maxcdn.bootstrapcdn.com/js/ie10-viewport-bug-workaround.js

popd
