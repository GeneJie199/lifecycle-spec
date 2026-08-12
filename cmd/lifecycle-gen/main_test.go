package main

import "testing"

func TestDigestCanonicalizesTextLineEndings(t *testing.T) {
	if digest([]byte("first\r\nsecond\r\n")) != digest([]byte("first\nsecond\n")) {
		t.Fatal("digest differs between CRLF and LF input")
	}
}
