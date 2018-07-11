// Copyright 2018 The go-hpb Authors
// This file is part of the go-hpb.
//
// The go-hpb is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hpb is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hpb. If not, see <http://www.gnu.org/licenses/>.


package voting

import (
	"math"
	"strconv"
	"math/rand"
    "fmt"
    "reflect"
    "errors"
    
	"github.com/hpb-project/go-hpb/common"
	//"github.com/hpb-project/go-hpb/consensus"
    "math/big"
	"github.com/hpb-project/go-hpb/consensus/snapshots"
	//"github.com/hpb-project/go-hpb/blockchain/storage"
	"github.com/hpb-project/go-hpb/network/p2p"
)


// 从网络中获取最优化的
func GetBestCadNodeFromNetwork(snap *snapshots.HpbNodeSnap) (*snapshots.CadWinner, error) {
		//str := strconv.FormatUint(number, 10)
		// 模拟从外部获取		
		type CadWinners []*snapshots.CadWinner
		cadWinners := []*snapshots.CadWinner{} 
		
		hpbAddresses := snap.GetHpbNodes()
		
		//cadNodeMap,_ := GetHpbNodeSnap(db,chain,number, hash)
		
		// 模拟从peer中获取
		// all nodes = Candidate node + HPB node
		peers := p2p.PeerMgrInst().PeersAll()
		
		fmt.Println("length:", len(peers))
		
		for _, peer := range peers {
			fmt.Println("this is test:", peer.TxsRate())
		}
		
		for i := 0; i < 1000; i++ {
			//加权算法
			networkBandwidth := float64(rand.Intn(1000)) * float64(0.3)
			transactionNum := float64(rand.Intn(1000)) * float64(0.7)
			VoteIndex := networkBandwidth + transactionNum
			
			strnum := strconv.Itoa(i)
			//cadNodeMap[uint64(VoteIndex)] = &snapshots.CadWinner{"192.168.2"+strnum,"0xd3b686a79f4da9a415c34ef95926719bb8dfcaf"+strnum,uint64(VoteIndex)}
			
			//在候选列表中获取，如果候选列表中含有，在进行加入
			//if cad,exists := cadNodeMap[string(i)]; exists == true{
			bigaddr, _ := new(big.Int).SetString("d3b686a79f4da9a415c34ef95926719bb8dfcafd", 16)
			bigaddr1, _ := new(big.Int).SetString("3ee4f38f985b4c1b658dadcac9f8fb946b4b0708", 16)

		    address := common.BigToAddress(bigaddr)
		    address1 := common.BigToAddress(bigaddr1)

			//fmt.Println("this is test:", i)
			//}
			
			if ok, err := Contain(address, hpbAddresses); !ok && err == nil {
				cadWinners = append(cadWinners,&snapshots.CadWinner{"192.168.2"+strnum,address1,uint64(VoteIndex)})
			}
			
		}
		
		
		// 先获取长度，然后进行随机获取
		lnlen := int(math.Log2(float64(len(cadWinners))))
		
		var lastCadWinners []*snapshots.CadWinner
		
		for i := 0 ; i < lnlen; i++{
			lastCadWinners = append(lastCadWinners,cadWinners[rand.Intn(len(cadWinners)-1)])
		}
		
		//开始进行排序获取最大值
		bigaddr, _ := new(big.Int).SetString("d3b686a79f4da9a415c34ef95926719bb8dfcafd", 16)
		address := common.BigToAddress(bigaddr)
		
		lastCadWinnerToChain := &snapshots.CadWinner{"192.168.2.33",address,uint64(0)}
		voteIndexTemp := uint64(0)
		
		
		for _, lastCadWinner := range lastCadWinners {
	        if(lastCadWinner.VoteIndex > voteIndexTemp){
	        	  voteIndexTemp = lastCadWinner.VoteIndex
	        	  lastCadWinnerToChain = lastCadWinner //返回最优的
	        }
	    }
		
		//fmt.Println("len:", voteIndexTemp)
		fmt.Println("len:", lastCadWinnerToChain.VoteIndex)
		return lastCadWinnerToChain,nil
}

func Contain(obj interface{}, target interface{}) (bool, error) {
    targetValue := reflect.ValueOf(target)
    switch reflect.TypeOf(target).Kind() {
    case reflect.Slice, reflect.Array:
        for i := 0; i < targetValue.Len(); i++ {
            if targetValue.Index(i).Interface() == obj {
                return true, nil
            }
        }
    case reflect.Map:
        if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
            return true, nil
        }
    }

    return false, errors.New("not in array")
}

/*
func GetCadNodeMap(db hpbdb.Database,chain consensus.ChainReader, number uint64, hash common.Hash) (map[string]*snapshots.CadWinner, error) {

	cadWinnerms := make(map[string]*snapshots.CadWinner)

	if cadNodeSnapformap, err  := GetCadNodeSnap(db, chain, number, hash); err == nil{
		for _, cws := range cadNodeSnapformap.CadWinners {
		    cadWinnerms[cws.NetworkId] = &snapshots.CadWinner{cws.NetworkId,cws.Address,cws.VoteIndex}
		}
	}

    return cadWinnerms,nil
}
*/


