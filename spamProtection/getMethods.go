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

import "strconv"

//---------------------------------------------------------

// Error returns you a string indicating what really happened
// at our servers.
func (e *Error) Error() string {
	str := "error in spamProtection package: "
	str += "with code " + strconv.Itoa(e.ErrorCode)
	str += " : " + e.Message

	return str
}

//---------------------------------------------------------

// HasError will check if there is an error in
// the response received from Intellivoid's servers.
func (r *APIResponse) HasError() bool {
	return r.Error != nil
}

// IsBlacklisted will check Attributes of the
// target, returns `true` if it's blacklisted,
// otherwise false.
func (r *APIResponse) IsBlacklisted() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsBlacklisted
}

// GetBlacklistFlag will return the BlacklistFlag in
// the Attributes of the target. it returns an empty
// string if the user is not blacklisted.
//  > see full list of flags here:
//  https://github.com/intellivoid/Intellivoid.SpamProtection-go/FLAGS.md
func (r *APIResponse) GetBlacklistFlag() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return r.Results.Attributes.BlacklistFlag
}

// HasFlag will check if the target has at least one of
// the passed flags or not. you can pass multiple flags to
// this method to check them all, at least one of them
// should exist, otherwise returns `false`.
func (r *APIResponse) HasFlag(flags ...string) bool {
	if flags == nil {
		return false
	}

	real := r.GetBlacklistFlag()
	if real == empty {
		return false
	}

	for _, f := range flags {
		if f == real {
			return true
		}
	}

	return false
}

// HasFlagNone will check if the Attributes of the target
// has None flag (0x0) or not.
func (r *APIResponse) HasFlagNone() bool {
	return r.HasFlag(NoneFlag)
}

// HasFlagNoneStr will check if the Attributes of the target
// has None flag (0x0) or not. It's equivalent to
// HasFlagNone(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagNoneStr() string {
	return yesNo(r.HasFlag(NoneFlag))
}

// HasFlagSpecial will check if the Attributes of the target
// has Special flag (0xSP) or not.
func (r *APIResponse) HasFlagSpecial() bool {
	return r.HasFlag(SpecialFlag)
}

// HasFlagSpecialStr will check if the Attributes of the target
// has Special flag (0xSP) or not. It's equivalent to
// HasFlagSpecial(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagSpecialStr() string {
	return yesNo(r.HasFlag(SpecialFlag))
}

// HasFlagSpam will check if the Attributes of the target
// has Spam flag (0xSPAM) or not.
func (r *APIResponse) HasFlagSpam() bool {
	return r.HasFlag(SpamFlag)
}

// HasFlagSpamStr will check if the Attributes of the target
// has Spam flag (0xSPAM) or not. It's equivalent to
// HasFlagSpam(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagSpamStr() string {
	return yesNo(r.HasFlag(SpamFlag))
}

// HasFlagEvade will check if the Attributes of the target
// has Evade( flag (0xEVADE) or not.
func (r *APIResponse) HasFlagEvade() bool {
	return r.HasFlag(EvadeFlag)
}

// HasFlagEvadeStr will check if the Attributes of the target
// has Evade( flag (0xEVADE) or not. It's equivalent to
// HasFlagEvade(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagEvadeStr() string {
	return yesNo(r.HasFlag(EvadeFlag))
}

// HasFlagChildAbuse will check if the Attributes of the target
// has child abuse flag (0xCACP) or not.
func (r *APIResponse) HasFlagChildAbuse() bool {
	return r.HasFlag(ChildAbuseFlag)
}

// HasFlagChildAbuseStr will check if the Attributes of the target
// has child abuse flag (0xCACP) or not. It's equivalent to
// HasFlagChildAbuse(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagChildAbuseStr() string {
	return yesNo(r.HasFlag(ChildAbuseFlag))
}

// HasFlagImper will check if the Attributes of the target
// has Imper flag (0xIMPER) or not.
func (r *APIResponse) HasFlagImper() bool {
	return r.HasFlag(ImperFlag)
}

// HasFlagImperStr will check if the Attributes of the target
// has Imper flag (0xIMPER) or not. It's equivalent to
// HasFlagImper(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagImperStr() string {
	return yesNo(r.HasFlag(ImperFlag))
}

