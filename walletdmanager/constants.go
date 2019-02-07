// Copyright (c) 2018, The WeaknessCoin Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in WEAK
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "weakness-service-session.log"
	logWalletdAllSessionsFilename        = "weakness-service.log"
	logWeaknessCoindCurrentSessionFilename = "WeaknessCoind-session.log"
	logWeaknessCoindAllSessionsFilename    = "WeaknessCoind.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "weakness-service"
	weaknesscoindCommandName               = "WeaknessCoind"
)
