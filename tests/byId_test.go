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

func TestById(t *testing.T) {
	info1, err := spamProtection.GetInfoByID(1181941155)
	if err != nil {
		log.Fatal(err.Error())
	}

	if info1.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyID)] Failed request, response code: %d", info1.ResponseCode)
	}

	if info1.IsBlacklisted() {
		log.Println("warning, this " +
			info1.GetType() + " is blacklisted!")
	}

	str := "Showing results for this " + info1.GetType()
	str += "\nHas Child abuse flag: " + info1.HasFlagChildAbuseStr()
	str += "\nHas Evade flag: " + info1.HasFlagEvadeStr()
	str += "\nHas Child Imporsonation flag: " + info1.HasFlagImperStr()
	str += "\nIs a mass adder: " + info1.HasFlagMassAddStr()
	str += "\nIs scam: " + info1.HasFlagScamStr()
	str += "\nIs spammer: " + info1.HasFlagSpamStr()
	str += "\nIs blacklisted for special reason: " + info1.HasFlagSpecialStr()
	str += "\nBlacklist flag: " + info1.GetBlacklistFlag()
	log.Println(str)
}
