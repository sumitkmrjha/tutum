package raft

import "github.com/coreos/raft"

func raft_init(){
	raft.DefaultElectionTimeout = 1
}
