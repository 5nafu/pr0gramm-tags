package main

import (
	"github.com/cznic/sortutil"
	"gopkg.in/cheggaaa/pb.v1"
)

type StoreBuilder struct {
	ShowProgress bool
	byteStore    ByteStore
	iterStore    *iterStore
}

func NewStoreBuilder() *StoreBuilder {
	byteStore := NewCppStore()
	return &StoreBuilder{
		byteStore: byteStore,
		iterStore: &iterStore{byteStore},
	}
}

func (sb *StoreBuilder) Push(word string, itemId int32) {
	hash := HashWord(word)

	// check if the hash is alredy known before adding it
	known := sb.byteStore.Contains(hash)

	sb.iterStore.PushInt(hash, int32(itemId))

	if !known {
		// add item virtual tag "__all"
		sb.iterStore.PushInt(0, int32(itemId))
	}
}

func (sb *StoreBuilder) Build() IterStore {
	var bar *pb.ProgressBar
	if sb.ShowProgress {
		bar = pb.StartNew(int(sb.iterStore.KeyCount()))
		bar.ShowFinalTime = true
		defer bar.Finish()
	}

	optimizedStore := NewIterStore(sb.byteStore)
	for _, key := range sb.iterStore.Keys() {
		if bar != nil {
			bar.Increment()
		}

		// get the list of items
		items := IteratorToList(nil, sb.iterStore.GetIterator(key))
		n := sortutil.Dedupe(sortutil.Int32Slice(items))

		// and optimize it
		optimizedStore.Replace(key, items[:n])
	}

	return optimizedStore
}