// HasFlagPiracy will check if the Attributes of the target
// has piracy flag (0xPIRACY) or not.
func (r *APIResponse) HasFlagPiracy() bool {
	return r.HasFlag(PiracyFlag)
}

// HasFlagPiracy will check if the Attributes of the target
// has piracy flag (0xPIRACY) or not. It's equivalent to
// HasFlagPiracy(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagPiracyStr() string {
	return yesNo(r.HasFlag(PiracyFlag))
}

// HasFlagNSFW will check if the Attributes of the target
// has NSFW flag (0xNSFW) or not.
func (r *APIResponse) HasFlagNSFW() bool {
	return r.HasFlag(NSFWFlag)
}

// HasFlagNSFW will check if the Attributes of the target
// has NSFW flag (0xNSFW) or not. It's equivalent to
// HasFlagNSFW(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagNSFWStr() string {
	return yesNo(r.HasFlag(NSFWFlag))
}

// HasFlagPrivate will check if the Attributes of the target
// has private flag (0xPRIVATE) or not.
func (r *APIResponse) HasFlagPrivate() bool {
	return r.HasFlag(PrivateFlag)
}

// HasFlagPrivateStr will check if the Attributes of the target
// has private flag (0xPRIVATE) or not. It's equivalent to
// HasFlagPrivate(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagPrivateStr() string {
	return yesNo(r.HasFlag(PrivateFlag))
}

// HasFlagRaid will check if the Attributes of the target
// has raid flag (0xRAID) or not.
func (r *APIResponse) HasFlagRaid() bool {
	return r.HasFlag(RaidFlag)
}

// HasFlagRaidStr will check if the Attributes of the target
// has raid flag (0xRAID) or not. It's equivalent to
// HasFlagRaid(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagRaidStr() string {
	return yesNo(r.HasFlag(RaidFlag))
}

// HasFlagScam will check if the Attributes of the target
// has scam flag (0xSCAM) or not.
func (r *APIResponse) HasFlagScam() bool {
	return r.HasFlag(ScamFlag)
}

// HasFlagScamStr will check if the Attributes of the target
// has scam flag (0xSCAM) or not. It's equivalent to
// HasFlagScam(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagScamStr() string {
	return yesNo(r.HasFlag(ScamFlag))
}

// HasFlagMassAdd will check if the Attributes of the target
// has mass add flag (0xMASSADD) or not.
func (r *APIResponse) HasFlagMassAdd() bool {
	return r.HasFlag(MassAddFlag)
}

// HasFlagMassAddStr will check if the Attributes of the target
// has mass add flag (0xMASSADD) or not. It's equivalent to
// HasFlagMassAdd(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagMassAddStr() string {
	return yesNo(r.HasFlag(MassAddFlag))
}

// HasFlagNameSpam will check if the Attributes of the target
// has name spam flag (0xNAMESPAM) or not.
func (r *APIResponse) HasFlagNameSpam() bool {
	return r.HasFlag(NameSpamFlag)
}

// HasFlagNameSpamStr will check if the Attributes of the target
// has name spam flag (0xNAMESPAM) or not. It's equivalent to
// HasFlagNameSpam(), but it returns "Yes" for `true`,
// and "No" for `false`.
func (r *APIResponse) HasFlagNameSpamStr() string {
	return yesNo(r.HasFlag(NameSpamFlag))
}

// GetBlacklistReason will return the reason of
// blacklisting of the target. if the target is not
// blacklisted in the first place, it will return an
// empty string.
func (r *APIResponse) GetBlacklistReason() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return r.Results.Attributes.BlacklistReason
}

// GetOriginId return the original private id of a user.
// if the user was blacklisted with 0xEVADE flag,
// then the original private telegram ID (ptid)
// will be available, otherwise it will be an empty string.
func (r *APIResponse) GetOriginId() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return r.Results.Attributes.OriginalPrivateID
}

// IsPotential will return `true` if the target
// is a potential spammer (based on past activities).
func (r *APIResponse) IsPotential() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsPotentialSpammer
}

