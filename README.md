# hqurlscann3r

[![release](https://img.shields.io/github/release/hueristiq/hqurlscann3r?style=flat&color=0040ff)](https://github.com/hueristiq/hqurlscann3r/releases) ![maintenance](https://img.shields.io/badge/maintained%3F-yes-0040ff.svg) [![open issues](https://img.shields.io/github/issues-raw/hueristiq/hqurlscann3r.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqurlscann3r/issues?q=is:issue+is:open) [![closed issues](https://img.shields.io/github/issues-closed-raw/hueristiq/hqurlscann3r.svg?style=flat&color=0040ff)](https://github.com/hueristiq/hqurlscann3r/issues?q=is:issue+is:closed) [![license](https://img.shields.io/badge/license-MIT-gray.svg?colorB=0040FF)](https://github.com/hueristiq/hqurlscann3r/blob/master/LICENSE) [![twitter](https://img.shields.io/badge/twitter-@itshueristiq-0040ff.svg)](https://twitter.com/itshueristiq)

A web application attack surface mapping tool. It takes in a list of urls then performs numerous probes

## Resources

* [Features](#features)
* [Installation](#installation)
	* [From Binary](#from-binary)
	* [From source](#from-source)
	* [From github](#from-github)
* [Usage](#usage)
* [Contribution](#contribution)

## Features

* Categorize URLs

	<details>
	<summary>URLs' categories</summary>

	```
	- endpoint
	- js {js}
	- style {css}
	- data {json|xml|csv}
	- archive {zip|tar|tar.gz}
	- doc {pdf|xlsx|doc|docx|txt}
	- media {jpg|jpeg|png|ico|svg|gif|webp|mp3|mp4|woff|woff2|ttf|eot|tif|tiff}
	```

	</details>

* Probe HTTP requests for `status_code`, `content_type`, e.t.c
* For every URL of category `endpoint` with a query:
	* Probe for commonly vulnerable parameters (inspired by [Somdev Sangwan](https://github.com/s0md3v)'s [Parth](https://github.com/s0md3v/Parth)).
	* Probe for reflected parameters (inspired by [Tom Hudson](https://github.com/tomnomnom)'s [kxss](https://github.com/tomnomnom/hacks/tree/master/kxss)).

## Installation

### From Binary

You can download the pre-built binary for your platform from this repository's [releases](https://github.com/hueristiq/hqurlscann3r/releases/) page, extract, then move it to your `$PATH`and you're ready to go.

### From Source

hqurlscann3r requires **go1.17+** to install successfully. Run the following command to get the repo

```bash
go install -v github.com/hueristiq/hqurlscann3r/cmd/hqurlscann3r@latest
```

### From Github

```bash
git clone https://github.com/hueristiq/hqurlscann3r.git && \
cd hqurlscann3r/cmd/hqurlscann3r/ && \
go build . && \
mv hqurlscann3r /usr/local/bin/ && \
hqurlscann3r -h
```

## Usage

To display help message for hqurlscann3r use the `-h` flag:

```bash
hqurlscann3r -h
```

```text
 _                      _                           _____
| |__   __ _ _   _ _ __| |___  ___ __ _ _ __  _ __ |___ / _ __
| '_ \ / _` | | | | '__| / __|/ __/ _` | '_ \| '_ \  |_ \| '__|
| | | | (_| | |_| | |  | \__ \ (_| (_| | | | | | | |___) | |
|_| |_|\__, |\__,_|_|  |_|___/\___\__,_|_| |_|_| |_|____/|_| v1.0.0
          |_|

USAGE:
  hqurlscann3r [OPTIONS]

OPTIONS:
   -c, --concurrency              concurrency level (default: 20)
   -d, --delay                    delay between requests (default: 100ms)
       --follow-redirects         follow redirects (default: false)
       --follow-host-redirects    follow internal redirects i.e, same host redirects (default: false)
       --http-proxy               HTTP Proxy URL
  -iL, --input-list               input urls list
  -nC, --no-color                 no color mode
   -o, --output                   JSON output file (default: ./hqurlscann3r.json)
   -t, --timeout                  HTTP request timeout (default: 10s)
  -ua, --user-agent               HTTP user agent
       --update-params            update params file
   -v, --verbose                  verbose mode
```

## Contribution

[Issues](https://github.com/hueristiq/hqurlscann3r/issues) and [Pull Requests](https://github.com/hueristiq/hqurlscann3r/pulls) are welcome!