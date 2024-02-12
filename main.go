package main

import "github.com/RND2002/onlineMarketGo/internals/routers"

func main() {
	r := routers.SetUpRouter()

	r.Run()
}
