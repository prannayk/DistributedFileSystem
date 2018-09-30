package main

import (
	"errors"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const hashKeySizeDefault = 64

var (
	ErrBadKeyLen = errors.New("Key length must be word aligned")
	ErrBadIDLen = errors.New("ID must key length in number of bytes")
)

type Config struct {
	KeySize 				int
	IDLength 				int
	FixNextFingerInterval 	time.Duration
	StabilizeInterval 		time.Duration
	ConnectionTimeout 		time.Duration
	RetryInterval 			time.Duration
	DialOption 				[]grpc.DialOption
	Log 					grpclog.Logger
}

func (cfg *Config) Validate() error {
	if cfg.KeySize %8 != 0 {
		return ErrBadKeyLen
	}
	if cfg.IDLength != cfg.KeySize / 8 {
		return ErrBadIDLen
	}
	return nil
}

func DefaultConfig () Config {
	return &Config{
		KeySize : hashKeySizeDefault,
		IDLength : hashKeySizeDefault / 8,
		FixNextFingerInterval : 50 * time.Millisecond, 
		StabilizeInterval : 100 * time.Millisecond,
		RetryInterval : 200 * time.Millisecond,
		ConnectionTimeout : 1000 * time.Millisecond,
		DialOption : []grpc.DialOption{grpc.WithInsecure(),},
		Log : log.New(os.stderr, "DHT : ", log.LstdFlags), 			// last part copied verbose, not set correctly
	}
}
