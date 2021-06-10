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

// Top-Level Struct.
// The value success will always be returned
// If it's true then results will be available as an object.
// If false the error object will return and
// it is pretty much the same across all of our APIs
type APIResponse struct {
	Success      bool     `json:"success"`
	ResponseCode int      `json:"response_code"`
	Results      *Results `json:"results"`
	Error        *Error   `json:"error"`
}

// Error contains an Error code, Type and Message.
// The way your client should handle API errors is by this;
// If applicable, apply the error code from
// the API to your exception + the message.
// For errors that you want to catch and resolve on your own,
// refer to the error code. For this case;
// SpamProtection only has one error code which is 0.
// Other more complicated APIs has a set of pre-determined errors
// that you can use to programmatically detect the error and
// resolve it if it's on your client's end.
// 0 Is pretty much telling you to read the message
// the server returned, because it's unique or something weird.
// Before the API would simply return a message and it
// is up to your client to determine the error
// by the HTTP response code;
// this isn't good because if we were to
// add more complicated methods in the future that
// returns different errors but with the same HTTP code,
// it can be harder for programs to automatically correct the error
// if required.
// So the 1.1 API update corrects this by using this new error
// standard which returns `error_code`, `type` and `message`.
type Error struct {
	ErrorCode int    `json:"error_code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type SpamPrediction struct {
	// The probability (confidence) of this generalized ham
	// (not spam) prediction, if no prediction is
	// available then it will be null
	HamPrediction float64 `json:"ham_prediction"`

	// The probability (confidence) of this generalized
	// spam prediction, if no prediction is available
	// then it will be zero
	SpamPrediction float64 `json:"spam_prediction"`
}

type LanguagePrediction struct {
	// The ISO 639-1 language code predicted by SpamProtectionBot,
	// if no prediction is available then it will be
	// an empty string
	Language string `json:"language"`

	// The probability (confidence) of this generalized
	// language prediction, if no prediction is available
	// then it will be zero
	Probability float64 `json:"probability"`
}

// There is no such thing as "Official users",
// only "Official channels" or groups,
// these are entities with the  "blue checkmark"
// implemented by us; not by Telegram.
// To show the public that this entity is the real thing.
//
// Operators are the people that maintains the SpamProtectionBot
// database by handling user appeals, blacklisting spammers
// and making sure that the bot is operating correctly.
// Basically the Administrators of SpamProtectionBot
//
// Agents are basically userbots that spies on groups
// where SpamProtectionBot isn't in and or resolve users
// that the bot doesn't know.
//
// There are no sudo users or sudo capabilities;
// operators cannot use the bot to ban people from your group,
// manage your group and or promote themselves to an administrator.
// Operators do not need to be administrators to preform
// their tasks since it's related to the database itself,
// not your group.
type Attributes struct {

	// Indicates if this entity was blacklisted by a operator or agent.
	IsBlacklisted bool `json:"is_blacklisted"`

	// If the entity is blacklisted, the blacklist flag will be
	//  available otherwise it will be null
	BlacklistFlag string `json:"blacklist_flag"`

	// If the entity is blacklisted, a user friendly message
	// of the blacklist flag will be available otherwise
	// it will be null
	BlacklistReason string `json:"blacklist_reason"`

	// If the entity was blacklisted for 0xEVADE then the
	// original private telegram ID (ptid) will be available,
	// otherwise it will be null
	OriginalPrivateID string `json:"original_private_id"`

	// Indicates if SpamProtection believes that
	// this user may be a potential spammer based on past activities
	IsPotentialSpammer bool `json:"is_potential_spammer"`

	// Indicates if this user is an operator for SpamProtectionBot
	IsOperator bool `json:"is_operator"`

	// Indicates if this user is an agent that automatically
	// reports spam to SpamProtectionBot (Userbot Agent)
	IsAgent bool `json:"is_agent"`

	// Indicates if this user is whitelisted and c
	// annot be blacklisted or seen as a potential spammer
	IsWhitelisted bool `json:"is_whitelisted"`

	// Indicates if this user has verified their
	// Telegram Account with Intellivoid Accounts
	IntellivoidAccountsVerified bool `json:"intellivoid_accounts_verified"`

	// Indicates if this entity is deemed
	//official by Intellivoid Technologies
	IsOfficial bool `json:"is_official"`
}

// Results contains the results recieved from API.
//  > see also: https://docs.intellivoid.net/spamprotection/v1/lookup
type Results struct {
	// The private Telegram ID of the entity
	PrivateTelegramID string `json:"private_telegram_id"`

	// The entity type, can be user, bot, group, supergroup or channel
	EntityType string `json:"entity_type"`

	// The attributes of the entity, such as blacklist information and so on.
	Attributes *Attributes `json:"attributes"`

	// Information about the generalized language prediction
	LanguagePrediction *LanguagePrediction `json:"language_prediction"`

	// Information about the generalized spam prediction
	SpamPrediction *SpamPrediction `json:"spam_prediction"`

	// The Unix Timestamp of when this was last updated in the database
	LastUpdated int `json:"last_updated"`
}
