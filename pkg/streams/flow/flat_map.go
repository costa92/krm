package flow

import "github.com/costa92/krm/pkg/streams"

// FlatMapFunction represents a FlatMap transformation function.
// zh: FlatMapFunction 表示一个 FlatMap 转换函数。
// 平摊映射（FlatMap）
type FlatMapFunction[T, R any] func(T) []R

// FlatMap takes one element and produces zero, one, or more elements.
//
// in  -- 1 -- 2 ---- 3 -- 4 ------ 5 --
//
// [ -------- FlatMapFunction -------- ]
//
// out -- 1' - 2' -------- 4'- 4" - 5' -.
type FlatMap[T, R any] struct {
	flatMapFunction FlatMapFunction[T, R]
	in              chan any
	out             chan any
	parallelism     uint
}

// Verify FlatMap satisfies the Flow interface.
// zh: 验证 FlatMap 满足 Flow 接口。
var _ streams.Flow = (*FlatMap[any, any])(nil)

// NewFlatMap returns a new FlatMap instance.
// zh: NewFlatMap 返回一个新的 FlatMap 实例。
// flatMapFunction is the FlatMap transformation function.
// parallelism is the flow parallelism factor. In case the events order matters, use parallelism = 1.
// zh: flatMapFunction 是 FlatMap 转换函数。
// zh: parallelism 是流并行度因子。如果事件顺序很重要，请使用 parallelism = 1。
func NewFlatMap[T, R any](flatMapFunction FlatMapFunction[T, R], parallelism uint) *FlatMap[T, R] {
	flatMap := &FlatMap[T, R]{
		flatMapFunction: flatMapFunction,
		in:              make(chan any),
		out:             make(chan any),
		parallelism:     parallelism,
	}
	go flatMap.doStream()

	return flatMap
}

// Via streams data through the given flow.
// zh: 通过给定的流传输数据。
func (fm *FlatMap[T, R]) Via(flow streams.Flow) streams.Flow {
	go fm.transmit(flow)
	return flow
}

// To streams data to the given sink.
// zh: 将数据流到给定的接收器。
func (fm *FlatMap[T, R]) To(sink streams.Sink) {
	fm.transmit(sink)
}

// Out returns an output channel for sending data.
// zh: Out 返回一个用于发送数据的输出通道。
func (fm *FlatMap[T, R]) Out() <-chan any {
	return fm.out
}

// In returns an input channel for receiving data.
// zh: In 返回一个用于接收数据的输入通道。
func (fm *FlatMap[T, R]) In() chan<- any {
	return fm.in
}

func (fm *FlatMap[T, R]) transmit(inlet streams.Inlet) {
	for element := range fm.Out() {
		inlet.In() <- element
	}
	close(inlet.In())
}

func (fm *FlatMap[T, R]) doStream() {
	sem := make(chan struct{}, fm.parallelism)
	for elem := range fm.in {
		sem <- struct{}{}
		go func(element T) {
			defer func() { <-sem }()
			result := fm.flatMapFunction(element)
			for _, item := range result {
				fm.out <- item
			}
		}(elem.(T))
	}
	for i := 0; i < int(fm.parallelism); i++ {
		sem <- struct{}{}
	}
	close(fm.out)
}
