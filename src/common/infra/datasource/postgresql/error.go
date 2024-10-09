package postgresql

type ScanRowError struct {
	Cause string
}

func NewScanRowError(cause string) ScanRowError {
	return ScanRowError{Cause: cause}
}

func (e ScanRowError) Error() string {
	return "Fail to scan row: " + e.Cause
}
