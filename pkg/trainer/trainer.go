package trainer

import (
	"fmt"
	"log/slog"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

const freezeWriteInterval = 200 * time.Millisecond

type Trainer struct {
	pid            int
	storageAddress uint64

	stop, stopped chan struct{}
}

type IndexedValue struct {
	Index int
	Value int32
}

func NewTrainer(storageAddress uint64) (*Trainer, error) {
	pid, err := findProcessByName("Game.exe")
	if err != nil {
		return nil, err
	}
	slog.Info("Found game pid", slog.Int("pid", pid))
	return &Trainer{
		pid:            pid,
		storageAddress: storageAddress,
	}, nil
}

func (t *Trainer) ReadStorageValues() ([]int32, error) {
	res := make([]int32, len(StorageLocations))
	local := make([]unix.Iovec, len(res))
	remote := make([]unix.RemoteIovec, len(res))
	for i := range res {
		local[i], remote[i] = t.createIovecPair(i, &res[i])
	}
	c, err := unix.ProcessVMReadv(t.pid, local, remote, 0)
	if err != nil {
		return nil, err
	}
	slog.Info("Refreshed storage values from game memory", slog.Int("bytesCopied", c), slog.Any("result", res))
	return res, nil
}

func (t *Trainer) WriteStorageValues(values []int32) error {
	local := make([]unix.Iovec, len(values))
	remote := make([]unix.RemoteIovec, len(values))
	for i := range values {
		local[i], remote[i] = t.createIovecPair(i, &values[i])
	}

	c, err := unix.ProcessVMWritev(t.pid, local, remote, 0)
	slog.Info("Wrote storage values to game memory", slog.Int("bytesCopied", c))
	return err
}

func (t *Trainer) FreezeStorageValues(values []IndexedValue) {
	if t.stop != nil {
		t.UnfreezeStorageValues()
	}
	t.stop = make(chan struct{})
	t.stopped = make(chan struct{})

	go func() {
		slog.Info("Started freezing values")
		defer close(t.stopped)

		ticker := time.NewTicker(freezeWriteInterval)
		defer ticker.Stop()

		local := make([]unix.Iovec, len(values))
		remote := make([]unix.RemoteIovec, len(values))

		for i, value := range values {
			local[i], remote[i] = t.createIovecPair(value.Index, &value.Value)
		}

		for {
			select {
			case <-t.stop:
				slog.Info("Stopped freezing values")
				return
			case <-ticker.C:
				_, err := unix.ProcessVMWritev(t.pid, local, remote, 0)
				if err != nil {
					switch err.Error() {
					case "no such process":
						slog.Error("Game process not found, stopping freeze", slog.String("error", err.Error()))
						return
					default:
						slog.Error("Failed to write frozen values to game memory", slog.String("error", err.Error()))
					}
				}
			}
		}
	}()
}

func (t *Trainer) UnfreezeStorageValues() {
	if t.stop != nil {
		close(t.stop)
		<-t.stopped
		t.stop = nil
		t.stopped = nil
	}
}

func (t *Trainer) createIovecPair(i int, value *int32) (unix.Iovec, unix.RemoteIovec) {
	// #nosec G103 -- Intended and needed usage of unsafe for low-level memory access
	p := (*byte)(unsafe.Pointer(value))
	size := unsafe.Sizeof(*value)

	if i < 0 {
		panic(fmt.Sprintf("index cannot be negative, but was %d", i))
	}
	iUint64 := uint64(i)
	local := unix.Iovec{
		Base: p,
		Len:  uint64(size),
	}
	remote := unix.RemoteIovec{
		Base: uintptr(t.storageAddress + iUint64*uint64(size)),
		Len:  int(size),
	}
	return local, remote
}
