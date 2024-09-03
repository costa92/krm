package flow

import "github.com/costa92/krm/pkg/streams"

// PassThrough retransmits incoming elements as is.
// zh: PassThrough 重新传输传入的元素。
//
// in  -- 1 -- 2 ---- 3 -- 4 ------ 5 --
//
// out -- 1 -- 2 ---- 3 -- 4 ------ 5 --.

// PassThrough retransmits incoming elements as is.
// zh: PassThrough 重新传输传入的元素。
// 直通（PassThrough）
type PassThrough struct {
	in  chan any
	out chan any
}

// Verify PassThrough satisfies the Flow interface.
// zh: 验证 PassThrough 是否满足 Flow 接口。
var _ streams.Flow = (*PassThrough)(nil)

// NewPassThrough  returns a new PassThrough instance.
// zh: NewPassThrough 返回一个新的 PassThrough 实例。
func NewPassThrough() *PassThrough {
	passThrough := &PassThrough{
		in:  make(chan any),
		out: make(chan any),
	}
	go passThrough.doStream()

	return passThrough
}

// Via streams data through the given flow.
// zh: 通过给定的流传输数据。
// Via streams data through the given flow.
func (pt *PassThrough) Via(flow streams.Flow) streams.Flow {
	go pt.transmit(flow)
	return flow
}

// To streams data to the given sink.
func (pt *PassThrough) To(sink streams.Sink) {
	pt.transmit(sink)
}

// Out returns an output channel for sending data.
func (pt *PassThrough) Out() <-chan any {
	return pt.out
}

// In returns an input channel for receiving data.
func (pt *PassThrough) In() chan<- any {
	return pt.in
}

// transmit streams data from the input to the output.
// zh: transmit 从输入流到输出流。
func (pt *PassThrough) transmit(inlet streams.Inlet) {
	for elem := range pt.Out() {
		inlet.In() <- elem
	}
	close(inlet.In())
}

// doStream streams data from the input to the output.
// zh: doStream 从输入流到输出流。
func (pt *PassThrough) doStream() {
	for elem := range pt.in {
		pt.out <- elem
	}
	close(pt.out)
}
