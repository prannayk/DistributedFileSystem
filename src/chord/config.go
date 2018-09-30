package main

import (
	"errors"
	"log"
	"os"
	"time"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math/big"
	"sync"
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

func DefaultConfig () *Config {
	return &Config{
		KeySize : hashKeySizeDefault,
		IDLength : hashKeySizeDefault / 8,
		FixNextFingerInterval : 50 * time.Millisecond, 
		StabilizeInterval : 100 * time.Millisecond,
		RetryInterval : 200 * time.Millisecond,
		ConnectionTimeout : 1000 * time.Millisecond,
		DialOption : []grpc.DialOption{grpc.WithInsecure(),},
		Log : log.New(os.Stderr, "DHT : ", log.LstdFlags), 			// last part copied verbose, not set correctly
	}
}

var ERR_SET_CONFIG = errors.New("DHT : Cannot set configuration more than once")

var config struct {  			// using anon struct
	Config
	Max *big.Int
	SyncOnce 	sync.Once
}

var Log grpclog.Logger

func init () {
	InitializeSet(DefaultConfig())
}

func InitializeSet(cfg *Config){				// same functionality as setter in the following function
	if err := cfg.Validate(); err != nil {
		config.Log.Fatalf("Error setting the configuration : %v", err);
	}
	
	config.Config = *cfg
	config.Max = GetMax()
	Log = config.Log

}

func Init(cfg *Config) error {
	err := ERR_SET_CONFIG
	config.SyncOnce.Do(func() {						// functionality unclear
		if err := cfg.Validate(); err != nil {
			return
		}
	
		config.Config = *cfg
		config.Max = GetMax()
		Log = config.Log
	})
	return err
}

func GetMax() *big.Int {						// creates the value 2^m modulo which we store the keys in fingerTable
	max := big.NewInt(2)
	b2 := big.NewInt(2)
	for i := 0 ; i < config.KeySize ; i++ {
		max.Mul(max, b2)
	}
	return max
}
