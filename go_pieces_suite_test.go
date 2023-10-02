package go_pieces_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoPieces(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoPieces Suite")
}
