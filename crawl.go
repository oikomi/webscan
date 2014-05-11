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

package main

import (
	"./glog"
	"./fetch"
	//"./conf"
	"fmt"
	"os"
)

func main() {
	glog.Info("start")
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <urls>\n", os.Args[0])
		os.Exit(-1)
	}
	urls := make([]string, 0)
	//for _, url := range os.Args[1:len(os.Args)-1] {
	//	fmt.Println(url)
	//	urls = append(urls, url)
	//}
	urls = append(urls, os.Args[1])
	
	fetch.MyCrawl([]string{"http://www.163.com"})
}
