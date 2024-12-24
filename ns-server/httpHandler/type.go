package httpHandler

type Response struct {
	Message string `json:"message"`
}

type Port uint16
type SideEnum int8
type ConnDesc struct {
	LocalPort  Port
	RemotePort Port
	RemoteAddr string
	LocalAddr  string
	Pid        uint32
	PidStr     string
	Protocol   uint32
	Side       SideEnum
	StreamId   int
	IsSsl      bool
}
type RecordToK8s struct {
	ReqStr  string `json:"reqStr"`
	RespStr string `json:"respStr"`
}
type ParsedMessage interface {
	FormatToString() string
	FormatToSummaryString() string
	TimestampNs() uint64
	ByteSize() int
	IsReq() bool
	Seq() uint64
}
type ResponseStatus int8
type AnnotatedRecord struct {
	ConnDesc
	RecordToK8s
	StartTs                      uint64
	EndTs                        uint64
	ReqPlainTextSize             int
	RespPlainTextSize            int
	ReqSize                      int
	RespSize                     int
	TotalDuration                float64
	BlackBoxDuration             float64
	CopyToSocketBufferDuration   float64
	ReadFromSocketBufferDuration float64
	// ReqSyscallEventDetails       []SyscallEventDetail
	// RespSyscallEventDetails      []SyscallEventDetail
	// ReqNicEventDetails           []NicEventDetail
	// RespNicEventDetails          []NicEventDetail
	// Json用于存入数据库
	ReqSyscallEventDetailsJson  string
	RespSyscallEventDetailsJson string
	ReqNicEventDetailsJson      string
	RespNicEventDetailsJson     string
}
type SyscallEventDetail PacketEventDetail
type NicEventDetail struct {
	PacketEventDetail
	Attributes map[string]any
}
type PacketEventDetail struct {
	ByteSize  int
	Timestamp uint64
}
