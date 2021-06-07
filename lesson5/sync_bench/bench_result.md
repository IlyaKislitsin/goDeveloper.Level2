# Результаты теста производительности
    cpu: Intel(R) Core(TM) i5-6300HQ CPU @ 2.30GHz

### sync.Mutex

##### запись 10%, чтение 90%
    BenchmarkSetAdd10Has90/#00-4     4527974               273.3 ns/op

##### запись 50%, чтение 50%
    BenchmarkSetAdd50Has50/#00-4     4937872               279.4 ns/op

##### запись 90%, чтение 10%
    BenchmarkSetAdd90Has10/#00-4     4297186               268.4 ns/op

### sync.RWMutex

##### запись 10%, чтение 90%
    BenchmarkRWSetAdd10Has90/#00-4           2856403               402.8 ns/op

##### запись 50%, чтение 50%
    BenchmarkRWSetAdd50Has50/#00-4           5827723               206.1 ns/op

##### запись 90%, чтение 10%
    BenchmarkRWSetAdd90Has10/#00-4           6726762               169.8 ns/op

### sync.Map

##### запись 10%, чтение 90%
    BenchmarkMapSetAdd10Has90/#00-4          6450200               193.0 ns/op

##### запись 50%, чтение 50%
    BenchmarkMapSetAdd50Has50/#00-4          5196926               221.9 ns/op

##### запись 90%, чтение 10%
    BenchmarkMapSetAdd90Has10/#00-4          5580976               231.8 ns/op
