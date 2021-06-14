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

func TestByUname1(t *testing.T) {
	info1, err := spamProtection.GetInfoByUsername("Falling_inside_the_black")
	if err != nil {
		log.Fatal(err.Error())
	}

	if info1.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", info1.ResponseCode)
	}
}

func TestByUname2(t *testing.T) {
	info1, err := spamProtection.GetInfoByUsername("I don't know what blah blag /NightShadowsHangout")
	if err != nil {
		log.Fatal(err.Error())
	}

	if info1.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", info1.ResponseCode)
	}
}

func TestByUname3(t *testing.T) {
	info1, err := spamProtection.GetInfoByUsername("facebook.me/NightShadowsHangout")
	if err != nil {
		log.Fatal(err.Error())
	}

	if info1.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", info1.ResponseCode)
	}
}

func TestByUname4(t *testing.T) {
	info1, err := spamProtection.GetInfoByUsername("facebook.me/NightShadowsHangout")
	if err != nil {
		log.Fatal(err.Error())
	}

	if info1.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", info1.ResponseCode)
	}
}
