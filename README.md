<!--
 *
 * This file is part of Intellivoid.SpamProtection-go (https://github.com/Intellivoid/Intellivoid.SpamProtection-go).
 * Copyright (c) 2021 Sayan Biswas, ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
-->

# <img src="https://intellivoid.net/assets/favicon/android-chrome-192x192.png" width="35px" align="left"></img> SpamProtection-go 
> Name:		SpamProtection-go	\
> Version:	v1.0.4				\
> Edit:		13 Jun 2021			\
> By:		Intellivoid (C)	

-----------------------------------------------------------
<!-- https://intellivoid.net/assets/favicon/android-chrome-192x192.png

https://intellivoid.net/assets/media/TextLogo2.svg
-->

[![Go Reference](https://pkg.go.dev/badge/github.com/Intellivoid/Intellivoid.SpamProtection-go.svg)](https://pkg.go.dev/github.com/Intellivoid/Intellivoid.SpamProtection-go)

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
- Type safe; no weird `interface{}` logic.
- Multiple searching for blacklist flags.

<hr/>

## Getting started

You can easily download the library with the standard `go get` command:

```bash
go get github.com/Intellivoid/Intellivoid.SpamProtection-go
```
or
```bash
dep ensure -add github.com/Intellivoid/Intellivoid.SpamProtection-go
```

Full documentation of this library, can be found [here](https://pkg.go.dev/github.com/Intellivoid/Intellivoid.SpamProtection-go).

<hr/>

## How to use

The `spamProtection` package contains all helpers you need!
You can lookup a user, channel, group or a bot and recieve it's status with calling only one function. (using either username or user id) 

```go
package main

import "github.com/Intellivoid/Intellivoid.SpamProtection-go/spamProtection"

func main() {
	// get information of a user, channel or a group with a telegram id (int64)
	info, err := spamProtection.GetInfoByID(1234567890)
	if err != nil {
		log.Fatal(err.Error())
	}
	if info.IsBlacklisted() {
		// use GetType() method to see if the target is a user, 
		// channel, group or bot?
		log.Println("This " + info.GetType() + 
			" is blacklisted, because of " + info.GetBlacklistReason())
		// an example of the output would be:
		// This user is blacklisted, because of RAID Initializer / Participator
	}

	// get information of a user, channel or a group with a telegram username (string)
	// you can use also use '@' at the first of username
	info2, err := spamProtection.GetInfoByUsername("<username>")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Yup! it supports multiple searching
	// pass more than one flag and see if the target has one of them or not!
	if info2.HasFlag(spamProtection.SpamFlag, spamProtection.ChildAbuseFlag) {
		log.Println("this " + info2.GetType() +
			" is either a spammer or a child abuser")
		log.Println("it has " + info2.GetBlacklistFlag() + " flag!")
		// an example of the output would be:
		// this user is either a spammer or a child abuser
		// it has 0xSPAM flag!
	}

	// or see if a user is a potential spammer or not 
	// (based on their last activities)
	if info2.IsPotential() {
		log.Println("Be careful! this " + info2.GetType() +
			" is a potential spammer!")
		// an example for the output would be:
		// "Be careful! this bot is a potential spammer!"
	}

	// also make sure you are joined in our official  group
	if info2.IsVerified() {
		log.Println("This " + info2.GetType() +
			" is verified by Intellivoid Technologies")
		// an example for the output would be:
		// This group is verified by Intellivoid Technologies
	}
}
```

Still need more samples? Take a look at our [samples directory](samples).

<hr/>

## Support and Contributions

 * If you want to be aware of most recent changes, you can join [Intellivoid updates channel](https://t.me/Intellivoid).

 * Have a problem with API? Servers are down? Something went wrong? Ensure that you are joined at our [Support chat](https://t.me/IntellivoidDiscussions)!

 * Having a problem with library? Wanna talk with repository's owner? Contact the [Maintainer](https://t.me/dank_as_fuck)!

 * Still don't know what's going on? Not sure about how API works? Be sure to read [Introduction](https://docs.intellivoid.net/spamprotection/introduction).

 * Want to read original documentation? Want to see how we receive data from Servers? You can read [API documents](https://docs.intellivoid.net/spamprotection/v1/lookup) then!

 * Want to guarantee your group security? Protecting your groups from spammers is in high-priority for you? You can add our official [SpamProtection bot](https://t.me/SpamProtectionBot) with full features of our API!

 * If you think you have found a bug or have a feature request, feel free to use our [issue tracker](https://github.com/Intellivoid/Intellivoid.SpamProtection-go/issues). Before opening a new issue, please search to see if your problem has already been reported or not.  Try to be as detailed as possible in your issue reports.
 
 * If you need help using Intellivoid's APIs or have other questions we suggest you to join our [telegram community](https://t.me/IntellivoidCommunity).  Please do not use the GitHub issue tracker for personal support requests.

<hr/>

## Links

 * [Official website](https://intellivoid.net)
 * [Coffehouse](https://coffeehouse.intellivoid.net)
 * [Official channel](https://t.me/Intellivoid)
 * [Support chat](https://t.me/IntellivoidDiscussions)
 * [Intellivoid Community](https://t.me/IntellivoidCommunity)
 * [Discord server](https://discord.gg/euNkxEKPJb)
 * [Maintainer Telegram](https://t.me/dank_as_fuck)
 * [API Introduction](https://docs.intellivoid.net/spamprotection/introduction)
 * [API documents](https://docs.intellivoid.net/spamprotection/v1/lookup)
 * [SpamProtection bot](https://t.me/SpamProtectionBot)

<hr/>

## License

<img src="https://raw.githubusercontent.com/aliwoto/aliwoto/main/resources/Something_that_looks_like_Diamond.png" width="25px"></img> The Intellivoid.SpamProtection-go project is under the [GPL-3.0](https://opensource.org/licenses/GPL-3.0) license. You can find the license file [here](LICENSE).


<img src="https://intellivoid.net/assets/media/TextLogo2.svg" width="195px">
