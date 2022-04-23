#
# Copyright 2021 The Sigstore Authors.
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

FROM golang:1.16-alpine as builder
WORKDIR /build

RUN apk add gcc libc-dev

COPY  ./go.mod ./go.sum ./
RUN go mod download

COPY ./mirroring/*.go ./mirroring/
COPY ./cmd/monitor/main.go ./
RUN GO111MODULE=on GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o monitor .

FROM alpine:latest
WORKDIR ./

COPY --from=builder /build/monitor ./monitor

ENTRYPOINT [ "./monitor" ]
