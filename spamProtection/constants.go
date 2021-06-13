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
	minUsername = 4
	maxUsername = 32
	empty       = ""
)

// Intellivoid Spam protection flags
const (
	NoneFlag       = "0x0"
	SpecialFlag    = "0xSP"
	SpamFlag       = "0xSPAM"
	EvadeFlag      = "0xEVADE"
	ChildAbuseFlag = "0xCACP"
	ImperFlag      = "0xIMPER"
	PiracyFlag     = "0xPIRACY"
	NSFWFlag       = "0xNSFW"
	PrivateFlag    = "0xPRIVATE"
	RaidFlag       = "0xRAID"
	ScamFlag       = "0xSCAM"
	MassAddFlag    = "0xMASSADD"
	NameSpamFlag   = "0xNAMESPAM"
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
