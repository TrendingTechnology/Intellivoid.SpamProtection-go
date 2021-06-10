/*
 * This file is part of Intellivoid.SpamProtection-go (https://github.com/Dank-del/Intellivoid.SpamProtection-go).
 * Copyright (c) 2021 Sayan Biswas.
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
	"github.com/Dank-del/Intellivoid.SpamProtection-go/SpamProtection"
	"log"
	"testing"
)

func TestAPI(t *testing.T) {
	responseID, err := SpamProtection.MakeRequestbyID(895373440)
	if err != nil {
		log.Fatal(err.Error())
	}
	if responseID.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyID)] Failed request, response code: %d", responseID.ResponseCode)
	}
	responseUsername, err := SpamProtection.MakeRequestbyUsername("dank_as_fuck")
	if err != nil {
		log.Fatal(err.Error())
	}
	if responseUsername.Success != true {
		t.Errorf("[Intellivoid.SpamProtection-go (MakeRequestbyUsername)] Failed request, response code: %d", responseUsername.ResponseCode)
	}
}