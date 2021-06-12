/*
 * This file is part of Intellivoid.SpamProtection-go (https://github.com/Intellivoid/Intellivoid.SpamProtection-go).
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

package spamProtection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// This function allows you to lookup a user, chat or channel
// by its id (which is a 64 bit integer)
// in the SpamProtection database, this method does not require
// authentication and there are no rate limits.
// this function will return you error if and only if there is
// an error during sending the request to the servers.
// if an error is in server-side and we get the full information,
// it will return you `nil` with the `info` value
// (and info value will contains the error).
func MakeRequestbyID(id int64) (info *APIResponse, err error) {
	if id == 0 {
		return nil, errors.New("ID cannot be zero")
	}

	addr := fmt.Sprintf(endPointI, id)
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res APIResponse

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// This function allows you to lookup a user, chat or channel
// by its username (which is a string)
// in the SpamProtection database, this method does not require
// authentication and there are no rate limits.
// this function will return you error if and only if there is
// an error during sending the request to the servers.
// if an error is in server-side and we get the full information,
// it will return you `nil` with the `info` value
// (and info value will contains the error).
func MakeRequestbyUsername(username string) (info *APIResponse, err error) {
	if len(username) == 0 {
		return nil, errors.New("username cannot be empty")
	}

	addr := fmt.Sprintf(endPointS, username)
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res APIResponse

	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
