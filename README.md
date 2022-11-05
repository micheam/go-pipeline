# pipeline

## Source

- [x] Stream[T] - given an input slice emits to output stream
- [x] StreamFn[T] - given an generator func emits result to output stream
- [ ] Queue - make queue onto which elements can be pushed for emitting from the source.
- [ ] Range - Emit each integer in a range, with an option to take bigger steps than 1.
- [ ] Repeat - Stream a single object repeatedly.

## [Sink](https://doc.akka.io/docs/akka/current/stream/operators/index.html#sink-operators)

- [x] Collect[T] - Collect all values emitted from the stream into a slice ..
- [ ] ForEach[T] - Invoke a given procedure for each element received.
- [ ] Head
- [ ] First
- [ ] Tail
- [ ] Last

## IO Sinks and Sources

- [ ] FromReader  - Emits the contents of a io.Reader.
- [ ] ToWriter    - Create a sink which will write incoming Byte to a given io.Writer.

## Fan-out

- [x] Broadcast[T] – (1 input, N outputs) given an input element emits to each output
- [ ] Balance[T] – (1 input, N outputs) given an input element emits to one of its output ports
- [ ] UnzipWith[In,A,B,...] – (1 input, N outputs) takes a function of 1 input that given a value for each input emits N output elements
- [ ] UnZip[A,B] – (1 input, 2 outputs) splits a stream of (A,B) tuples into two streams, one of type A and one of type B

## Fan-in

- [x] Merge[In] – (N inputs , 1 output) picks randomly from inputs pushing them one by one to its output
- [ ] MergePreferred[In] – like Merge but if elements are available on preferred port, it picks from it, otherwise randomly from others
- [ ] MergePrioritized[In] – like Merge but if elements are available on all input ports, it picks from them randomly based on their priority
- [ ] MergeLatest[In] – (N inputs, 1 output) emits List[In], when i-th input stream emits element, then i-th element in emitted list is updated
- [ ] MergeSequence[In] – (N inputs, 1 output) emits List[In], where the input streams must represent a partitioned sequence that must be merged back together in order
- [ ] ZipWith[A,B,...,Out] – (N inputs, 1 output) which takes a function of N inputs that given a value for each input emits 1 output element
- [ ] Zip[A,B] – (2 inputs, 1 output) is a ZipWith specialised to zipping input streams of A and B into a (A,B) tuple stream
- [ ] Concat[A] – (2 inputs, 1 output) concatenates two streams (first consume one, then the second one)

## Misc

- [x] Take   - Pass N of incoming elements downstream and then close.
- [ ] Filter - Apply filter Fn to incoming elements, then pass downstream.


## Usage

TBD

## Requirements

TBD

## Installation

TBD

## License

TBD

## Author

Michito Maeda <https://micheam.com>

## Acknowledgments

- API heavily inspired by [Operators • Akka Documentation](https://doc.akka.io/docs/akka/current/stream/operators/index.html)

