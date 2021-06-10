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

package spamProtection

// the endpoint used to get responses from Intellivoid servers.
// for more information please read our documents:
//  https://docs.intellivoid.net/spamprotection/v1/lookup
const (
	endPoint  = "https://api.intellivoid.net/spamprotection/v1/lookup?query=%"
	endPointI = endPoint + "d"
	endPointS = endPoint + "s"
)