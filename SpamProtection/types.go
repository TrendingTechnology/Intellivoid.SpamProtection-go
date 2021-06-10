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

type APIResponse struct {
	Success      bool `json:"success"`
	ResponseCode int  `json:"response_code"`
	Results      []Results `json:"results"`
	Error        []Error   `json:"error"`
}

type Error struct {
	ErrorCode int    `json:"error_code"`
	Type      string `json:"type"`
	Message   string `json:"message"`
}

type SpamPrediction struct {
	HamPrediction  float64 `json:"ham_prediction"`
	SpamPrediction float64 `json:"spam_prediction"`
}

type LanguagePrediction struct {
	Language    string  `json:"language"`
	Probability float64 `json:"probability"`
}

type Attributes struct {
	IsBlacklisted               bool        `json:"is_blacklisted"`
	BlacklistFlag               interface{} `json:"blacklist_flag"`
	BlacklistReason             interface{} `json:"blacklist_reason"`
	OriginalPrivateID           interface{} `json:"original_private_id"`
	IsPotentialSpammer          bool        `json:"is_potential_spammer"`
	IsOperator                  bool        `json:"is_operator"`
	IsAgent                     bool        `json:"is_agent"`
	IsWhitelisted               bool        `json:"is_whitelisted"`
	IntellivoidAccountsVerified bool        `json:"intellivoid_accounts_verified"`
	IsOfficial                  bool        `json:"is_official"`
}

// Top Level Struct
type Results struct {
	PrivateTelegramID string `json:"private_telegram_id"`
	EntityType        string `json:"entity_type"`
	Attributes []Attributes `json:"attributes"`
	LanguagePrediction         []LanguagePrediction `json:"language_prediction"`
	SpamPrediction             []SpamPrediction `json:"spam_prediction"`
	LastUpdated int `json:"last_updated"`
}