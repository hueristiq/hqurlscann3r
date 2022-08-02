package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/hueristiq/hqurlscann3r/internal/configuration"
	"github.com/hueristiq/hqurlscann3r/pkg/hqurlscann3r"
	"github.com/hueristiq/hqurlscann3r/pkg/params"
	"github.com/logrusorgru/aurora/v3"
)

type options struct {
	delay        int
	concurrency  int
	output       string
	noColor      bool
	URLs         string
	updateParams bool
	verbose      bool
}

var (
	co options
	au aurora.Aurora
	ro configuration.Options
)

func banner() {
	fmt.Fprintln(os.Stderr, aurora.BrightBlue(configuration.BANNER).Bold())
}

func init() {
	// general options
	flag.StringVar(&co.URLs, "iL", "", "")
	flag.IntVar(&co.concurrency, "c", 20, "")
	flag.IntVar(&co.concurrency, "concurrency", 20, "")
	flag.BoolVar(&co.updateParams, "update-params", false, "")
	// http options
	flag.IntVar(&co.delay, "d", 100, "")
	flag.IntVar(&co.delay, "delay", 100, "")
	flag.BoolVar(&ro.FollowRedirects, "follow-redirects", false, "")
	flag.BoolVar(&ro.FollowHostRedirects, "follow-host-redirects", false, "")
	flag.StringVar(&ro.HTTPProxy, "http-proxy ", "", "")
	flag.IntVar(&ro.Timeout, "t", 10, "")
	flag.IntVar(&ro.Timeout, "timeout", 10, "")
	flag.StringVar(&ro.UserAgent, "ua", "", "")
	flag.StringVar(&ro.UserAgent, "user-agent", "", "")
	// output options
	flag.BoolVar(&co.noColor, "nC", false, "")
	flag.BoolVar(&co.noColor, "no-color", false, "")
	flag.StringVar(&co.output, "o", "./hqurlscann3r.json", "")
	flag.StringVar(&co.output, "output", "./hqurlscann3r.json", "")
	flag.BoolVar(&co.verbose, "v", false, "")
	flag.BoolVar(&co.verbose, "verbose", false, "")

	flag.Usage = func() {
		banner()

		h := "USAGE:\n"
		h += "  hqurlscann3r [OPTIONS]\n"

		h += "\nOPTIONS:\n"
		h += "   -c, --concurrency              concurrency level (default: 20)\n"
		h += "   -d, --delay                    delay between requests (default: 100ms)\n"
		h += "       --follow-redirects         follow redirects (default: false)\n"
		h += "       --follow-host-redirects    follow internal redirects i.e, same host redirects (default: false)\n"
		h += "       --http-proxy               HTTP Proxy URL\n"
		h += "  -iL, --input-list               input urls list\n"
		h += "  -nC, --no-color                 no color mode\n"
		h += "   -o, --output                   JSON output file (default: ./hqurlscann3r.json)\n"
		h += "   -t, --timeout                  HTTP request timeout (default: 10s)\n"
		h += "  -ua, --user-agent               HTTP user agent\n"
		h += "       --update-params            update params file\n"
		h += "   -v, --verbose                  verbose mode\n"

		fmt.Fprint(os.Stderr, h)
	}

	flag.Parse()
	ro.Parse()

	au = aurora.NewAurora(!co.noColor)
}

func main() {
	banner()

	if co.updateParams {
		if err := params.UpdateOrDownload(params.File()); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("[", au.BrightBlue("INF"), "] params file updated successfully :)")

		os.Exit(0)
	}

	URLs := make(chan string, co.concurrency)

	go func() {
		defer close(URLs)

		var (
			f   *os.File
			err error
		)

		switch {
		case hasStdin():
			f = os.Stdin
		case co.URLs != "":
			f, err = os.Open(co.URLs)
			if err != nil {
				log.Fatalln(err)
			}
		default:
			log.Fatalln("hqurlscann3r takes input from stdin or file using '-d' flag")
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			URL := scanner.Text()

			if URL != "" {
				URLs <- URL
			}
		}

		if scanner.Err() != nil {
			log.Fatalln(scanner.Err())
		}
	}()

	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	var output hqurlscann3r.Results

	for i := 0; i < co.concurrency; i++ {
		wg.Add(1)

		time.Sleep(time.Duration(co.delay) * time.Millisecond)

		go func() {
			defer wg.Done()

			runner, err := hqurlscann3r.New(&ro)
			if err != nil {
				log.Fatalln(err)
			}

			for URL := range URLs {
				results, err := runner.Process(URL)
				if err != nil {
					// fmt.Println(au.BrightRed(" -"), results.URL, au.BrightRed("...failed!"))

					if co.verbose {
						fmt.Fprintf(os.Stderr, err.Error()+"\n")
					}

					continue
				}

				mutex.Lock()
				// fmt.Println(au.BrightGreen(" +"), results.URL, au.BrightGreen("...done!"))
				// fmt.Println(au.BrightGreen(" +"), results.URL, au.BrightGreen("...done!"))
				x := fmt.Sprintf("[ %s ] %s", results.Category, results.URL)
				fmt.Println(x)
				output = append(output, results)
				mutex.Unlock()
			}
		}()
	}

	wg.Wait()

	if err := output.SaveToJSON(co.output); err != nil {
		log.Fatalln(err)
	}
}

func hasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	isPipedFromChrDev := (stat.Mode() & os.ModeCharDevice) == 0
	isPipedFromFIFO := (stat.Mode() & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}
