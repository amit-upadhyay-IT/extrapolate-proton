package system_design

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/* NOTE BEGIN
 * Currently, I am not writing details about it, coz it will consume much of my energy :P. But I will code it now.
 * This kind of consistent hashing I am implementing is for `n` nodes and each of the node having replications factor of
 * `r`.
 * For implementation I will map the hash values in range [0, 1), this hash range will be for machines and the keys.

NOTE END */

func System_design_consistent_hashing() string {

	ch := GetConsistentHash(6, 3)
	var res strings.Builder
	for _, v := range ch.hashMapping {
		res.WriteString("[" + strconv.Itoa(v.nodeId) + "," + strconv.Itoa(v.repId) + "," + fmt.Sprint(v.hashVal) + "]\n")
	}
	res.WriteString("\n")
	key := "amit"
	machineId, repId := ch.GetMachineId(key)
	res.WriteString("sending request for key: " + key + " with hash: " + FloatToString(ModifyHashRange(GetMD5Hash(key))) + "\n")
	res.WriteString("result: machineId is " + strconv.Itoa(machineId) + " and replication id is " + strconv.Itoa(repId) + "\n\n")

	key = "upadhyay"
	machineId, repId = ch.GetMachineId(key)
	res.WriteString("sending request for key: " + key + " with hash: " + FloatToString(ModifyHashRange(GetMD5Hash(key))) + "\n")
	res.WriteString("result: machineId is " + strconv.Itoa(machineId) + " and replication id is " + strconv.Itoa(repId) + "\n\n")

	key = "amit upadhyay"
	machineId, repId = ch.GetMachineId(key)
	res.WriteString("sending request for key: " + key + " with hash: " + FloatToString(ModifyHashRange(GetMD5Hash(key))) + "\n")
	res.WriteString("result: machineId is " + strconv.Itoa(machineId) + " and replication id is " + strconv.Itoa(repId))
	return res.String()
}

// this would represent the machine number, replication number of that machine and the hash value (or the point) which
// it is representing on a circular representation
type MachineHashMapping struct {
	nodeId int
	repId int
	hashVal float64
}

// represents the structure of the consistent hash, i.e. machine mapping details and count of machines present for
// performing the hashing
// TODO: make this singleton
type ConsistentHash struct {
	nodesCount int
	replicasCount int
	hashMapping []MachineHashMapping
}

// methods for ConsistentHash

// this will return the machine id and the replication id
func (ch *ConsistentHash) GetMachineId(key string) (int, int) {
	// hash the input key
	keyHash := ModifyHashRange(GetMD5Hash(key))
	// edge case where keyHash can be greater than the hashValue of last node in our circular cycle of nodes
	// in that case we should return first node present
	if keyHash > ch.hashMapping[len(ch.hashMapping)-1].hashVal {
		return ch.hashMapping[0].nodeId, ch.hashMapping[0].repId
	}
	// lets find the next machine where request can be mapped
	nodePosition := ch.getNextGreaterIndex(keyHash)
	return ch.hashMapping[nodePosition].nodeId, ch.hashMapping[nodePosition].repId
}

func (ch *ConsistentHash) getNextGreaterIndex(hashVal float64) int {
	for i, det := range ch.hashMapping {
		if hashVal < det.hashVal {
			return i
		}
	}
	return len(ch.hashMapping)-1
}

func GetConsistentHash(nodesCnt, repCnt int) *ConsistentHash {

	// construct machine hash mapping, we will have nodesCnt * repCount items to be mapped on the circle
	// lets calculate hash for each of them
	hashMappingDetails := make([]MachineHashMapping, 0)
	for x := 0; x < nodesCnt; x++ {
		for y := 1; y <= repCnt; y++ {
			hashBytes := GetMD5Hash(strconv.Itoa(x) + "_" + strconv.Itoa(y))
			hashRep := ModifyHashRange(hashBytes)
			hashMappingDetails = append(hashMappingDetails, MachineHashMapping{
				nodeId:  x,
				repId:   y,
				hashVal: hashRep,
			})
		}
	}
	// lets sort the hash mapping details based on the hash value, coz it will help up tell the next node in optimal time
	sort.Slice(hashMappingDetails, func(i, j int) bool {
		return hashMappingDetails[i].hashVal < hashMappingDetails[j].hashVal
	})
	return &ConsistentHash{nodesCount:nodesCnt, replicasCount:repCnt, hashMapping:hashMappingDetails}
}

func ModifyHashRange(hash []byte) float64 {
	decimalPrecision := int64(10000)
	hexString := hex.EncodeToString(hash)[:8]
	n, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return (float64(n % decimalPrecision))/float64(decimalPrecision)
}

func countDigitsInFloatingNum(num float64) int {
	cnt := 0
	for num > 1 {
		num = num/10
		cnt++
	}
	return cnt
}

func GetMD5Hash(text string) []byte {
	hash := md5.Sum([]byte(text))
	return hash[:]
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}