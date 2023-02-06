package privatecache

import (
	"github.com/daefrom/go-dae/common"
	"github.com/daefrom/go-dae/core/state"
	"github.com/daefrom/go-dae/ethdb"
	"github.com/daefrom/go-dae/trie"
)

type segregatedCacheProvider struct {
	db     ethdb.Database
	config *trie.Config
}

func (p *segregatedCacheProvider) GetCache() state.Database {
	return state.NewDatabase(p.db)
}

func (p *segregatedCacheProvider) GetCacheWithConfig() state.Database {
	return state.NewDatabaseWithConfig(p.db, p.config)
}

func (p *segregatedCacheProvider) Commit(db state.Database, hash common.Hash) error {
	return db.TrieDB().Commit(hash, false, nil)
}
func (p *segregatedCacheProvider) Reference(child, parent common.Hash) {
	// do nothing
}
