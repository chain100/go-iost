package merkletree

import (
	"testing"
	"log"
	//"bytes"
	. "github.com/smartystreets/goconvey/convey"
	//"github.com/golang/protobuf/proto"
	"github.com/iost-official/Go-IOS-Protocol/core/new_tx"
	"encoding/hex"
	"fmt"
)

func TestTXRMerkleTree(t *testing.T) {
	Convey("Test of Build", t, func() {
		m := TXRMerkleTree{}
		txrs := []tx.TxReceipt{
			{TxHash:[]byte("node1")},
			{TxHash:[]byte("node2")},
			{TxHash:[]byte("node3")},
			{TxHash:[]byte("node4")},
			{TxHash:[]byte("node5")},
		}
		err := m.Build(txrs)
		if err != nil {
			log.Panic(err)
		}
		So(hex.EncodeToString(m.TX2TXR["node1"]), ShouldEqual, "0a056e6f6465311a00")
		txr, err := m.GetTXR([]byte("node1"))
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("\n%s",hex.EncodeToString(txr.Encode()))


		//So(hex.EncodeToString(m.HashList[1]), ShouldEqual, "e5e1a9ed8c02ed449057a4c17618127fa8e0a1e1c19fa15a371810371ac7530b")
		//So(hex.EncodeToString(m.HashList[2]), ShouldEqual, "de333248f6058db0367c9dc3e4731ea37324d4bfbbeee22ffd3d5a4e0c28330a")
		//rootHash, err := m.RootHash()
		//So(hex.EncodeToString(rootHash), ShouldEqual, "0f8a9f1e9450978a41ff06e77df3de64866b55261ed20651c90eb6cb462b1409")
		//mp, err := m.MerklePath([]byte("node5"))
		//if err != nil {
		//	log.Panic(err)
		//}
		//So(hex.EncodeToString(mp[0]), ShouldEqual, "6e6f646535")
		//So(hex.EncodeToString(mp[1]), ShouldEqual, "946f804875563d1f73fb27b1fc8af795850e9128281954028e145108db4a0ab9")
		//So(hex.EncodeToString(mp[2]), ShouldEqual, "e5e1a9ed8c02ed449057a4c17618127fa8e0a1e1c19fa15a371810371ac7530b")
		//success, err := m.MerkleProve([]byte("node5"), rootHash, mp)
		//So(success, ShouldBeTrue)
		//b, err := proto.Marshal(&m)
		//if err != nil {
		//	log.Panic(err)
		//}
		//var m_read MerkleTree
		//err = proto.Unmarshal(b, &m_read)
		//if err != nil {
		//	log.Panic(err)
		//}
		//for i := 0; i < 16; i++ {
		//	So(bytes.Equal(m.HashList[i], m_read.HashList[i]), ShouldBeTrue)
		//}
	})
}