package main

var _headerTemplate = `
// Code generated by atreus tool genbts. DO NOT EDIT.

NEWLINE
/* 
  Package {{.PkgName}} is a generated cache proxy package.
  It is generated from:
  ARGS
*/
NEWLINE

package {{.PkgName}}

import (
	"context"
	{{if .EnableBatch }}"sync"{{end}}
NEWLINE
	"github.com/mapgoo-lab/atreus/pkg/cache"
	{{if .EnableBatch }}"github.com/mapgoo-lab/atreus/pkg/sync/errgroup"{{end}}
	{{.ImportPackage}}
NEWLINE
	{{if .EnableSingleFlight}}	"golang.org/x/sync/singleflight" {{end}}
)

{{if .UseBTS}}
var _ _bts
{{end }}
{{if .EnableSingleFlight}}
var cacheSingleFlights = [SFCOUNT]*singleflight.Group{SFINIT} 
{{end }}
`
