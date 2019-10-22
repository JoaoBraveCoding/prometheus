// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Allow Go 1.11 until Go 1.12 is used.
// Go 1.12 is only required to support darwin platforms properly.
// +build go1.11

// Package goversion enforces the go version suported by the tsdb module.
package goversion

const _SoftwareRequiresGOVERSION1_12 = uint8(0)