// IsPotentialStr will return `"Yes"` if the target
// is a potential spammer (based on past activities),
// otherwise returns `"No".`
func (r *APIResponse) IsPotentialStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IsPotentialSpammer)
}

// IsOperator returns `true` if this user is an operator
// of Spam Protection.
func (r *APIResponse) IsOperator() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsOperator
}

// IsOperatorStr returns `"Yes"` if this user is an operator
// of Spam Protection, otherwise returns `"No"`.
func (r *APIResponse) IsOperatorStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IsOperator)
}

// IsAgent will check if this user is an agent or not.
// Agents are basically userbots that spies on groups
// where SpamProtectionBot isn't in and or resolve users
// that the bot doesn't know.
func (r *APIResponse) IsAgent() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsAgent
}

// IsAgentStr will check if this user is an agent or not.
// Agents are basically userbots that spies on groups
// where SpamProtectionBot isn't in and or resolve users
// that the bot doesn't know.
// it returns `"Yes"` if the user is an agent, otherwise
// `"No"`.
func (r *APIResponse) IsAgentStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IsAgent)
}

// IsWhitelisted will check if this user is a whitelisted
// use or not. if this user is whitelisted, it annot be blacklisted
// or seen as a potential spammer
func (r *APIResponse) IsWhitelisted() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsWhitelisted
}

// IsWhitelistedStr will check if this user is a whitelisted
// use or not. if this user is whitelisted, it cannot be blacklisted
// or seen as a potential spammer.
// it returns `"Yes"` if the user is whitelisted,
// otherwise returns `"No"`.
func (r *APIResponse) IsWhitelistedStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IsWhitelisted)
}

// IsVerified returns `true` if the target
// (which can be user, group and channel), is
// verified by intellivoid or not.
func (r *APIResponse) IsVerified() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IntellivoidAccountsVerified
}

// IsVerified returns `"Yes"` if the target
// (which can be user, group and channel), is
// verified by intellivoid or not, otherwise
// returns `"No"`.
func (r *APIResponse) IsVerifiedStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IntellivoidAccountsVerified)
}

// IsOfficial will check if the target is official or not.
// There is no such thing as "Official users",
// only "Official channels" or groups,
// these are entities with the  "blue checkmark"
// implemented by us; not by Telegram.
// To show the public that this entity is the real thing.
func (r *APIResponse) IsOfficial() bool {
	if r.Results.Attributes == nil {
		return false
	}

	return r.Results.Attributes.IsOfficial
}

// IsOfficial will check if the target is official or not.
// There is no such thing as "Official users",
// only "Official channels" or groups,
// these are entities with the  "blue checkmark"
// implemented by us; not by Telegram.
// To show the public that this entity is the real thing.
// it returns `"Yes"` if it's official, otherwise returns
// `"No"`.
func (r *APIResponse) IsOfficialStr() string {
	if r.Results.Attributes == nil {
		return empty
	}

	return yesNo(r.Results.Attributes.IsOfficial)
}

// GetType will return the type of the target.
// it can be user, bot, group, supergroup or channel.
func (r *APIResponse) GetType() string {
	if r.Results == nil {
		return empty
	}

	return r.Results.EntityType
}

// IsUser will check if the target is a
// telegram user account or not.
func (r *APIResponse) IsUser() bool {
	return r.GetType() == UserType
}

// IsBot will check if the target is a
// bot or not.
func (r *APIResponse) IsBot() bool {
	return r.GetType() == BotType
}

// IsGroup will check if the target is
// either a super group or a normal group.
func (r *APIResponse) IsGroup() bool {
	t := r.GetType()
	return t == GroupType ||
		t == SupergroupType
}

// IsNormalGroup will check if the target is a
// normal telegram group or not.
func (r *APIResponse) IsNormalGroup() bool {
	return r.GetType() == GroupType
}

// IsSuperGroup will check if the target is a
// super group or not.
func (r *APIResponse) IsSuperGroup() bool {
	return r.GetType() == SupergroupType
}

// IsChannel will check if the target is a
// telegram channel or not.
func (r *APIResponse) IsChannel() bool {
	return r.GetType() == ChannelType
}

//---------------------------------------------------------
