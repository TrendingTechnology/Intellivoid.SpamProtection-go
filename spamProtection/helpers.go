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

package spamProtection

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// This function allows you to lookup a user, chat or channel
// by its id (which is a 64 bit integer)
// in the SpamProtection database, this method does not require
// authentication and there are no rate limits.
// if something went wrong in server-side and we get the
// full information in response, it will return you that
// `error`, so the `info` value will be nil.
// negative IDs (like group's id) are supported.
func GetInfoByID(id int64) (info *APIResponse, err error) {
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

	if res.HasError() {
		// if something went wrong at server-side,
		// return it so the user can find out that
		// duty of this function ended in failure.
		return nil, res.Error
	}

	return &res, nil
}

// This function allows you to lookup a user, chat or channel
// by its username (which is a string)
// in the SpamProtection database, this method does not require
// authentication and there are no rate limits.
// if something went wrong in server-side and we get the
// full information in response, it will return you that
// `error`, so the `info` value will be nil.
// it supports these formats:
//  * the_username
//  * @the_username
//  * t.me/falling_inside_the_black
//  * http://telesco.pe/NightShadowsHangout
func GetInfoByUsername(username string) (info *APIResponse, err error) {
	if len(username) == 0 {
		return nil, errors.New("username cannot be empty")
	}

	// fix the username so it doesn't containt white space and '@'
	username = FixUsername(username)
	if !IsValidUsername(username) {
		// always check if the username is valid or not.
		// sending http request blindly without any checking
		// will only increase running time order.
		return nil, errors.New("the username \"" +
			username + "\" is invalid")
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

	if res.HasError() {
		// if something went wrong at server-side,
		// return it so the user can find out that
		// duty of this function ended in failure.
		return nil, res.Error
	}

	return &res, nil
}

// IsValidUsername will check if `text` is a valid username or not.
// please fix the string before using this function,
// for example you have to do TrimPrefix, TrimSuffix, etc...
func IsValidUsername(text string) bool {
	l := len(text)
	if l <= minUsername || l > maxUsername {
		return false
	}

	for i, c := range text {
		// when we are in the base index (0)
		// the character should be english,
		// otherwise return false
		if i == baseIndex {
			if !isEnglish(c) {
				return false
			}

			continue
		}

		// check if it's english, number or underline ('_'),
		// if it's not return false
		if !isNumOrEng(c) {
			return false
		}
	}

	// all checks have been passes, so return true
	return true
}

// FixUsername will fix the username and will return you a valid
// username.
// examples:
//  " @Falling_inside_the_black  " => "Falling_inside_the_black"
//  "t.me/dank_as_fuck " => "dank_as_fuck"
func FixUsername(value string) string {
	if len(value) == 0 {
		return empty
	}

	value = strings.TrimSpace(value)
	value = strings.TrimPrefix(value, atSign)
	value = strings.TrimSpace(value)

	if strings.Contains(value, slash) {
		strs := strings.Split(value, slash)
		l := len(strs)
		if l == baseIndex {
			return empty
		}

		l--
		return strs[l]
	}

	return value
}

// isEnglish return `true` if `r` is an english letter,
// otherwise `false`.
func isEnglish(r rune) bool {
	if r >= letterA && r <= letterZ {
		return true
	} else {
		return r >= capLetterA && r <= capLetterZ
	}
}

// isNum returns `true` if `r` is considered as a
// number rune, otherwise `false`.
func isNum(r rune) bool {
	return r >= letterZero || r <= letterNine
}

// isNumOrEng will check if `r` is rather english,
// number or underline ('_') rune or not.
// return `true` if yes, otherwise `false`.
func isNumOrEng(r rune) bool {
	return isNum(r) ||
		isEnglish(r) ||
		r == underlineChar
}

// yesNo returns "Yes" for `true`,
// and "No" for `false`.
func yesNo(value bool) string {
	if value {
		return yesStr
	} else {
		return noStr
	}
}
