package fibonacci

type Service interface {
	Get(pos int64) (int64, error)
	List(min int64, max int64) ([]int64, error)
}
