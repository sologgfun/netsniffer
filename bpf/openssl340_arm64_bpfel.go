// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package bpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// LoadOpenssl340 returns the embedded CollectionSpec for Openssl340.
func LoadOpenssl340() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Openssl340Bytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Openssl340: %w", err)
	}

	return spec, err
}

// LoadOpenssl340Objects loads Openssl340 and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*Openssl340Objects
//	*Openssl340Programs
//	*Openssl340Maps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadOpenssl340Objects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadOpenssl340()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// Openssl340Specs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl340Specs struct {
	Openssl340ProgramSpecs
	Openssl340MapSpecs
}

// Openssl340Specs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl340ProgramSpecs struct {
	SSL_readEntryNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.ProgramSpec `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.ProgramSpec `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.ProgramSpec `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.ProgramSpec `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.ProgramSpec `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.ProgramSpec `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.ProgramSpec `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.ProgramSpec `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.ProgramSpec `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.ProgramSpec `ebpf:"SSL_write_ret_offset"`
}

// Openssl340MapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type Openssl340MapSpecs struct {
	ActiveSslReadArgsMap  *ebpf.MapSpec `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.MapSpec `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.MapSpec `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.MapSpec `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.MapSpec `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.MapSpec `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.MapSpec `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.MapSpec `ebpf:"filter_pidns_map"`
	Rb                    *ebpf.MapSpec `ebpf:"rb"`
	SslDataMap            *ebpf.MapSpec `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.MapSpec `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.MapSpec `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.MapSpec `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.MapSpec `ebpf:"syscall_rb"`
}

// Openssl340Objects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl340Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl340Objects struct {
	Openssl340Programs
	Openssl340Maps
}

func (o *Openssl340Objects) Close() error {
	return _Openssl340Close(
		&o.Openssl340Programs,
		&o.Openssl340Maps,
	)
}

// Openssl340Maps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl340Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl340Maps struct {
	ActiveSslReadArgsMap  *ebpf.Map `ebpf:"active_ssl_read_args_map"`
	ActiveSslWriteArgsMap *ebpf.Map `ebpf:"active_ssl_write_args_map"`
	ConnEvtRb             *ebpf.Map `ebpf:"conn_evt_rb"`
	ConnInfoMap           *ebpf.Map `ebpf:"conn_info_map"`
	FilterMntnsMap        *ebpf.Map `ebpf:"filter_mntns_map"`
	FilterNetnsMap        *ebpf.Map `ebpf:"filter_netns_map"`
	FilterPidMap          *ebpf.Map `ebpf:"filter_pid_map"`
	FilterPidnsMap        *ebpf.Map `ebpf:"filter_pidns_map"`
	Rb                    *ebpf.Map `ebpf:"rb"`
	SslDataMap            *ebpf.Map `ebpf:"ssl_data_map"`
	SslRb                 *ebpf.Map `ebpf:"ssl_rb"`
	SslUserSpaceCallMap   *ebpf.Map `ebpf:"ssl_user_space_call_map"`
	SyscallDataMap        *ebpf.Map `ebpf:"syscall_data_map"`
	SyscallRb             *ebpf.Map `ebpf:"syscall_rb"`
}

func (m *Openssl340Maps) Close() error {
	return _Openssl340Close(
		m.ActiveSslReadArgsMap,
		m.ActiveSslWriteArgsMap,
		m.ConnEvtRb,
		m.ConnInfoMap,
		m.FilterMntnsMap,
		m.FilterNetnsMap,
		m.FilterPidMap,
		m.FilterPidnsMap,
		m.Rb,
		m.SslDataMap,
		m.SslRb,
		m.SslUserSpaceCallMap,
		m.SyscallDataMap,
		m.SyscallRb,
	)
}

// Openssl340Programs contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadOpenssl340Objects or ebpf.CollectionSpec.LoadAndAssign.
type Openssl340Programs struct {
	SSL_readEntryNestedSyscall    *ebpf.Program `ebpf:"SSL_read_entry_nested_syscall"`
	SSL_readEntryOffset           *ebpf.Program `ebpf:"SSL_read_entry_offset"`
	SSL_readExEntryNestedSyscall  *ebpf.Program `ebpf:"SSL_read_ex_entry_nested_syscall"`
	SSL_readExRetNestedSyscall    *ebpf.Program `ebpf:"SSL_read_ex_ret_nested_syscall"`
	SSL_readRetNestedSyscall      *ebpf.Program `ebpf:"SSL_read_ret_nested_syscall"`
	SSL_readRetOffset             *ebpf.Program `ebpf:"SSL_read_ret_offset"`
	SSL_writeEntryNestedSyscall   *ebpf.Program `ebpf:"SSL_write_entry_nested_syscall"`
	SSL_writeEntryOffset          *ebpf.Program `ebpf:"SSL_write_entry_offset"`
	SSL_writeExEntryNestedSyscall *ebpf.Program `ebpf:"SSL_write_ex_entry_nested_syscall"`
	SSL_writeExRetNestedSyscall   *ebpf.Program `ebpf:"SSL_write_ex_ret_nested_syscall"`
	SSL_writeRetNestedSyscall     *ebpf.Program `ebpf:"SSL_write_ret_nested_syscall"`
	SSL_writeRetOffset            *ebpf.Program `ebpf:"SSL_write_ret_offset"`
}

func (p *Openssl340Programs) Close() error {
	return _Openssl340Close(
		p.SSL_readEntryNestedSyscall,
		p.SSL_readEntryOffset,
		p.SSL_readExEntryNestedSyscall,
		p.SSL_readExRetNestedSyscall,
		p.SSL_readRetNestedSyscall,
		p.SSL_readRetOffset,
		p.SSL_writeEntryNestedSyscall,
		p.SSL_writeEntryOffset,
		p.SSL_writeExEntryNestedSyscall,
		p.SSL_writeExRetNestedSyscall,
		p.SSL_writeRetNestedSyscall,
		p.SSL_writeRetOffset,
	)
}

func _Openssl340Close(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed openssl340_arm64_bpfel.o
var _Openssl340Bytes []byte
