package net

import (
	"fmt"
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
)

type Version struct {
	ObfuscationKeys []uint32                        `yaml:"obfuscation_keys"`
	Packets         map[PacketID]PacketVersionEntry `yaml:"packets"`
}

type PacketVersionEntry struct {
	Packet string `yaml:"packet"`
	Size   int    `yaml:"size"`
}

var PacketVersions map[int]*Version

func init() {
	file, err := ioutil.ReadFile("./config/packets.yml")

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	err = yaml.Unmarshal(file, &PacketVersions)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	i := 0
	keys := make([]int, len(PacketVersions))

	for k := range PacketVersions {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	for i := range keys {
		if i == 0 {
			continue
		}

		pv := PacketVersions[keys[i]]
		last := PacketVersions[keys[i-1]]

		if pv.ObfuscationKeys == nil || len(pv.ObfuscationKeys) == 0 {
			pv.ObfuscationKeys = last.ObfuscationKeys
		}

		if pv.Packets == nil {
			pv.Packets = make(map[PacketID]PacketVersionEntry)
		}

		for id, v := range last.Packets {
			_, found := pv.Packets[id]

			if !found {
				pv.Packets[id] = v
			}
		}
	}
}
