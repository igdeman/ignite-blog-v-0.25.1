package keeper

import (
	"blog/x/blog/mongo"
	"encoding/binary"
	"log"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"blog/x/blog/types"
)

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	postCountByteKey := []byte(types.PostCountKey)
	postCountByteValue := store.Get(postCountByteKey)

	if postCountByteValue == nil {
		return 0
	}

	return binary.BigEndian.Uint64(postCountByteValue)
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostCountKey))
	postCountByteKey := []byte(types.PostCountKey)

	countByteValue := make([]byte, 8)
	binary.BigEndian.PutUint64(countByteValue, count)

	store.Set(postCountByteKey, countByteValue)
}

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PostKey))

	postByteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(postByteKey, post.Id)

	postByteValue := k.cdc.MustMarshal(&post)

	store.Set(postByteKey, postByteValue)

	k.SetPostCount(ctx, count+1)

	// Insert post into the Mongo DB
	err := mongo.AddPostMongo(post)
	if err != nil {
		log.Fatal(err)
	}

	return count
}
