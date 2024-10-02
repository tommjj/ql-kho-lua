package utils

// ExecuteAsync execute async f func return channel
func ExecuteAsync[T any](fc func() (T, error)) (<-chan T, <-chan error) {
	resultChan := make(chan T, 1)
	errChan := make(chan error, 1)

	go func() {
		result, err := fc()
		errChan <- err
		resultChan <- result
		close(resultChan)
		close(errChan)
	}()

	return resultChan, errChan
}
