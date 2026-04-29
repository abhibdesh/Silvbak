// approver.go
package approvalengine

import (
	// "crypto/sha256"
	// "encoding/hex"
	// "fmt"
	"sync"
)

type ApprovalInput struct {
	TenantID        string
	ApproverID      string
	SubjectID       string
	Action          string
	IdempotencyKey  string
	ClientTimestamp int64
	Metadata        map[string]string
}

type ApprovalReceipt struct {
	GlobalHash      string
	GlobalSeq       int64
	AdminHash       string
	AdminSeq        int64
	AdminSignature  string
	SystemSignature string
}
type Engine struct {
	mu        sync.Mutex
	lasthash  string
	globalseq int64
	adminseq  map[string]int64
	seenkeys  map[string]ApprovalReceipt
}

func NewEngine() *Engine {
	return &Engine{}
}

func ApproveRecord() {

}
