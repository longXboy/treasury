package eth

import (
	"context"
	"fmt"
	"testing"
)

func TestNotFound(t *testing.T) {
	_, _, err := GetTransactionStatus(context.Background(), "0xee6e2c97f244d641eb1209fb9ebd13bb97234caa7f6fe207b8f7b604ab2669a5")
	fmt.Println(err.Error() == "not found")
}
