/*
Copyright © 2023 Vladimir Voitenko vladimirdev635@gmail.com

*/
package main

import (
	"fmt"
	"osiris/cmd"
)

var Version = "0.0";

func main() {
	fmt.Println(Version)
	cmd.Execute()
}
