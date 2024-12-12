package bpf

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"kyanos/common"
	"os"
	"unsafe"

	"github.com/cilium/ebpf/perf"
)

func PullProcessExitEvents(ctx context.Context, channels []chan *AgentProcessExitEvent) error {
	pageSize := os.Getpagesize()
	perCPUBuffer := pageSize * 4
	eventSize := int(unsafe.Sizeof(AgentProcessExitEvent{}))
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "ProcExitEvents"), perCPUBuffer)
	if err == nil {
		go func(*perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				if evt, err := parseExitEvent(record.RawSample); err != nil {
					common.AgentLog.Errorf("[dataReader] handleKernEvt err: %s\n", err)
					continue
				} else {
					for _, ch := range channels {
						ch <- evt
					}
				}
			}
		}(reader)
	}
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseExitEvent(rawSample []byte) (*AgentProcessExitEvent, error) {
	event := AgentProcessExitEvent{}
	if err := binary.Read(bytes.NewBuffer(rawSample), binary.LittleEndian, &event); err != nil {
		return nil, fmt.Errorf("parse event: %w", err)
	}
	return &event, nil
}

func PullProcessExecEvents(ctx context.Context, channels []chan *AgentProcessExecEvent) error {
	pageSize := os.Getpagesize()
	perCPUBuffer := pageSize * 4
	eventSize := int(unsafe.Sizeof(AgentProcessExecEvent{}))
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "ProcExecEvents"), perCPUBuffer)
	if err == nil {
		go func(*perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				if evt, err := parseExecEvent(record.RawSample); err != nil {
					common.AgentLog.Errorf("[dataReader] handleKernEvt err: %s\n", err)
					continue
				} else {
					for _, ch := range channels {
						ch <- evt
					}
				}
			}
		}(reader)
	}
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseExecEvent(rawSample []byte) (*AgentProcessExecEvent, error) {
	event := AgentProcessExecEvent{}
	if err := binary.Read(bytes.NewBuffer(rawSample), binary.LittleEndian, &event); err != nil {
		return nil, fmt.Errorf("parse event: %w", err)
	}
	return &event, nil
}

func PullSyscallDataEvents(ctx context.Context, channels []chan *SyscallEventData, perfCPUBufferPageNum int, hook SyscallEventHook) error {
	pageSize := os.Getpagesize()
	perCPUBuffer := pageSize * perfCPUBufferPageNum
	eventSize := int(unsafe.Sizeof(AgentProcessExecEvent{}))
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "SyscallRb"), perCPUBuffer)
	if err == nil {
		go func(*perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				if evt, err := parseSyscallDataEvent(record.RawSample); err != nil {
					common.AgentLog.Errorf("[dataReader] handle syscall data err: %s\n", err)
					continue
				} else {
					if hook != nil {
						hook(evt)
					}
					tgidfd := evt.SyscallEvent.Ke.ConnIdS.TgidFd
					ch := channels[int(tgidfd)%len(channels)]
					ch <- evt
				}
			}
		}(reader)
	}
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseSyscallDataEvent(rawSample []byte) (*SyscallEventData, error) {
	event := new(SyscallEventData)
	err := binary.Read(bytes.NewBuffer(rawSample), binary.LittleEndian, &event.SyscallEvent)
	if err != nil {
		return nil, err
	}
	msgSize := event.SyscallEvent.BufSize
	buf := make([]byte, msgSize)
	if msgSize > 0 {
		headerSize := uint(unsafe.Sizeof(event.SyscallEvent)) - 4
		err = binary.Read(bytes.NewBuffer(rawSample[headerSize:]), binary.LittleEndian, &buf)
		if err != nil {
			return nil, err
		}
	}
	event.Buf = buf

	// tgidFd := event.SyscallEvent.Ke.ConnIdS.TgidFd
	return event, nil
}

func PullSslDataEvents(ctx context.Context, channels []chan *SslData, perfCPUBufferPageNum int, hook SslEventHook) error {
	pageSize := os.Getpagesize()
	perCPUBuffer := pageSize * perfCPUBufferPageNum
	eventSize := int(unsafe.Sizeof(SslData{}))
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "SslRb"), perCPUBuffer)
	if err == nil {
		go func(*perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				if evt, err := parseSslDataEvent(record.RawSample); err != nil {
					common.AgentLog.Errorf("[dataReader] ssl data event err: %s\n", err)
					continue
				} else {
					if hook != nil {
						hook(evt)
					}
					tgidfd := evt.SslEventHeader.Ke.ConnIdS.TgidFd
					ch := channels[int(tgidfd)%len(channels)]
					ch <- evt
				}
			}
		}(reader)
	}
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseSslDataEvent(record []byte) (*SslData, error) {
	event := new(SslData)
	err := binary.Read(bytes.NewBuffer(record), binary.LittleEndian, &event.SslEventHeader)
	if err != nil {
		return nil, err
	}
	msgSize := event.SslEventHeader.BufSize
	headerSize := uint(unsafe.Sizeof(event.SslEventHeader))
	buf := make([]byte, msgSize)
	err = binary.Read(bytes.NewBuffer(record[headerSize:]), binary.LittleEndian, &buf)
	if err != nil {
		return nil, err
	}
	event.Buf = buf

	return event, nil
}

