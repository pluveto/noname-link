package types

const (
	// CoinType
	CoinType = 60

	// FullFundraiserPath is the parts of the BIP44 HD path
	// BIP44 is a wallet standard used to determine the derivation path for 
	// addresses and private keys in cryptocurrency wallets.
	//  HD (Hierarchical Deterministic) signifies that
	//  this wallet utilizes a hierarchical deterministic
	//  approach to generate addresses and private keys.
	// 
	// Here,
	//  m denotes the master path, 
	// 44' is the coin index, 
	// 60' is the derived coin index, 
	// 0' is the account index,
	// 0 represents the external/internal chain index, 
	// and the final 0 is the address index.
	FullFundraiserPath = "m/44'/60'/0'/0/0"
)
