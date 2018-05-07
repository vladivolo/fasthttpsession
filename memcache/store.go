package memcache

import (
	"github.com/phachon/fasthttpsession"
	"github.com/valyala/fasthttp"
	"github.com/bradfitz/gomemcache/memcache"
)

// session memCache store

// new default memCache store
func NewMemCacheStore(sessionId string) *Store {
	memCacheStore := &Store{}
	memCacheStore.Init(sessionId, make(map[string]interface{}))
	return memCacheStore
}

// new memCache store data
func NewMemCacheStoreData(sessionId string, data map[string]interface{}) *Store {
	memCacheStore := &Store{}
	memCacheStore.Init(sessionId, data)
	return memCacheStore
}

type Store struct {
	fasthttpsession.Store
}

// save store
func (mcs *Store) Save(ctx *fasthttp.RequestCtx) error {

	value, err := utils.GobEncode(mcs.GetAll())
	if err != nil {
		return err
	}

	return provider.memCacheClient.Set(&memcache.Item{
		Key: provider.getMemCacheSessionKey(mcs.GetSessionId()),
		Value: value,
		Expiration: int32(provider.maxLifeTime),
	})
}