// PullConnDataEvents 从 eBPF 映射中读取连接数据事件，并将其分发到指定的通道
// ctx: 上下文，用于控制 goroutine 的生命周期
// channels: 用于传递连接事件的通道数组
// perfCPUBufferPageNum: 每个 CPU 的缓冲区页数
// hook: 处理连接事件的钩子函数
func PullConnDataEvents(ctx context.Context, channels []chan *AgentConnEvtT, perfCPUBufferPageNum int, hook ConnEventHook) error {
	// 获取系统页面大小
	pageSize := os.Getpagesize()
	// 计算每个 CPU 的缓冲区大小
	perCPUBuffer := pageSize * perfCPUBufferPageNum
	// 获取事件的大小
	eventSize := int(unsafe.Sizeof(AgentConnEvtT{}))
	// 如果事件大小大于等于缓冲区大小，则调整缓冲区大小
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	// 创建 perf 事件读取器
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "ConnEvtRb"), perCPUBuffer)
	if err == nil {
		// 启动一个 goroutine 来读取事件
		go func(reader *perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					// 上下文取消，退出 goroutine
					return
				default:
				}
				// 读取事件记录
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						// 读取器关闭，退出 goroutine
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					// 读取错误，继续读取
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				// 解析连接事件
				if evt, err := parseConnEvent(record.RawSample); err != nil {
					// 解析错误，记录日志并继续读取
					common.AgentLog.Errorf("[dataReader] conn event err: %s\n", err)
					continue
				} else {
					// 如果有钩子函数，调用钩子函数处理事件
					if hook != nil {
						hook(evt)
					}
					// 计算事件对应的通道索引
					tgidFd := uint64(evt.ConnInfo.ConnId.Upid.Pid)<<32 | uint64(evt.ConnInfo.ConnId.Fd)
					ch := channels[int(tgidFd)%len(channels)]
					// 将事件发送到对应的通道
					ch <- evt
				}
			}
		}(reader)
	}
	// 如果创建读取器失败，记录警告日志
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseConnEvent(record []byte) (*AgentConnEvtT, error) {
	var event AgentConnEvtT
	err := binary.Read(bytes.NewBuffer(record), binary.LittleEndian, &event)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func PullKernEvents(ctx context.Context, channels []chan *AgentKernEvt, perfCPUBufferPageNum int, hook KernEventHook) error {
	pageSize := os.Getpagesize()
	perCPUBuffer := pageSize * perfCPUBufferPageNum
	eventSize := int(unsafe.Sizeof(AgentKernEvt{}))
	if eventSize >= perCPUBuffer {
		perCPUBuffer = perCPUBuffer * (1 + (eventSize / perCPUBuffer))
	}
	reader, err := perf.NewReader(GetMapFromObjs(Objs, "Rb"), perCPUBuffer)
	if err == nil {
		go func(*perf.Reader) {
			defer reader.Close()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}
				record, err := reader.Read()
				if err != nil {
					if errors.Is(err, perf.ErrClosed) {
						common.BPFLog.Debug("[dataReader] Received signal, exiting..")
						return
					}
					common.BPFLog.Debugf("[dataReader] reading from reader: %s\n", err)
					continue
				}

				if evt, err := parseKernEvent(record.RawSample); err != nil {
					common.AgentLog.Errorf("[dataReader] kern event err: %s\n", err)
					continue
				} else {
					if hook != nil {
						hook(evt)
					}
					tgidFd := evt.ConnIdS.TgidFd
					ch := channels[int(tgidFd)%len(channels)]
					ch <- evt
				}
			}
		}(reader)
	}
	if err != nil {
		common.BPFLog.Warningf("[bpf] set up perf reader failed: %s\n", err)
	}
	return err
}

func parseKernEvent(record []byte) (*AgentKernEvt, error) {
	var event AgentKernEvt
	err := binary.Read(bytes.NewBuffer(record), binary.LittleEndian, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

type SyscallEventHook func(evt *SyscallEventData)
type SslEventHook func(evt *SslData)
type ConnEventHook func(evt *AgentConnEvtT)
type KernEventHook func(evt *AgentKernEvt)
