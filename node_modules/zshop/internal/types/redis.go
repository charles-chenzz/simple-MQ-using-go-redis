package types

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net"
	"time"
)

// redis config

type Options func(*redis.Options)

func RedisOptions(options ...Options) *redis.Options {
	opts := &redis.Options{}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithAddr(add string) Options {
	return func(options *redis.Options) {
		options.Addr = add
	}
}

func WithNetwork(network string) Options {
	return func(options *redis.Options) {
		options.Network = network
	}
}

func WithClientName(clientname string) Options {
	return func(options *redis.Options) {
		options.ClientName = clientname
	}
}

func WithDialer(dialer func(ctx context.Context, network string, addr string) (net.Conn, error)) Options {
	return func(options *redis.Options) {
		options.Dialer = dialer
	}
}

func WithUserName(username string) Options {
	return func(options *redis.Options) {
		options.Username = username
	}
}

func WithPassword(password string) Options {
	return func(options *redis.Options) {
		options.Password = password
	}
}

func WithCredentialProvider(credentialProvider func() (username string, password string)) Options {
	return func(options *redis.Options) {
		options.CredentialsProvider = credentialProvider
	}
}

func WithDB(db int) Options {
	return func(options *redis.Options) {
		options.DB = db
	}
}

func WithMaxRetries(retry int) Options {
	return func(options *redis.Options) {
		options.MaxRetries = retry
	}
}

func WithMinRetryBackoff(backoff time.Duration) Options {
	return func(options *redis.Options) {
		options.MinRetryBackoff = backoff
	}
}

func WithMaxRetryBackoff(backoff time.Duration) Options {
	return func(options *redis.Options) {
		options.MaxRetryBackoff = backoff
	}
}

func WithDialTimeout(dialTimeout time.Duration) Options {
	return func(options *redis.Options) {
		options.DialTimeout = dialTimeout
	}
}

func WithReadTimeout(readTimeout time.Duration) Options {
	return func(options *redis.Options) {
		options.ReadTimeout = readTimeout
	}
}

func WithWriteTimeout(writeTimeout time.Duration) Options {
	return func(options *redis.Options) {
		options.WriteTimeout = writeTimeout
	}
}
