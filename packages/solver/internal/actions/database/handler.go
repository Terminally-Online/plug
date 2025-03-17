package database

import (
	"solver/bindings/plug_database"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	SetUint256  = "set_uint256"
	GetUint256  = "get_uint256"
	SetInt256   = "set_int256"
	GetInt256   = "get_int256"
	SetBytes32  = "set_bytes32"
	GetBytes32  = "get_bytes32"
	SetBytes    = "set_bytes"
	GetBytes    = "get_bytes"
	SetAddress  = "set_address"
	GetAddress  = "get_address"
	SetBool     = "set_bool"
	GetBool     = "get_bool"
	SetString   = "set_string"
	GetString   = "get_string"
	RemoveValue = "remove_value"
	BatchSet    = "batch_set"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Database",
			Icon:   "https://cdn.onplug.io/protocols/database.png",
			Tags:   []string{"database", "storage", "state"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				SetUint256: {
					Sentence:     "Set uint256 {0<name:string>} to {1<value:uint256>}",
					Handler:      HandleSetUint256,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setUint",
				},
				GetUint256: {
					Sentence:     "Get uint256 value of {0<name:string>}",
					Handler:      HandleGetUint256,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getUint",
				},
				SetInt256: {
					Sentence:     "Set int256 {0<name:string>} to {1<value:int256>}",
					Handler:      HandleSetInt256,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setInt",
				},
				GetInt256: {
					Sentence:     "Get int256 value of {0<name:string>}",
					Handler:      HandleGetInt256,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getInt",
				},
				SetBytes32: {
					Sentence:     "Set bytes32 {0<name:string>} to {1<value:bytes32>}",
					Handler:      HandleSetBytes32,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "set",
				},
				GetBytes32: {
					Sentence:     "Get bytes32 value of {0<name:string>}",
					Handler:      HandleGetBytes32,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "get",
				},
				SetBytes: {
					Sentence:     "Set bytes {0<name:string>} to {1<value:bytes>}",
					Handler:      HandleSetBytes,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setBytes",
				},
				GetBytes: {
					Sentence:     "Get bytes value of {0<name:string>}",
					Handler:      HandleGetBytes,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getBytes",
				},
				SetAddress: {
					Sentence:     "Set address {0<name:string>} to {1<value:address>}",
					Handler:      HandleSetAddress,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setAddress",
				},
				GetAddress: {
					Sentence:     "Get address value of {0<name:string>}",
					Handler:      HandleGetAddress,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getAddress",
				},
				SetBool: {
					Sentence:     "Set boolean {0<name:string>} to {1<value:bool>}",
					Handler:      HandleSetBool,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setBool",
				},
				GetBool: {
					Sentence:     "Get boolean value of {0<name:string>}",
					Handler:      HandleGetBool,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getBool",
				},
				SetString: {
					Sentence:     "Set string {0<name:string>} to {1<value:string>}",
					Handler:      HandleSetString,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "setString",
				},
				GetString: {
					Sentence:     "Get string value of {0<name:string>}",
					Handler:      HandleGetString,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "getString",
				},
				RemoveValue: {
					Sentence:     "Remove value for {0<name:string>}",
					Handler:      HandleRemoveValue,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "remove",
				},
				BatchSet: {
					Sentence:     "Batch set {0<count:uint256>} bytes32 values",
					Handler:      HandleBatchSet,
					Metadata:     plug_database.PlugDatabaseMetaData,
					FunctionName: "batchSet",
				},
			},
		},
	)
}
