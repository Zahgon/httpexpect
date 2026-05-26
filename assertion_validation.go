package httpexpect

func validateAssertion(failure *AssertionFailure) error { _ = "STUB: not implemented"; return nil }

func validateType(failure *AssertionFailure) error { _ = "STUB: not implemented"; return nil }

type fieldRequirement uint

const (
	fieldOptional fieldRequirement = iota
	fieldRequired
	fieldDenied
)

type fieldTraits struct {
	Actual   fieldRequirement
	Expected fieldRequirement
	Range    fieldRequirement
	List     fieldRequirement
}

func validateTraits(failure *AssertionFailure, traits fieldTraits) error {
	_ = "STUB: not implemented"
	return nil
}
