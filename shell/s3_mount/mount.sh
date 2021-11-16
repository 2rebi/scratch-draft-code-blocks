#!/bin/sh

s3fs $1 $2 -o use_cache=/tmp -o allow_other -o uid=$(id -u $(whoami)) -o gid=$(id -g $(whoami)) -o multireq_max=5 -o use_path_request_style -o url=https://s3-ap-northeast-2.amazonaws.com