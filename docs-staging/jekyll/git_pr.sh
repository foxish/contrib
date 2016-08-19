#!/usr/bin/env bash

# Copyright 2016 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
LOG="/log.txt"

for i in "$@"
do
case $i in
    -p=*|--pull-request=*)
    PR="${i#*=}"
    shift
    ;;
    *)
    # unknown option
    ;;
esac
done

# check if PR not supplied.
if [ "$PR" = "" ]; then
	echo "Empty PR" >> ${LOG}
    exit 1
fi

# check if already serving
curl http://localhost:$PR
if [ $? -eq 0 ] ; then
	echo "PR ${PR} serving" >> ${LOG}
	exit 0
fi

# check if we need to pull
mkdir /pulls/$PR
if [ $? -eq 0 ]; then
    cd /pulls/$PR
	git init
	git fetch https://github.com/kubernetes/kubernetes.github.io refs/pull/$PR/head:$PR
	git checkout $PR
	echo "PR ${PR} pulled" >> ${LOG}
fi

jekyll serve --port $PR &
echo "Attempting to serve PR ${PR}" >> ${LOG}


