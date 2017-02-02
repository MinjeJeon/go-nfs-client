package nfs

import (
	"fmt"
	"os"
)

const (
	NFS3_OK             = 0
	NFS3ERR_PERM        = 1
	NFS3ERR_NOENT       = 2
	NFS3ERR_IO          = 5
	NFS3ERR_NXIO        = 6
	NFS3ERR_ACCES       = 13
	NFS3ERR_EXIST       = 17
	NFS3ERR_XDEV        = 18
	NFS3ERR_NODEV       = 19
	NFS3ERR_NOTDIR      = 20
	NFS3ERR_ISDIR       = 21
	NFS3ERR_INVAL       = 22
	NFS3ERR_FBIG        = 27
	NFS3ERR_NOSPC       = 28
	NFS3ERR_ROFS        = 30
	NFS3ERR_MLINK       = 31
	NFS3ERR_NAMETOOLONG = 63
	NFS3ERR_NOTEMPTY    = 66
	NFS3ERR_DQUOT       = 69
	NFS3ERR_STALE       = 70
	NFS3ERR_REMOTE      = 71
	NFS3ERR_BADHANDLE   = 10001
	NFS3ERR_NOT_SYNC    = 10002
	NFS3ERR_BAD_COOKIE  = 10003
	NFS3ERR_NOTSUPP     = 10004
	NFS3ERR_TOOSMALL    = 10005
	NFS3ERR_SERVERFAULT = 10006
	NFS3ERR_BADTYPE     = 10007
)

func NFS3Error(errnum uint32) error {
	switch errnum {
	case NFS3ERR_PERM:
		return os.ErrPermission
	case NFS3ERR_EXIST:
		return os.ErrExist
	case NFS3ERR_NOENT:
		return os.ErrNotExist
	default:
		return &Error{fmt.Sprintf("error: %d", errnum)}
	}

	return nil
}

// Error represents an unexpected I/O behavior.
type Error struct {
	ErrorString string
}

func (err *Error) Error() string { return err.ErrorString }
