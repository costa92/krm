package streams

// Inlet represents a type that exposes one open input.
type Inlet interface {
	In() chan<- any
}

// Outlet represents a type that exposes one open output.
// zh: Outlet 表示一个暴露一个开放输出的类型。
type Outlet interface {
	Out() <-chan any
}

// Source represents a set of stream processing steps that has one open output.
// zh: Source 表示一组流处理步骤，它有一个开放的输出。
type Source interface {
	Outlet
	Via(Flow) Flow
}

// Flow represents a set of stream processing steps that has one open input and one open output.
// zh: Flow 表示一组流处理步骤，它有一个开放的输入和一个开放的输出。
type Flow interface {
	Inlet
	Outlet
	Via(Flow) Flow
	To(Sink)
}

// Sink represents a set of stream processing steps that has one open input.
// zh: Sink 表示一组流处理步骤，它有一个开放的输入。
type Sink interface {
	Inlet
}
