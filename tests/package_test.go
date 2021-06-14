/*
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
 */

package main

import (
	"log"
	"testing"

	"github.com/Intellivoid/Intellivoid.SpamProtection-go/spamProtection"
)

func TestFix(t *testing.T) {
	str := spamProtection.FixUsername("https://t.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("http://t.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("t.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("bullshit.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("twitter.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("whatsapp.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("whatsapp.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("facebook.me/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("I don't know what blah blag /NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("everything is supported/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("Intellivoid.net/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("  @ NightShadowsHangout  ")
	log.Println(str)

	str = spamProtection.FixUsername("@NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("http://telesco.pe/NightShadowsHangout")
	log.Println(str)

	str = spamProtection.FixUsername("http://unknown_host/g/215600")
	log.Println(str)
}
