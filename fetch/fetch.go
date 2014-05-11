package fetch

import (
	//"crypto/md5"
	"fmt"
	//"sync"
	"os"
	//"html/template"
	//"io"
	//"io/ioutil"
	//"net/http"
	"net/url"
	"regexp"
	"strings"
	//"../httplib"
	"container/list"
)

type FetchData struct {
	u    string
	//body string
	//hash string
}

type iFetcher interface {
	Fetch(u string) (*FetchData, []string, error)
}

type Fetcher struct {
	entryPoint []string
}

func Fetch(u string) (*FetchData, []string, error) {
	addUrl := make([]string, 0)
	resp, err := Get(u).getResponse()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(u)
	//fmt.Println(resp.Status)
	if resp.Status == HTTP_OK {
		body, err := Get(u).String()
		if err != nil {
			fmt.Println(err)
		}
		bodyStr := string(body)
		
		r, _ := regexp.Compile("<[a|A].*?href=\"([^#?\">]*).*?>")
		//r, _ := regexp.Compile("([<a.*\\/a>]+)")
		for _, ul := range r.FindAllStringSubmatch(bodyStr, -1) {
			//fmt.Println(ul)
			var nu string

			if strings.HasPrefix(ul[1], "http") {
				nu = strings.TrimSpace(ul[1])
				addUrl = append(addUrl, nu)
			} else if strings.HasPrefix(ul[1], "//") {
				nu = "http:" + strings.TrimSpace(ul[1])
				addUrl = append(addUrl, nu)
			} else {
				r, _ := regexp.Compile("http://[^/]+")
				nu = r.FindString(u) + strings.TrimSpace(ul[1])
				addUrl = append(addUrl, nu)
			}
		}
	}
	//h := md5.New()
	//io.WriteString(h, escapeBody)
	//hash := fmt.Sprintf("%x", h.Sum(nil))
	//fmt.Println(addUrl)
	return &FetchData{u}, addUrl, nil
}


//var wg sync.WaitGroup

func crawler(urlSet map[string]bool, urlCache map[string]bool, urls []string) {
	//defer func() { wg.Done() }()
	
	if len(urls) == 0 {
		os.Exit(-1)
	}

	for _, ul := range urls {
		//fmt.Println(urls)
		delete(urlSet, ul)

		if _, seen := urlCache[ul]; seen {
			continue
		}
		urlCache[ul] = true

		fmt.Println(ul) // debug
		_, us, err := Fetch(ul)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(us)

		//wg.Add(1)
		//go crawler(urlSet, urlCache, urls)
	}
}

func Crawl(urls []string) {
	urlSet := make(map[string]bool)
	urlCache := make(map[string]bool)
	//fmt.Println(urls)
	for _, ul := range urls {
		urlSet[ul] = true
	}
	crawler(urlSet, urlCache, urls)
	
	
	
	//wg.Add(1)
	//go crawler(urlSet, urlCache, urls)

	//wg.Wait()
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getDomain(u string) string {
	up, err := url.Parse(u)
    if err != nil {
        panic(err)
    }
	host := up.Host
	fmt.Println(up.Host)
	if strings.HasPrefix(up.Host, "www") {
		host = up.Host[4:]
	}
	
	return host
}

func MyCrawl(urls []string) {
	urlsSet := make([]string, 0)
	urlsSet = append(urlsSet, urls[0])
	urlsVisit := make([]string, 0)
	//urlsVisitFlag := make(map[string]bool)
	urlsGet := make([]string, 0)

	urlsList := list.New()
	urlsList.PushBack(urls[0])
	
	domain := getDomain(urls[0])
	
	
	for urlsList.Len() != 0 {
		fmt.Println(urlsList.Len())
		fmt.Println(urlsVisit)
		e := urlsList.Front()
		if uStr, ok := e.Value.(string); ok {
			if !stringInSlice(uStr, urlsGet) {
				urlsGet = append(urlsGet, uStr)
			}
			_, us, err := Fetch(uStr)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			
			urlsList.Remove(e)
			for _, ul := range us { 
				if !stringInSlice(ul, urlsVisit) {
					if getDomain(ul) == domain {
						urlsVisit = append(urlsVisit, ul)
						urlsList.PushBack(ul)
					}
				}
			}
		}
	}
}
