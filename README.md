# <img src="https://intellivoid.net/assets/favicon/android-chrome-192x192.png" width="35px" align="left"></img> SpamProtection-go 
> Name:		SpamProtection-go			\
> Version:	v1.0.2				\
> Edit:		11 Jun 2021			\
> By:		Intellivoid (C)	

-----------------------------------------------------------
<!-- https://intellivoid.net/assets/favicon/android-chrome-192x192.png

https://intellivoid.net/assets/media/TextLogo2.svg
-->

[![Go Reference](https://pkg.go.dev/badge/github.com/Dank-del/Intellivoid.SpamProtection-go.svg)](https://pkg.go.dev/github.com/Dank-del/Intellivoid.SpamProtection-go)

SpamProtection-Go is an Official [golang](https://go.dev) wrapper for [Intellivoid](https://intellivoid.net) SpamProtection API, which is fast, secure and requires no additional packages to be installed.

### Preview:

 * [What is SpamProtection?](#what-is-spamprotection)
 * [Supported Go versions](#supported-go-versions)
 * [Features](#features)
 * [Getting started](#getting-started)
 * [How to use](#how-to-use)
 * [Support and Contributions](#support-and-contributions)
 * [Links](#links)
 * [License](#license)


<img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="50px"></img><img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="50px"></img><img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="50px"></img><img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="50px"></img><img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/798246901916499998.gif" width="50px"></img>

<hr/>


## What is SpamProtection?

SpamProtection is a community powered solution with the goal to effectively combat spam on Telegram using machine learning and manual spam reporting.
This API allows you to lookup a user, chat or channel in the SpamProtection database, this method does not require authentication and there are no rate limits. 

<hr/>

## Supported Go versions

For better experience, we recommend you to use latest version of Go (v1.16), but this library requires at least Go version 1.15.

<hr/>

## Features

- Uses official Intellivoid API, which makes this library:
   - Easy to update
   - Guaranteed to match the docs
   - No third party endpoints
   - Self-documenting (Contains all pre-existing Intellivoid's docs)
- It's in pure go, no need to install any kind of plugin or include any kind of additional files.
- No third party library bloat; only uses standard library.
- Type safe; no weird interface{} logic.

<hr/>

## Getting started

You can easily download the library with the standard `go get` command:

```bash
go get github.com/Dank-del/Intellivoid.SpamProtection-go
```
or
```bash
dep ensure -add github.com/Dank-del/Intellivoid.SpamProtection-go
```

Full documentation of this library, can be found [here](https://pkg.go.dev/github.com/Dank-del/Intellivoid.SpamProtection-go).

<hr/>

## How to use

The `spamProtection` package contains all helpers you need!

```go
package main

import "github.com/Dank-del/Intellivoid.SpamProtection-go/spamProtection"

func main() {
	// get information of a user, channel or a group with a telegram id (int64)
	respByID, err := spamProtection.MakeRequestbyID(1234567890)
	if err != nil {
		log.Fatal(err.Error())
	}
	// use your respByID...

	// // get information of a user, channel or a group with a telegram username (string)
	respUname, err := spamProtection.MakeRequestbyUsername("<username>")
	if err != nil {
		log.Fatal(err.Error())
	}
	// use your respUname...
}
```

<hr/>

## Support and Contributions

 * If you want to be aware of most recent changes, you can join [Intellivoid updates channel](https://t.me/Intellivoid).
 * Have a problem with API? Servers are down? Somethign went wrong? Ensure that you are joined at our [Support chat](https://t.me/IntellivoidDiscussions)!
 * Have a personal problem with library? Feel lonely? Have to talk with repository's owner? Contact the [Maintainer](https://t.me/dank_as_fuck)!
 * Still don't know what's going on? Not sure about how API works? Be sure to read [Introduction](https://docs.intellivoid.net/spamprotection/introduction).
 * Want to read original documentation? Want to see how we receive data from Servers? You can read [API documents](https://docs.intellivoid.net/spamprotection/v1/lookup) then!
 * Want to guarantee your group security? Want to protect your groups from spammers? You can add our official [SpamProtection bot](https://t.me/SpamProtectionBot) with full features of our API!
 * If you think you have found a bug or have a feature request, feel free to use our [issue tracker](https://github.com/Dank-del/Intellivoid.SpamProtection-go/issues). Before opening a new issue, please search to see if your problem has already been reported or not.  Try to be as detailed as possible in your issue reports.
 * If you need help using Intellivoid's APIs or have other questions we suggest you to join our [telegram community](https://t.me/IntellivoidCommunity).  Please do not use the GitHub issue tracker for personal support requests.

<hr/>

## Links

 * [Official website](https://intellivoid.net)
 * [Coffehouse](https://coffeehouse.intellivoid.net)
 * [Official channel](https://t.me/Intellivoid)
 * [Support chat](https://t.me/IntellivoidDiscussions)
 * [Intellivoid Community](https://t.me/IntellivoidCommunity)
 * [Discord server](https://discord.gg/euNkxEKPJb)
 * [Netkas](https://t.me/Netkas)
 * [API Introduction](https://docs.intellivoid.net/spamprotection/introduction)
 * [API documents](https://docs.intellivoid.net/spamprotection/v1/lookup)
 * [SpamProtection bot](https://t.me/SpamProtectionBot)

<hr/>

## License

<img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/Something_that_looks_like_Diamond.png" width="25px"></img> The Intellivoid.SpamProtection-go project is under the [GPL3.0](https://opensource.org/licenses/GPL-3.0). You can find the license file [here](LICENSE).


<img src="https://intellivoid.net/assets/media/TextLogo2.svg" width="195px">
