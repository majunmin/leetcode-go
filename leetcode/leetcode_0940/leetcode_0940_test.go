package leetcode_0940

import (
	"fmt"
	"testing"
)

func Test_distinctSubseqII(t *testing.T) {
	fmt.Println(distinctSubseqII("abc"))
	fmt.Println(distinctSubseqII("aba"))
	fmt.Println(distinctSubseqII("aaa"))
	// 97915677
	fmt.Println(distinctSubseqII("zchmliaqdgvwncfatcfivphddpzjkgyygueikthqzyeeiebczqbqhdytkoawkehkbizdmcnilcjjlpoeoqqoqpswtqdpvszfaksn"))
	// 584215525
	fmt.Println(distinctSubseqII("ajxjagdwzxxehvwbxhenrxtoydfobqrlugeuklytwonkrilsthwokzobvtraitboxlsazxstwnjnwnouzuzsskwteuapmmyexvdj"))
	fmt.Println(distinctSubseqII("trmnysxnqebcexbadivzlydqreqzxxnegygoddmhiywgfdlbomkauddngxrolekxdchoimomztkfdobtjwdomdrouyuvpmafqkvi"))
}
