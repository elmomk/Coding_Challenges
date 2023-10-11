package main_test

import (
	"log"
	"testing"

	"github.com/elmomk/Coding_Challenges/1_wc/go/ggwc/cmd"
)

func TestLine(t *testing.T) {
  get := cmd.LinesInFile("./test.txt")
  log.Printf("Lines In file: %d", get)
  // if err != nil {
  //   log.Fatalln(err)
  //   }
  want := 7145 // wc -l test.txt 
  if get != want {
    log.Fatalf("got %d instead of %d", get, want)
  }
}

func TestBytes(t *testing.T) {
  get := cmd.BytesInFile("./test.txt")
  log.Printf("Bytes In file: %d", get)
  var want int64 =  342190 // wc -c test.txt
  if get != want {
    log.Fatalf("got %d instead of %d", get, want)
  }
}
