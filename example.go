// An exemple application written with Gofast
//
// Author: Vincent Composieux <vincent.composieux@gmail.com>

package main

import (
    "./lib"
)

func main() {
    g := gofast.Bootstrap()
    g.Handle()
}