// TODO: Add beautiful header
package main

const PORT int = 6379

func main() {
	server := &ReddishServer{}
	server.Start(PORT)
}
