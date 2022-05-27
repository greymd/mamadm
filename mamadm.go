package main

import (
	"fmt"
	"log"
	"os"

	"github.com/greymd/mamadm/generator"
)

var appVersion = `Mama-katsu direct message generator version 0.0.0
Copyright (c) 2022 Yamada, Yasuhiro
Released under the MIT License.
https://github.com/greymd/mamadm`

var usage = `Usage:
  mamadm [options]

Options:
  -h, --help      ヘルプを表示.
  -V, --version   バージョンを表示.`

func main() {
        result, err := generator.Start()
        if err != nil {
                log.Fatal(err)
                os.Exit(1)
        }
        fmt.Printf("%s\n", result)
}

