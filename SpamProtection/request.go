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

package SpamProtection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MakeRequestbyID(id int) (info *APIResponse, err error) {
	if id == 0 {
		return nil, errors.New("ID cannot be empty")
	}

	addr := fmt.Sprintf("https://api.intellivoid.net/spamprotection/v1/lookup?query=%d", id)
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

func MakeRequestbyUsername(username string) (info *APIResponse, err error) {
	if len(username) == 0 {
		return nil, errors.New("username cannot be empty")
	}

	addr := fmt.Sprintf("https://api.intellivoid.net/spamprotection/v1/lookup?query=%s", username)
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

