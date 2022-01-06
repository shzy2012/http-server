package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

//localIP 输出local ip
func localIP() []net.IP {
	ips := make([]net.IP, 0)
	if ifaces, err := net.Interfaces(); err == nil {
		// handle err
		for _, i := range ifaces {
			addrs, _ := i.Addrs()
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}

				if ip.To4() != nil && !strings.Contains(ip.String(), "169.254") {
					ips = append(ips, ip)
				}
			}
		}
	}

	return ips
}

//loggingHandler 输出log
func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(*r.URL)
		start := time.Now()
		h.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr
		ua := r.UserAgent()
		log.Printf(
			"%s\t%s\t%s\t%s\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			ua,
			time.Since(start),
		)
	})
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func portMustBeAvailable(port int) {
	p := strconv.Itoa(port)
	ln, err := net.Listen("tcp", ":"+p)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't listen on port %q. Please use other port\n", p)
		os.Exit(1)
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't stop listening on port %q: %s\n", p, err)
		os.Exit(1)
	}

	// fmt.Printf("TCP Port %q is available\n", p)
}

func main() {
	var port int
	var open bool
	flag.IntVar(&port, "p", 8080, "http 端口号  -p=8080")
	flag.BoolVar(&open, "o", false, "Open browser automatically")
	flag.Parse()

	portMustBeAvailable(port)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	_, cancel := context.WithCancel(context.Background())
	go func() {

		ips := localIP()
		ipString := ""
		for _, ip := range ips {
			ipString = ipString + fmt.Sprintf("        http://%s:%v\n", ip, port)
		}
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Starting up http-server, serving dir", path)
		fmt.Println("Available on:")
		fmt.Printf("%s", ipString)
		fmt.Println("Hit CTRL-C to stop the server")

		fs := http.FileServer(http.Dir(path))
		http.Handle("/", loggingHandler(fs))

		if open && len(ips) > 0 {
			openbrowser(fmt.Sprintf("http://%s:%v", ips[1].String(), port))
		}

		log.Println(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	}()

	go func() {
		<-sigs
		cancel()
		done <- true
	}()

	<-done
	log.Println("http service stop.")
}
