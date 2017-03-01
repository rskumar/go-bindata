// Copyright 2017 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a Modified
// BSD License that can be found in the LICENSE file.

package bindata

import (
	"path/filepath"
	"text/template"
)

func init() {
	template.Must(baseTemplate.New("header").Funcs(template.FuncMap{
		"toslash": filepath.ToSlash,
	}).Parse(`{{- /* This makes e.g. Github ignore diffs in generated files. */ -}}
// Code generated by go-bindata.
{{if $.Dev -}}
	//  debug: dev
{{else if $.Debug -}}
	//  debug: true
{{end -}}
{{- if $.MemCopy -}}
	//  memcopy: true
{{end -}}
{{- if $.Compress -}}
	//  compress: true
{{end -}}
{{- if and $.Compress $.DecompressOnce -}}
	//  decompress: once
{{end -}}
{{- if $.Metadata -}}
	//  metadata: true
{{end -}}
{{- if gt $.Mode 0 -}}
	//  mode: {{printf "%04o" $.Mode}}
{{end -}}
{{- if gt $.ModTime 0 -}}
	//  modtime: {{$.ModTime}}
{{end -}}
{{- if $.AssetDir -}}
	//  asset-dir: true
{{end -}}
{{- if $.Restore -}}
	//  restore: true
{{end -}}
{{- if ne $.HashFormat 0 -}}
	//  hash-format: {{$.HashFormat}}
{{end -}}
{{- if and (ne $.HashLength 0) (ne $.HashLength 16) -}}
	//  hash-length: {{$.HashLength}}
{{end -}}
{{- if ne $.HashEncoding 0 -}}
	//  hash-encoding: {{$.HashEncoding}}
{{end -}}
{{- if $.HashKey -}}
	//  hash-key: <omitted>
{{end -}}
// sources:
{{range .Assets -}}
	//  {{toslash .Path}}
{{end -}}
// DO NOT EDIT!

{{if $.Tags -}}
	// +build {{$.Tags}}

{{end -}}

package {{$.Package}}`))
}
