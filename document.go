package main

import (
  "fmt"
  "log"
  "strings"
	//"compress/gzip"
	//"encoding/xml"
	"os"
  "path/filepath"
  //"golang.org/x/net/html"
  //"github.com/PuerkitoBio/goquery"
)

// document represents a Wikipedia abstract dump document.
//type document struct {
//	Title string `xml:"title"`
//	URL   string `xml:"url"`
//	Text  string `xml:"abstract"`
//	ID    int
//}

// document represents a tbc blog post
type document struct {
	Title string
  Text  string
	ID    int
}

func visit(filenames *[]string) filepath.WalkFunc {
  return func(path string, info os.FileInfo, err error) error {
    if err != nil {
      log.Fatal(err)
    }
    if info.IsDir() {
      return nil
    }
    fn := filepath.Base(path)
    first_char := fn[0]
    if first_char != '0' && first_char != '1' {
      return nil
    }
    i := strings.LastIndex(fn, ".")
    if -1 == i {
      return nil
    }
    if i > len(fn) - 4 {
      return nil
    }
    i += 1
    ext := fn[i:i+3]
    if ext != "htm" {
      fmt.Println(ext)
      return nil
    }
    *filenames = append(*filenames, fn)
    return nil
  }
}

func loadDocuments(path string) ([]document, error) {
  
  var filenames []string

  err := filepath.Walk(path, visit(&filenames))
  if err != nil {
    panic(err)
  }
  for _, fn := range filenames {
    fmt.Println(fn)
  }
  
	//dec := xml.NewDecoder(gz)
	dump := struct {
		Documents []document `xml:"doc"`
	}{}
	//if err := dec.Decode(&dump); err != nil {
	//	return nil, err
	//}
	
  docs := dump.Documents
	
  //for i := range docs {
	//	docs[i].ID = i
	//}

  //z := html.NewTokenizer(gz)

	return docs, nil
}
