// Code generated by go generate internal/cmd/pairs_gen; DO NOT EDIT.
package types

import (
	"time"
)

// All available pairs.
const (
	AccessKey    = "access_key"
	Checksum     = "checksum"
	Count        = "count"
	Delimiter    = "delimiter"
	Expire       = "expire"
	Host         = "host"
	Location     = "location"
	Name         = "name"
	Port         = "port"
	Protocol     = "protocol"
	Recursive    = "recursive"
	SecretKey    = "secret_key"
	Size         = "size"
	StorageClass = "storage_class"
	Type         = "type"
	UpdatedAt    = "updated_at"
)

// WithAccessKey will apply access_key value to Options
func WithAccessKey(v string) *Pair {
	return &Pair{
		Key:   AccessKey,
		Value: v,
	}
}

// WithChecksum will apply checksum value to Options
func WithChecksum(v string) *Pair {
	return &Pair{
		Key:   Checksum,
		Value: v,
	}
}

// WithCount will apply count value to Options
func WithCount(v int64) *Pair {
	return &Pair{
		Key:   Count,
		Value: v,
	}
}

// WithDelimiter will apply delimiter value to Options
func WithDelimiter(v string) *Pair {
	return &Pair{
		Key:   Delimiter,
		Value: v,
	}
}

// WithExpire will apply expire value to Options
func WithExpire(v int) *Pair {
	return &Pair{
		Key:   Expire,
		Value: v,
	}
}

// WithHost will apply host value to Options
func WithHost(v string) *Pair {
	return &Pair{
		Key:   Host,
		Value: v,
	}
}

// WithLocation will apply location value to Options
func WithLocation(v string) *Pair {
	return &Pair{
		Key:   Location,
		Value: v,
	}
}

// WithName will apply name value to Options
func WithName(v string) *Pair {
	return &Pair{
		Key:   Name,
		Value: v,
	}
}

// WithPort will apply port value to Options
func WithPort(v int) *Pair {
	return &Pair{
		Key:   Port,
		Value: v,
	}
}

// WithProtocol will apply protocol value to Options
func WithProtocol(v string) *Pair {
	return &Pair{
		Key:   Protocol,
		Value: v,
	}
}

// WithRecursive will apply recursive value to Options
func WithRecursive(v bool) *Pair {
	return &Pair{
		Key:   Recursive,
		Value: v,
	}
}

// WithSecretKey will apply secret_key value to Options
func WithSecretKey(v string) *Pair {
	return &Pair{
		Key:   SecretKey,
		Value: v,
	}
}

// WithSize will apply size value to Options
func WithSize(v int64) *Pair {
	return &Pair{
		Key:   Size,
		Value: v,
	}
}

// WithStorageClass will apply storage_class value to Options
func WithStorageClass(v string) *Pair {
	return &Pair{
		Key:   StorageClass,
		Value: v,
	}
}

// WithType will apply type value to Options
func WithType(v string) *Pair {
	return &Pair{
		Key:   Type,
		Value: v,
	}
}

// WithUpdatedAt will apply updated_at value to Options
func WithUpdatedAt(v time.Time) *Pair {
	return &Pair{
		Key:   UpdatedAt,
		Value: v,
	}
}
