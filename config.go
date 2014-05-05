/*
    config.go -
 */
package main

import (
    "fmt"
//    "os"
 //   "path"
//    "path/filepath"
//    "runtime"
//    "bitbucket.org/miranr/artistic/utils"
)

func handleConfigFile(ac *ArtisticCtrl) error {
    
    path := ac.configFile
    fmt.Printf("Config file: %q\n", path)

    return nil
}

