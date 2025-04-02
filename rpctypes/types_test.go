package rpctypes

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
)

func TestEthSendBundleArgsValidate(t *testing.T) {
	// from https://github.com/flashbots/rbuilder/blob/develop/crates/rbuilder/src/primitives/serialize.rs#L607
	inputs := []struct {
		Payload           json.RawMessage
		ExpectedHash      string
		ExpectedUUID      string
		ExpectedUniqueKey string
	}{
		{
			Payload: []byte(`{
				"blockNumber": "0x1136F1F",
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": ["0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"]
        	}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "d9a3ae52-79a2-5ce9-a687-e2aa4183d5c6",
		},
		{
			Payload: []byte(`{
				"blockNumber": "0x1136F1F",
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": ["0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"]
			}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "d9a3ae52-79a2-5ce9-a687-e2aa4183d5c6",
		},
		{
			Payload: []byte(`{
				"blockNumber": "0xA136F1F",
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": []
			}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "5d5bf52c-ac3f-57eb-a3e9-fc01b18ca516",
		},
		{
			Payload: []byte(`{
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": []
			}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "e9ced844-16d5-5884-8507-db9338950c5c",
		},
		{
			Payload: []byte(`{
		                "blockNumber": "0x0",
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": []
			}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "e9ced844-16d5-5884-8507-db9338950c5c",
		},
		{
			Payload: []byte(`{
		                "blockNumber": null,
				"txs": ["0x02f9037b018203cd8405f5e1008503692da370830388ba943fc91a3afd70395cd496c647d5a6cc9d4b2b7fad8780e531581b77c4b903043593564c000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000064f390d300000000000000000000000000000000000000000000000000000000000000030b090c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001e0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000009184e72a0000000000000000000000000000000000000000000000000000080e531581b77c400000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000b5ea574dd8f2b735424dfc8c4e16760fc44a931b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000c001a0a9ea84ad107d335afd5e5d2ddcc576f183be37386a9ac6c9d4469d0329c22e87a06a51ea5a0809f43bf72d0156f1db956da3a9f3da24b590b7eed01128ff84a2c1"],
				"revertingTxHashes": []
			}`),
			ExpectedHash: "0xcf3c567aede099e5455207ed81c4884f72a4c0c24ddca331163a335525cd22cc",
			ExpectedUUID: "e9ced844-16d5-5884-8507-db9338950c5c",
		},
		// different empty bundles have the same uuid, they must have different unique key
		{
			Payload: []byte(`{
		                 "replacementUuid": "e9ced844-16d5-5884-8507-db9338950c5c"
			}`),
			ExpectedHash:      "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
			ExpectedUUID:      "35718fe4-5d24-51c8-93bf-9c961d7c3ea3",
			ExpectedUniqueKey: "1655edd0-29a6-5372-a19b-1ddedda14b20",
		},
		{
			Payload: []byte(`{
		                 "replacementUuid": "35718fe4-5d24-51c8-93bf-9c961d7c3ea3"
			}`),
			ExpectedHash:      "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
			ExpectedUUID:      "35718fe4-5d24-51c8-93bf-9c961d7c3ea3",
			ExpectedUniqueKey: "3c718cb9-3f6c-5dc0-9d99-264dafc0b4e9",
		},
		{
			Payload: []byte(`  {
            "version": "v2",
            "txs": [
                "0x02f86b83aa36a780800982520894f24a01ae29dec4629dfb4170647c4ed4efc392cd861ca62a4c95b880c080a07d37bb5a4da153a6fbe24cf1f346ef35748003d1d0fc59cf6c17fb22d49e42cea02c231ac233220b494b1ad501c440c8b1a34535cdb8ca633992d6f35b14428672"
            ],
            "blockNumber": "0x0",
            "minTimestamp": 123,
            "maxTimestamp": 1234,
            "revertingTxHashes": ["0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"],
            "droppingTxHashes": ["0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"],
            "refundPercent": 1,
            "refundRecipient": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5",
            "refundTxHashes": ["0x75662ab9cb6d1be7334723db5587435616352c7e581a52867959ac24006ac1fe"]
        }`),
			ExpectedHash:      "0xee3996920364173b0990f92cf6fbeb8a4ab832fe5549c1b728ac44aee0160f02",
			ExpectedUUID:      "e2bdb8cd-9473-5a1b-b425-57fa7ecfe2c1",
			ExpectedUniqueKey: "a54c1e8f-936f-5868-bded-f5138c60b34a",
		},
		{
			Payload: []byte(`  {
            "version": "v2",
            "txs": [
                "0x02f90408018303f1d4808483ab318e8304485c94a69babef1ca67a37ffaf7a485dfff3382056e78c8302be00b9014478e111f60000000000000000000000007f0f35bbf44c8343d14260372c469b331491567b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c4f4ff52950000000000000000000000000000000000000000000000000be75df44ebec5390000000000000000000000000000000000000000000036404c073ad050000000000000000000000000000000000000000000003e91fd871e8a6021ca93d911920000000000000000000000000000000000000000000000000000e91615b961030000000000000000000000000000000000000000000000000000000067eaa0b7ff8000000000000000000000000000000000000000000000000000000001229300000000000000000000000000000000000000000000000000000000f90253f9018394919fa96e88d67499339577fa202345436bcdaf79f9016ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513782a000000000000000000000000000000000000000000000000000000000000000a1a0bfd358e93f18da3ed276c3afdbdba00b8f0b6008a03476a6a86bd6320ee6938ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513785a00000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000a0a00000000000000000000000000000000000000000000000000000000000000002a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513783a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513784a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000004f85994c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2f842a060802b93a9ac49b8c74d6ade12cf6235f4ac8c52c84fd39da757d0b2d720d76fa075245230289a9f0bf73a6c59aef6651b98b3833a62a3c0bd9ab6b0dec8ed4d8fd6947f0f35bbf44c8343d14260372c469b331491567bc0f85994d533a949740bb3306d119cc777fa900ba034cd52f842a07a7ff188ddb962db42160fb3fb573f4af0ebe1a1d6b701f1f1464b5ea43f7638a03d4653d86fe510221a71cfd2b1168b2e9af3e71339c63be5f905dabce97ee61f01a0c9d68ec80949077b6c28d45a6bf92727bc49d705d201bff8c62956201f5d3a81a036b7b953d7385d8fab8834722b7c66eea4a02a66434fc4f38ebfe8f5218a87b0"
            ],
            "blockNumber": "0x0",
            "minTimestamp": 123,
            "maxTimestamp": 1234,
            "refundPercent": 20,
            "refundRecipient": "0xFF82BF5238637B7E5E345888BaB9cd99F5Ebe331",
            "refundTxHashes": ["0xffd9f02004350c16b312fd14ccc828f587c3c49ad3e9293391a398cc98c1a373"]
        }`),
			ExpectedHash:      "0x90551b655a8a5b424064e802c0ec2daae864d8b786a788c2c6f9d7902feb42d2",
			ExpectedUUID:      "e785c7c0-8bfa-508e-9c3f-cb24f1638de3",
			ExpectedUniqueKey: "fb7bff94-6f0d-5030-ab69-33adf953b742",
		},
		{
			Payload: []byte(`  {
            "version": "v2",
            "txs": [
                "0x02f90408018303f1d4808483ab318e8304485c94a69babef1ca67a37ffaf7a485dfff3382056e78c8302be00b9014478e111f60000000000000000000000007f0f35bbf44c8343d14260372c469b331491567b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c4f4ff52950000000000000000000000000000000000000000000000000be75df44ebec5390000000000000000000000000000000000000000000036404c073ad050000000000000000000000000000000000000000000003e91fd871e8a6021ca93d911920000000000000000000000000000000000000000000000000000e91615b961030000000000000000000000000000000000000000000000000000000067eaa0b7ff8000000000000000000000000000000000000000000000000000000001229300000000000000000000000000000000000000000000000000000000f90253f9018394919fa96e88d67499339577fa202345436bcdaf79f9016ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513782a000000000000000000000000000000000000000000000000000000000000000a1a0bfd358e93f18da3ed276c3afdbdba00b8f0b6008a03476a6a86bd6320ee6938ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513785a00000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000a0a00000000000000000000000000000000000000000000000000000000000000002a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513783a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513784a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000004f85994c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2f842a060802b93a9ac49b8c74d6ade12cf6235f4ac8c52c84fd39da757d0b2d720d76fa075245230289a9f0bf73a6c59aef6651b98b3833a62a3c0bd9ab6b0dec8ed4d8fd6947f0f35bbf44c8343d14260372c469b331491567bc0f85994d533a949740bb3306d119cc777fa900ba034cd52f842a07a7ff188ddb962db42160fb3fb573f4af0ebe1a1d6b701f1f1464b5ea43f7638a03d4653d86fe510221a71cfd2b1168b2e9af3e71339c63be5f905dabce97ee61f01a0c9d68ec80949077b6c28d45a6bf92727bc49d705d201bff8c62956201f5d3a81a036b7b953d7385d8fab8834722b7c66eea4a02a66434fc4f38ebfe8f5218a87b0"
            ],
            "blockNumber": "0x0",
            "minTimestamp": 123,
            "maxTimestamp": 1234,
            "refundPercent": 20,
            "refundRecipient": "0xFF82BF5238637B7E5E345888BaB9cd99F5Ebe331"
        }`),
			ExpectedHash:      "0x90551b655a8a5b424064e802c0ec2daae864d8b786a788c2c6f9d7902feb42d2",
			ExpectedUUID:      "e785c7c0-8bfa-508e-9c3f-cb24f1638de3",
			ExpectedUniqueKey: "",
		},
		{
			Payload: []byte(`  {
            "version": "v2",
            "txs": [
                "0x02f90408018303f1d4808483ab318e8304485c94a69babef1ca67a37ffaf7a485dfff3382056e78c8302be00b9014478e111f60000000000000000000000007f0f35bbf44c8343d14260372c469b331491567b000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c4f4ff52950000000000000000000000000000000000000000000000000be75df44ebec5390000000000000000000000000000000000000000000036404c073ad050000000000000000000000000000000000000000000003e91fd871e8a6021ca93d911920000000000000000000000000000000000000000000000000000e91615b961030000000000000000000000000000000000000000000000000000000067eaa0b7ff8000000000000000000000000000000000000000000000000000000001229300000000000000000000000000000000000000000000000000000000f90253f9018394919fa96e88d67499339577fa202345436bcdaf79f9016ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513782a000000000000000000000000000000000000000000000000000000000000000a1a0bfd358e93f18da3ed276c3afdbdba00b8f0b6008a03476a6a86bd6320ee6938ba0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513785a00000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000000a0a00000000000000000000000000000000000000000000000000000000000000002a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513783a0bf1c200cc6dee22da7e010c51ff8e5210da52f1c78d2171dbb5d4f739e513784a00000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000004f85994c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2f842a060802b93a9ac49b8c74d6ade12cf6235f4ac8c52c84fd39da757d0b2d720d76fa075245230289a9f0bf73a6c59aef6651b98b3833a62a3c0bd9ab6b0dec8ed4d8fd6947f0f35bbf44c8343d14260372c469b331491567bc0f85994d533a949740bb3306d119cc777fa900ba034cd52f842a07a7ff188ddb962db42160fb3fb573f4af0ebe1a1d6b701f1f1464b5ea43f7638a03d4653d86fe510221a71cfd2b1168b2e9af3e71339c63be5f905dabce97ee61f01a0c9d68ec80949077b6c28d45a6bf92727bc49d705d201bff8c62956201f5d3a81a036b7b953d7385d8fab8834722b7c66eea4a02a66434fc4f38ebfe8f5218a87b0"
            ],
            "blockNumber": "0x0",
            "minTimestamp": 123,
            "maxTimestamp": 1234,
            "refundPercent": 20
        }`),
			ExpectedHash:      "0x90551b655a8a5b424064e802c0ec2daae864d8b786a788c2c6f9d7902feb42d2",
			ExpectedUUID:      "e785c7c0-8bfa-508e-9c3f-cb24f1638de3",
			ExpectedUniqueKey: "",
		},
		{
			Payload: []byte(`{
            "version": "v2",
            "txs": [
                "0x02f86b83aa36a780800982520894f24a01ae29dec4629dfb4170647c4ed4efc392cd861ca62a4c95b880c080a07d37bb5a4da153a6fbe24cf1f346ef35748003d1d0fc59cf6c17fb22d49e42cea02c231ac233220b494b1ad501c440c8b1a34535cdb8ca633992d6f35b14428672"
            ],
            "blockNumber": "0x0",
            "revertingTxHashes": []
        }`),
			ExpectedHash:      "0xee3996920364173b0990f92cf6fbeb8a4ab832fe5549c1b728ac44aee0160f02",
			ExpectedUUID:      "22dc6bf0-9a12-5a76-9bbd-98ab77423415",
			ExpectedUniqueKey: "",
		},
	}

	for i, input := range inputs {
		t.Run(fmt.Sprintf("inout-%d", i), func(t *testing.T) {
			bundle := &EthSendBundleArgs{}
			require.NoError(t, json.Unmarshal(input.Payload, bundle))
			hash, uuid, err := bundle.Validate()
			uniqueKey := bundle.UniqueKey()
			require.NoError(t, err)
			require.Equal(t, input.ExpectedHash, hash.Hex())
			require.Equal(t, input.ExpectedUUID, uuid.String())
			if input.ExpectedUniqueKey != "" {
				require.Equal(t, input.ExpectedUniqueKey, uniqueKey.String())
			}
		})
	}
}

func TestMevSendBundleArgsValidate(t *testing.T) {
	// From: https://github.com/flashbots/rbuilder/blob/91f7a2c22eaeaf6c44e28c0bda98a2a0d566a6cb/crates/rbuilder/src/primitives/serialize.rs#L700
	// NOTE: I had to dump the hash in a debugger to get the expected hash since the test above uses a computed hash
	raw := []byte(`{
		"version": "v0.1",
		"inclusion": {
		  "block": "0x1"
		},
		"body": [
		  {
			"bundle": {
			  "version": "v0.1",
			  "inclusion": {
				"block": "0x1"
			  },
			  "body": [
				{
				  "tx": "0x02f86b0180843b9aca00852ecc889a0082520894c87037874aed04e51c29f582394217a0a2b89d808080c080a0a463985c616dd8ee17d7ef9112af4e6e06a27b071525b42182fe7b0b5c8b4925a00af5ca177ffef2ff28449292505d41be578bebb77110dfc09361d2fb56998260",
				  "canRevert": true
				},
				{
				  "tx": "0x02f8730180843b9aca00852ecc889a008288b894c10000000000000000000000000000000000000088016345785d8a000080c001a07c8890151fed9a826f241d5a37c84062ebc55ca7f5caef4683dcda6ac99dbffba069108de72e4051a764f69c51a6b718afeff4299107963a5d84d5207b2d6932a4"
				}
			  ],
			  "validity": {
				"refund": [
				  {
					"bodyIdx": 0,
					"percent": 90
				  }
				],
				"refundConfig": [
				  {
					"address": "0x3e7dfb3e26a16e3dbf6dfeeff8a5ae7a04f73aad",
					"percent": 100
				  }
				]
			  }
			}
		  },
		  {
			"tx": "0x02f8730101843b9aca00852ecc889a008288b894c10000000000000000000000000000000000000088016345785d8a000080c001a0650c394d77981e46be3d8cf766ecc435ec3706375baed06eb9bef21f9da2828da064965fdf88b91575cd74f20301649c9d011b234cefb6c1761cc5dd579e4750b1"
		  }
		],
		"validity": {
		  "refund": [
			{
			  "bodyIdx": 0,
			  "percent": 80
			}
		  ]
		},
		"metadata": {
			"signer": "0x4696595f68034b47BbEc82dB62852B49a8EE7105"
		}
	}`)

	bundle := &MevSendBundleArgs{}
	require.NoError(t, json.Unmarshal(raw, bundle))
	hash, err := bundle.Validate()
	require.NoError(t, err)
	require.Equal(t, "0x3b1994ad123d089f978074cfa197811b644e43b2b44b4c4710614f3a30ee0744", hash.Hex())
}

func TestEthsendRawTransactionArgsJSON(t *testing.T) {
	data := hexutil.MustDecode("0x1234")

	rawTransaction := EthSendRawTransactionArgs(data)

	out, err := json.Marshal(rawTransaction)
	require.NoError(t, err)

	require.Equal(t, `"0x1234"`, string(out))

	var roundtripRawTransaction EthSendRawTransactionArgs
	err = json.Unmarshal(out, &roundtripRawTransaction)
	require.NoError(t, err)
	require.Equal(t, rawTransaction, roundtripRawTransaction)
}
