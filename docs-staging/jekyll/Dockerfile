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

FROM ubuntu:14.04

# TODO: switch to standard and pin down versions.
RUN apt-get update && \
	apt-get install -y software-properties-common && \
	apt-add-repository -y ppa:brightbox/ruby-ng && \
	apt-get update && \
	apt-get install -y software-properties-common ruby2.2 ruby2.2-dev build-essential zlib1g-dev git nginx nginx-extras curl lua5.1 liblua5.1-0 liblua5.1-0-dev && \
	gem install github-pages

COPY nginx.conf /nginx.conf
COPY git_pr.sh /git_pr.sh
RUN chmod -c 755 /git_pr.sh
CMD nginx -c /nginx.conf