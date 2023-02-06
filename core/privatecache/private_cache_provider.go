package privatecache

import (
	"github.com/daefrom/go-dae/common"
	"github.com/daefrom/go-dae/core/state"
	"github.com/daefrom/go-dae/ethdb"
	"github.com/daefrom/go-dae/log"
	"github.com/daefrom/go-dae/trie"
)

type Provider interface {
	GetCache() state.Database
	GetCacheWithConfig() state.Database
	Commit(db state.Database, hash common.Hash) error
	Reference(child, parent common.Hash)
}

func NewPrivateCacheProvider(db ethdb.Database, config *trie.Config, cache state.Database, privateCacheEnabled bool) Provider {
	if privateCacheEnabled {
		log.Info("Using UnifiedCacheProvider.")
		return &unifiedCacheProvider{
			cache: cache,
		}
	}
	log.Info("Using SegregatedCacheProvider.")
	return &segregatedCacheProvider{db: db, config: config}
}
