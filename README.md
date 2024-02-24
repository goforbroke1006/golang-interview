# Golang Interview

Interview questions for Software (Go/Golang) engineer position.

## Slices

1. What is slice?
2. What is default value for slices?
3. How append function works?
4. Why append inside another function is not good idea? (related to Q1 and Q3)
5. Could you please tell about cases, when you send slice as argument, modify it and everything is okay?

#### Useful links

* `slice` struct [source code](https://go.dev/src/runtime/slice.go#L15)
* `growslice` func [source code](https://go.dev/src/runtime/slice.go#L155)
* `nextslicecap` func [source code](https://go.dev/src/runtime/slice.go#L267)

## Concurrency

6. Why goroutines faster that OS threads? 
7. Could you tell about synchronizing tool from stdlib? 
8. What's difference sync.Mutex and sync.RWMutex? 
9. Why atomics faster that counter with mutex? 
10. How goroutines communicate? How to send information between goroutines? 
11. Why replacing channel with slice+mutex is not good idea? 
12. What's difference between non-buffered channel and buffered channel with size=1.

#### Useful links

* My atomics research [project](https://github.com/goforbroke1006/hello-go-atomic)
* The `hchan` struct [source code](https://go.dev/src/runtime/chan.go#L33)
* The Go Memory Model [off doc](https://go.dev/ref/mem)
* Share Memory By Communicating [off doc](https://go.dev/blog/codelab-share). _"Do not communicate by sharing memory; instead, share memory by communicating."_
* M:N model [discuss on SoF](https://stackoverflow.com/questions/44982392/what-are-the-disadvantages-of-an-mn-threading-model-e-g-goroutines)
* Scheduling in Go [part 1](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)
* Scheduling in Go [part 2](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html)
* Scheduling in Go [part 3](https://www.ardanlabs.com/blog/2018/12/scheduling-in-go-part3.html)

### Useful links

* [Daniil Podolskiy "Collegues, you make me sad" (RU)](https://habr.com/ru/companies/oleg-bunin/articles/521582/)
