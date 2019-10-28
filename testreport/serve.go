package main
import (
	"flag"
	"fmt"
	"net"
	"./github.com/spf13/pflag"
	// "fmt"
	"net/http"
	"os"
	"./testreport/handler"
	// "strings"
	"log"
)

var (
	argPort          = pflag.Int("port", 8080, "The port to listen to for incoming HTTP requests")
	argBindAddress   = pflag.IP("bind-address", net.IPv4(0, 0, 0, 0), "The IP address on which to serve the --port (set to 0.0.0.0 for all interfaces).")
)

func main() {
	var nodes = flag.String("nodelist","","nodelist")
	http.HandleFunc("/getlist", handler.GetMes)
	flag.Parse()
	fmt.Println(*nodes)
	fmt.Println(os.Args)
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Printf("Using HTTP port: %d", *argPort)
	log.Print(http.ListenAndServe(fmt.Sprintf("%s:%d", *argBindAddress, *argPort), nil)) //设置监听的端口
}
