package main

import (
  "bufio"
  "fmt"
  "log"
  "regexp"
  "strings"
  "strconv"
	"os"
  "path/filepath"
  "github.com/PuerkitoBio/goquery"
)

// document represents a tbc blog post
type document struct {
	Title string
  Text  string
  Url   string
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
      //fmt.Println(ext)
      return nil
    }
    *filenames = append(*filenames, fn)
    return nil
  }
}

func scanurls(path string) ([]string, error) {

  //pattern := regexp.MustCompile("<tr><td align="right">\d{4}</td><td>\d{4}-\d{2}-\d{2}</td><td><a href="(http[^"]*)">([^\<]*)</a>&nbsp;&nbsp;&nbsp;<a href="([^"]*)">^</a>&nbsp;&nbsp;</td><td>[^\<]*</td></tr>")  
  //pattern := regexp.MustCompile("<tr><td align=\"right\">\\d{4}</td><td>\\d{4}-\\d{2}-\\d{2}</td><td><a href=\"(http[^\"]*)\">[^<]*</a>&nbsp;&nbsp;&nbsp;<a href=\"[^\"]*\">^</a>&nbsp;&nbsp;</td><td>[^<]*</td></tr>")
  //pattern := regexp.MustCompile("<tr><td align=\"right\">\\d{4}</td><td>\\d{4}-\\d{2}-\\d{2}</td><td><a href=\"(http[^\"]*)\">[^<]*</a>")
  
  pattern := regexp.MustCompile("<tr><td align=\"right\">(\\d{4}).* href=\"(http[^\"]*thebuildingcoder.typepad.com[^\"]*)\"")  
 
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
 
  defer file.Close()
 
  scanner := bufio.NewScanner(file)
 
  scanner.Split(bufio.ScanLines) 
 
  var urls []string
  
  j := 0
 
  for scanner.Scan() {
    line := scanner.Text()
    matches := pattern.FindSubmatch([]byte(line))
    if( 0 < len(matches) ) {
      //fmt.Println(line, "-->", string(matches[1]))
      id, _ := strconv.Atoi(string(matches[1]))
      j++
      if( j != id ) {
        log.Fatal("Blog post URL number out of sync with index: ", id, " != ", j)
      }
      urls = append(urls, string(matches[2]))
    }
  }
  return urls, nil
}

func loadDocuments(path string) ([]document, error) {
  
  // List the blog post source files
  
  var filenames []string
  err := filepath.Walk(path, visit(&filenames))
  if err != nil {
    panic(err)
  }
  n := len(filenames)
  //fmt.Println(n, "files")
  
  // Load URLs from index.html
  
  urls, err := scanurls(path + "/index.html")
  if err != nil {
    panic(err)
  }
  //for i, url := range urls {
  //  fmt.Println(i, url)
  //}
  m := len(urls)
  if( m != n ) {
    log.Fatal("Expected equal number of blog post docs and urls, but ", n, " != ", m)
  }

  // Retrieve blog post document content
  
  docs := make([]document,0,n)

  for _, fn := range filenames {

    fp := path + "/" + fn

    f, err := os.Open(fp)
    if err != nil {
      log.Fatal(err)
    }
    //defer func() {
    //  if err = f.Close(); err != nil {
    //    log.Fatal(err)
    //  }
    //}()

    r := bufio.NewReader(f)    
    doc, _ := goquery.NewDocumentFromReader(r)
    id, _ := strconv.Atoi(fn[0:4])
    if(!(0<id)){
      log.Fatal("Expected positive blog post number, not ", id)
    }
    title := doc.Find("h3").First().Text()
    i := id - 1
    docs = append(docs, document{ title, doc.Text(), urls[i], i })
    f.Close()
  }
  for i := range docs {
    if( docs[i].ID != i ) {
      fmt.Println(i, docs[i].ID, docs[i].Title)
      log.Fatal("Doc index out of sync with blog post number")
    }
	}
	return docs, nil
}
