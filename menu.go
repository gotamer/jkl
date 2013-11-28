// menu.go
package main

import (
	"strings"
)

var menus = make(map[int]menu)

type menu struct {
	Url  string
	Name string
	Cat  []string
}

// m "Linux/Code/Go/GoTamer/sbs.html"
func menuAdd(url string) {
	link := strings.Split(url, "/")
	n := len(link)
	if n == 0 {
		return
	}
	m := menu{}
	m.Name = strings.TrimSuffix(link[n-1], ".html")
	m.Url = url
	if n > 1 {
		for i := 0; i < n-1; i++ {
			m.Cat = append(m.Cat, link[i])
		}
	}
	id := len(menus) + 1
	menus[id] = m
}

type TmplMenu struct {
	Name, Url string
}

func mainMenu() (m []TmplMenu) {
	mt := make(map[string]string)
	for _, v := range menus {
		if len(v.Cat) == 0 {
			if v.Name == "index" {
				v.Name = "Lobby"
				v.Url = ""
			}
			mt[v.Name] = v.Url
		} else {
			mt[v.Cat[0]] = v.Cat[0]
		}
	}
	var mm []TmplMenu
	for k, v := range mt {
		mm = append(mm, TmplMenu{k, "/" + v})
	}

	// Reverse
	for i := len(mm) - 1; i >= 0; i-- {
		m = append(m, mm[i])
	}
	return
}

func subMenu(url string) (m []TmplMenu) {
	var cats []string

	link := strings.Split(url, "/")
	n := len(link)
	if n > 1 {
		cats = link[0 : n-1]
		for _, mv := range menus {
			if len(mv.Cat) != 0 {
				if ok := compareSlice(mv.Cat, cats); ok {
					m = append(m, TmplMenu{mv.Name, "/" + mv.Url})
				}
			}
		}
	}
	return
}

func compareSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
