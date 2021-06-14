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

// the endpoint used to get responses from Intellivoid servers.
// for more information please read our documents:
//  https://docs.intellivoid.net/spamprotection/v1/lookup
const (
	endPoint  = "https://api.intellivoid.net/spamprotection/v1/lookup?query=%"
	endPointI = endPoint + "d"
	endPointS = endPoint + "s"
)

// telegram limitations constants
const (
	baseIndex   = 0x0
	minUsername = 4
	maxUsername = 32
	empty       = ""
)

// Intellivoid Spam protection flags
const (
	// None or an Undefined reason
	NoneFlag = "0x0"

	// Special reason, like consulting operators
	SpecialFlag = "0xSP"

	// Spam or Unwanted promotion
	SpamFlag = "0xSPAM"

	// Ban Evade using alt accounts
	EvadeFlag = "0xEVADE"

	// Child pornography or Child abuse
	ChildAbuseFlag = "0xCACP"

	// Malicious Impersonation
	ImperFlag = "0xIMPER"

	// Promote or spam of pirated content
	PiracyFlag = "0xPIRACY"

	// Promote or spam of NSFW content
	NSFWFlag = "0xNSFW"

	// Unsolicited Spam or Promote
	PrivateFlag = "0xPRIVATE"

	// Raid initializer or Participator
	RaidFlag = "0xRAID"

	// Scamming
	ScamFlag = "0xSCAM"

	// Mass adding users to group or channels
	MassAddFlag = "0xMASSADD"

	// Promotion or spam via name or bio
	NameSpamFlag = "0xNAMESPAM"
)

// The entity type constants,
// can be user, bot, group, supergroup or channel
const (
	UserType       = "user"
	BotType        = "bot"
	GroupType      = "group"
	SupergroupType = "supergroup"
	ChannelType    = "channel"
)

// another string constants
const (
	yesStr = "Yes"
	noStr  = "No"
)

// prefixes used in fixing a username
const (
	atSign = "@"
	slash  = "/"
)

const (
	underlineChar = '_' // _
	letterA       = 'a' // a
	letterZ       = 'z' // z
	capLetterA    = 'A' // A
	capLetterZ    = 'Z' // Z
	letterZero    = '0' // 0
	letterNine    = '9' // 9
)
