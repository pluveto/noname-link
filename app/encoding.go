package noname

import (
	"github.com/pluveto/demo1/app/params"
	"github.com/evmos/ethermint/encoding/codec"
)

// MakeEncodingConfig creates an EncodingConfig for testing
func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()	
	codec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	codec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
