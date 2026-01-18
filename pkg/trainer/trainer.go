package trainer

import (
	"log/slog"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Trainer struct {
	pid int
}

func NewTrainer() (*Trainer, error) {
	pid, err := findProcessByName("Game.exe")
	if err != nil {
		return nil, err
	}
	slog.Info("Found game pid", slog.Int("pid", pid))
	return &Trainer{pid: pid}, nil
}

func (t *Trainer) ReadStorageValues(addr uint64) ([]int32, error) {
	res := make([]int32, len(StorageLocations))
	local := make([]unix.Iovec, len(res))
	remote := make([]unix.RemoteIovec, len(res))
	for i := range res {
		p := (*byte)(unsafe.Pointer(&res[i]))
		size := int(unsafe.Sizeof(res[i]))
		local[i] = unix.Iovec{
			Base: p,
			Len:  uint64(size),
		}
		remote[i] = unix.RemoteIovec{
			Base: uintptr(addr + uint64(i*size)),
			Len:  size,
		}
	}
	c, err := unix.ProcessVMReadv(t.pid, local, remote, 0)
	if err != nil {
		return nil, err
	}
	slog.Info("Result", slog.Int("bytesCopied", c), slog.Any("result", res))
	return res, nil
}
