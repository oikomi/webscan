//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
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

package chans

const (
	MaxUrlsChan = 100
	MaxRespChan = 50
)

type UrlsChan struct {
	unrefUrls chan string
	refUrls chan string
}

func CreateUrlsChan() *UrlsChan {
	urlsChan := &UrlsChan{
		unrefUrls : make(chan string, MaxUrlsChan),
		refUrls : make(chan string, MaxUrlsChan),
	}

	return urlsChan
}

type RespChan struct {
	Respchans chan string
}

func CreateRespChan() *RespChan {
	respChan := &RespChan{
		Respchans : make(chan string, MaxRespChan),
	}

	return respChan
}