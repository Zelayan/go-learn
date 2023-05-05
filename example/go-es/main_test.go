package main

import (
	"github.com/olivere/elastic/v7"
	"testing"
)

func Test_scollSearch(t *testing.T) {
	type args struct {
		client    *elastic.Client
		indexName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				client:    client(),
				indexName: indexName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scrollSearch(tt.args.client, tt.args.indexName)
		})
	}
}

func Test_search(t *testing.T) {
	type args struct {
		client    *elastic.Client
		indexName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				client:    client(),
				indexName: indexName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			search(tt.args.client, tt.args.indexName)
		})
	}
}

func Test_fromSizeSearch(t *testing.T) {
	type args struct {
		client    *elastic.Client
		indexName string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				client:    client(),
				indexName: indexName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fromSizeSearch(tt.args.client, tt.args.indexName)
		})
	}
}
