package main

import (
	"io"
	"testing"
)

func TestByteCounter_Write(t *testing.T) {

	tests := []struct {
		name  string
		str   string
		wantN int
	}{
		{"first", "", 0},
		{"second", "hello", 5},
		{"third", "bro", 3},
		{"fourth", "sis", 3},
	}
	var b LineCounter
	w, _ := countingWriter(&b)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := w.Write([]byte(tt.str))

			if err != nil {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantN)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestCountingWriter_Write(t *testing.T) {
	type fields struct {
		writer io.Writer
		count  int64
	}
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n1 := &CountingWriter{
				writer: tt.fields.writer,
				count:  tt.fields.count,
			}
			gotN, err := n1.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Write() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
