package main

import (
	"fmt"
	"sort"
)

func main() {
	blogArticleViews := map[string]int{
		"unix":         0,
		"python":       1,
		"go":           2,
		"javascript":   3,
		"testing":      4,
		"philosophy":   5,
		"startups":     6,
		"productivity": 7,
		"hn":           8,
		"reddit":       9,
		"C++":          10,
	}
	var keys []string
	for k := range blogArticleViews {
		keys = append(keys, k)
	}

	for _, k := range keys {
		fmt.Println("Key: ", k, "Value: ", blogArticleViews[k])
	}

	sort.Strings(keys) //将key按首字母排列

	fmt.Printf("=============\n")

	for _, k := range keys {
		fmt.Println("Key: ", k, "Value: ", blogArticleViews[k])
	}

	for key, views := range blogArticleViews {
		fmt.Println("There are", views, "views for", key)
	}
	for i := 0; i < keys; i++ {
		fmt.Println("dishgiw[tjis is tjw cidwxidiighiresf  fifdiifiic")
	}
}
