# sigurlscann3r

[![release](https://img.shields.io/github/release/signedsecurity/sigurlscann3r?style=flat&color=0040ff)](https://github.com/signedsecurity/sigurlscann3r/releases) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/signedsecurity/sigurlscann3r.svg?style=flat&color=0040ff)](https://github.com/signedsecurity/sigurlscann3r/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/signedsecurity/sigurlscann3r.svg?style=flat&color=0040ff)](https://github.com/signedsecurity/sigurlscann3r/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/signedsecurity/sigurlscann3r/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@signedsecurity-0040ff.svg)](https://twitter.com/signedsecurity)

sigurlscann3r a web application attack surface mapping tool, it does ...:

* Categorize URLs

	<details>
	<summary>URLs' categories</summary>

	```
	> endpoint
	> js {js}
	> style {css}
	> data {json|xml|csv}
	> archive {zip|tar|tar.gz}
	> doc {pdf|xlsx|doc|docx|txt}
	> media {jpg|jpeg|png|ico|svg|gif|webp|mp3|mp4|woff|woff2|ttf|eot|tif|tiff}
	```

	</details>

* Next, probe HTTP requests to the URLs for `status_code`, `content_type`, e.t.c
* Next, for every URL of category `endpoint` with a query:
	* Probe for commonly vulnerable parameters (inspired by [Somdev Sangwan](https://github.com/s0md3v)'s [Parth](https://github.com/s0md3v/Parth)).
	* Probe for reflected parameters (inspired by [Tom Hudson](https://github.com/tomnomnom)'s [kxss](https://github.com/tomnomnom/hacks/tree/master/kxss)).

## Resources

* [Usage](#usage)
* [Installation](#installation)
	* [From Binary](#from-binary)
	* [From source](#from-source)
	* [From github](#from-github)
* [Contribution](#contribution)

## Usage

To display help message for sigurlscann3r use the `-h` flag:

```
$ sigurlscann3r -h

     _                  _                           _____
 ___(_) __ _ _   _ _ __| |___  ___ __ _ _ __  _ __ |___ / _ __
/ __| |/ _` | | | | '__| / __|/ __/ _` | '_ \| '_ \  |_ \| '__|
\__ \ | (_| | |_| | |  | \__ \ (_| (_| | | | | | | |___) | |
|___/_|\__, |\__,_|_|  |_|___/\___\__,_|_| |_|_| |_|____/|_| v1.0.0
       |___/

USAGE:
  sigurlscann3r [OPTIONS]

GENERAL OPTIONS:
  -iL                       input urls list (use `-iL -` to read from stdin)
  -threads                  number concurrent threads (default: 20)
  -update-params            update params file

HTTP OPTIONS:
  -delay                    delay between requests (default: 100ms)
  -follow-redirects         follow redirects (default: false)
  -follow-host-redirects    follow internal redirects i.e, same host redirects (default: false)
  -http-proxy               HTTP Proxy URL
  -timeout                  HTTP request timeout (default: 10s)
  -UA                       HTTP user agent

OUTPUT OPTIONS:
  -nC                       no color mode
  -oJ                       JSON output file (default: ./sigurlscann3r.json)
  -v                        verbose mode
```

## Installation

#### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/signedsecurity/sigurlscann3r/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

#### From Source

sigurlscann3r requires **go1.14+** to install successfully. Run the following command to get the repo

```bash
GO111MODULE=on go get -v -u github.com/signedsecurity/sigurlscann3r/cmd/sigurlscann3r
```

#### From Github

```
git clone https://github.com/signedsecurity/sigurlscann3r.git && \
cd sigurlscann3r/cmd/sigurlscann3r/ && \
go build . && \
mv sigurlscann3r /usr/local/bin/ && \
sigurlscann3r -h
```

## Contribution

[Issues](https://github.com/signedsecurity/sigurlscann3r/issues) and [Pull Requests](https://github.com/signedsecurity/sigurlscann3r/pulls) are welcome!