package main_test

import (
	"testing"

	"github.com/elmomk/Coding_Challenges/1_wc/go/ggwc/cmd"
)

func TestLine(t *testing.T) {
	get := cmd.LinesInFile("./test.txt")
	t.Logf("Lines In file: %d", get)
	want := 7145 // wc -l test.txt
	if get != want {
		t.Fatalf("got %d instead of %d", get, want)
	}
}

func TestBytes(t *testing.T) {
	get := cmd.BytesInFile("./test.txt")
	t.Logf("Bytes In file: %d", get)
	var want int64 = 342190 // wc -c test.txt
	if get != want {
		t.Fatalf("got %d instead of %d", get, want)
	}
}

func TestWords(t *testing.T) {
	get := cmd.WordsInFile("./test.txt")
	t.Logf("Words in file: %d", get)
	want := 58164 // wc test.txt --words
	if get != want {
		t.Fatalf("got %d instead of %d", get, want)
	}
}

func TestChars(t *testing.T) {
	get := cmd.CharsInFile("./test.txt")
	t.Logf("Chars in file: %d", get)
	want := 339292 // wc test.txt -m
	if get != want {
		t.Fatalf("got %d instead of %d", get, want)
	}
}
