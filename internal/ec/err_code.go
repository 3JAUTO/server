package ec

type ErrCode = uint16

const (
	ErrCodeSuccess    = iota
	ErrCodeParameter  = iota
	ErrCodeCredential = iota
	ErrCodeAuth       = iota
	ErrCodePrivilege  = iota
)
