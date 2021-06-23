package main

import "golang-api/routers"

func main() {
	r := routers.StartRouters()

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
