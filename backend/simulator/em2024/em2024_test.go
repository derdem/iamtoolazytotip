package em2024_test

import (
	"fmt"
	"testing"

	"github.com/derdem/iamtoolazytotip/simulator/em2024"
)

func TestMultipleGroupSetup(t *testing.T) {
	groups := em2024.CreateEm2024Groups()
	fmt.Println(groups)
}
