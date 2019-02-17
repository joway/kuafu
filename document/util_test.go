package document

import (
	"fmt"
	"testing"
)

func TestGenerateDocId(t *testing.T) {
	id := GenerateDocId()
	fmt.Println(id)
	fmt.Println(len(id))
}
