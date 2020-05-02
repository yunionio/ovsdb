package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"

	"yunion.io/x/ovsdb/types"
)

const (
	genSchema = "schema"
	genTypes  = "types"
)

func main() {
	var (
		gen    string
		fname  string
		outDir string
	)
	flag.StringVar(&gen, "gen", genSchema, "what to generate, 'schema' or 'types'")
	flag.StringVar(&fname, "schema", "", "path to ovsdb schema")
	flag.StringVar(&outDir, "outdir", "", "dir for outputing go files")
	flag.Parse()

	switch gen {
	case genSchema:
		if outDir == "" {
			if suf := ".ovsschema"; strings.HasSuffix(fname, suf) {
				outDir = strings.TrimSuffix(fname, suf)
			}
			if outDir == "" {
				glog.Fatalln("please specify a output dir")
			}
		}

		d, gopkg := filepath.Split(outDir)
		gopkg = strings.ReplaceAll(gopkg, "-", "_")
		outDir = filepath.Join(d, gopkg)

		f, err := os.Open(fname)
		if err != nil {
			glog.Fatalf("open %s: %v", fname, err)
		}

		sch, err := types.ParseSchema(f)
		if err != nil {
			glog.Fatalf("parse: %v", err)
		}
		gen := types.NewGenerator().
			Schema(sch).
			OutDir(outDir)
		if err := gen.Gen(); err != nil {
			glog.Fatalf("gen: %v", err)
		}
	case genTypes:
		ag := types.NewAtomicGen().OutDir("./types")
		if err := ag.Gen(); err != nil {
			glog.Fatalf("atomic gen: %v", err)
		}
	}
}
