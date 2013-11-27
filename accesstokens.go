package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// This file defines helper functions for working with access tokens
var accesstokens = make(map[string]*AccessToken)

type Entity struct {
	id int
}

type AccessToken struct {
	id      int
	hash    string
	created time.Time
}

func makeAccessToken(e Entity) *AccessToken {
	salt := uint64(rand.New(rand.NewSource(time.Now().Unix())).Int63())
	hasher := sha1.New()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, salt+uint64(e.id))
	hasher.Write(b)
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	prefix := ""
	if (reflect.TypeOf(e) == reflect.TypeOf(&Player{})) {
		prefix = "p"
	} else if (reflect.TypeOf(e) == reflect.TypeOf(&User{})) {
		prefix = "u"
	} else {
		fmt.Println("Unknown type %T", reflect.ValueOf(e).Kind())
	}
	at_hash := prefix + hash
	at := AccessToken{e.id, at_hash, time.Now()}
	fmt.Println("AccessToken: " + at_hash)
	accesstokens[at_hash] = &at
	return &at
}
