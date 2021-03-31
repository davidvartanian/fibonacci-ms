package fibonacci

type fibonacciService struct{}

func NewService() Service {
	return &fibonacciService{}
}

func (w *fibonacciService) Get(pos int64) (int64, error) {
	return Mem(pos), nil
}

func (w *fibonacciService) List(min int64, max int64) ([]int64, error) {
	return MemList(min, max), nil
}

func CalculateFib(pos int64, cache map[int64]int64) int64 {
	if pos < 2 {
		cache[pos] = 1
		return 1
	}

	if _, ok := cache[pos]; !ok {
		cache[pos] = CalculateFib(pos-1, cache) + CalculateFib(pos-2, cache)
	}

	return cache[pos]
}

func Mem(pos int64) int64 {
	cache := make(map[int64]int64)
	return CalculateFib(pos, cache)
}

func MemList(min int64, max int64) []int64 {
	cache := make(map[int64]int64)
	var result []int64
	for pos := min; pos <= max; pos++ {
		result = append(result, CalculateFib(pos, cache))
	}
	return result
}
