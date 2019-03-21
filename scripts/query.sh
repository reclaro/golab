#! /bin/sh

curl https://api.github.com/search/code\?q\=TODO+in:file+repo:fnproject/fn | jq .
