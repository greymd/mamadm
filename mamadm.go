package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/greymd/mamadm/generator"
)

var appVersion = "0.0.0"
var usagePrefix = "mamadm: Mama-katsu direct message generator.\n\n"
var copyright = `
Copyright (c) 2022 Yamada, Yasuhiro
Released under the MIT License.
https://github.com/greymd/mamadm
`

var help = flag.Bool("h", false, "ヘルプを表示");
var version = flag.Bool("V", false, "バージョンを表示");
var taste = flag.Int("t", 0, "0 = 文体をランダムに設定（デフォルト）, 1 = 風情ある文体にする, 2 = 通常の文体にする");

func main() {
  flag.Parse()

  if *help {
    fmt.Fprintf(os.Stderr, usagePrefix)
    flag.Usage()
    fmt.Fprintf(os.Stderr, copyright)
    os.Exit(0)
  }

  if *version {
    fmt.Fprintf(os.Stderr, "mamadm %s\n", appVersion)
    os.Exit(0)
  }

  result, err := generator.Generate(*taste)

  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  fmt.Printf("%s\n", result)
}

