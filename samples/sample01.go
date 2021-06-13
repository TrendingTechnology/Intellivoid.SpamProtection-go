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
package samples

import (
	"log"

	"github.com/Intellivoid/Intellivoid.SpamProtection-go/spamProtection"
)

func UsernameSample01() {
	info, err := spamProtection.GetInfoByUsername("@Falling_inside_the_black")
	if err != nil {
		log.Fatal(err)
	}

	if info.Success {
		log.Println("successed to get information of use from " +
			"intellivoid's databases")
	}

	if info.IsBlacklisted() {
		log.Println("target is blacklisted")
		f := info.GetBlacklistFlag()
		log.Println("with falg: " + f)
	}

	if info.IsVerified() {
		log.Println("target is verified by Intellivoid")
	}

	if info.HasFlag(spamProtection.NSFWFlag, spamProtection.ChildAbuseFlag) {
		log.Println("Oh no, the target has either NSFW flag or child abuse")
	}

	if info.HasFlagEvade() {
		log.Println("the user has a evade flag")
	}

	log.Println("The type is: " + info.Results.EntityType)
}

func IdSample01() {
	info, err := spamProtection.GetInfoByID(1234567890)
	if err != nil {
		log.Fatal(err)
	}

	if info.Success {
		log.Println("successed to get information of use from " +
			"intellivoid's database")
	}

	if info.IsBlacklisted() {
		log.Println("target is blacklisted")
		f := info.GetBlacklistFlag()
		log.Println("with falg: " + f)
	}

	if info.IsOfficial() {
		log.Println("target is an official Intellivoid " +
			info.GetType())
	}

	if info.HasFlag(spamProtection.NameSpamFlag, spamProtection.ChildAbuseFlag) {
		log.Println("Oh no, the target has either" +
			" name spam flag or child abuse")
	}

	if info.HasFlagPrivate() {
		// for example for a user, this log will show us:
		//  this user has a private flag
		log.Println("this " + info.GetType() + " has a private flag")
	}

	log.Println("The type is: " + info.Results.EntityType)
	// or we can use:
	log.Println("The type is: " + info.GetType())
	// both of them will give us same result

	//**********printing some information******************
	str := "Information about this " + info.GetType()
	str += "Has Child abuse flag: " + info.HasFlagChildAbuseStr()
	str += "Has Evade flag: " + info.HasFlagEvadeStr()
	str += "Has Child Imporsonation flag: " + info.HasFlagImperStr()
	str += "Is a mass adder: " + info.HasFlagMassAddStr()
	str += "Is scam: " + info.HasFlagScamStr()
	str += "Is spammer: " + info.HasFlagSpamStr()
	str += "Is blacklisted for special reason: " + info.HasFlagSpecialStr()
}
