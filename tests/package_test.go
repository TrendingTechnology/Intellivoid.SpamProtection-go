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

func TestAPI(t *testing.T) {
	responseID, err := spamProtection.GetInfoByID(1181941155)
	if err != nil {
		log.Fatal(err.Error())
	}

	if responseID.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyID)] Failed request, response code: %d", responseID.ResponseCode)
	}

	responseUsername, err := spamProtection.GetInfoByUsername("Falling_inside_the_black")
	if err != nil {
		log.Fatal(err.Error())
	}

	if responseUsername.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", responseUsername.ResponseCode)
	}
}
