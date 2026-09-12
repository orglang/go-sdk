package prog

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/orglang/go-sdk/internal/prog"
	pooltypedef "github.com/orglang/go-sdk/pool/typedef"
	pooltypeexp "github.com/orglang/go-sdk/pool/typeexp"
	proctypedef "github.com/orglang/go-sdk/proc/typedef"
	proctypeexp "github.com/orglang/go-sdk/proc/typeexp"
)

func TestMsgFromText_Success(t *testing.T) {
	tests := map[string]struct {
		text string
		spec Spec
	}{
		"one pool type": {
			text: `
				pool type t1 {}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1"},
			}},
		},
		"two pool types": {
			text: `
				pool type t1 {}
				pool type t2 {}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1"},
				pooltypedef.DefSpec{TypeQN: "t2"}},
			},
		},
		"pool type and proc type (direct order)": {
			text: `
				pool type t1 {}
				proc type t2 {}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1"},
				proctypedef.DefSpec{TypeQN: "t2"}},
			},
		},
		"pool type and proc type (reverse order)": {
			text: `
				proc type t1 {}
				pool type t2 {}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				proctypedef.DefSpec{TypeQN: "t1"},
				pooltypedef.DefSpec{TypeQN: "t2"}},
			},
		},
		"pool type with single exp (kind of 1st field)": {
			text: `
				pool type t1 {
					up {}
				}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1", TypeExp: pooltypeexp.ExpSpec{
					K:  "up",
					Up: &pooltypeexp.ShiftSpec{},
				}},
			}},
		},
		"pool type with single exp (kind of 2nd field)": {
			text: `
				pool type t1 {
					down {}
				}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1", TypeExp: pooltypeexp.ExpSpec{
					K:    "down",
					Down: &pooltypeexp.ShiftSpec{},
				}},
			}},
		},
		"pool type with multiple exps": {
			text: `
				pool type t1 {
					up {
						with (foo) {
							down {
								link t1
							}
						}
					}
				}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				pooltypedef.DefSpec{TypeQN: "t1", TypeExp: pooltypeexp.ExpSpec{
					K: "up",
					Up: &pooltypeexp.ShiftSpec{ContExp: pooltypeexp.ExpSpec{
						K: "with",
						With: &pooltypeexp.LaborSpec{ProcQNs: []string{"foo"}, ContExp: pooltypeexp.ExpSpec{
							K: "down",
							Down: &pooltypeexp.ShiftSpec{ContExp: pooltypeexp.ExpSpec{
								K:    "link",
								Link: &pooltypeexp.LinkSpec{TypeQN: "t1"},
							}},
						}},
					}},
				}},
			}},
		},
		"proc type with multiple exps": {
			text: `
				proc type t1 {
					up {
						with {
							@foo {
								down {
									link t1
								}
							}
							@bar {
								down {
									link t1
								}
							}
						}
					}
				}
			`,
			spec: Spec{Exps: []prog.ProgExp{
				proctypedef.DefSpec{TypeQN: "t1", TypeExp: proctypeexp.ExpSpec{
					K: "up",
					Up: &proctypeexp.ShiftSpec{Cont: proctypeexp.ExpSpec{
						K: "with",
						With: &proctypeexp.SumSpec{Choices: []proctypeexp.ChoiceSpec{
							{LabQN: "foo", Cont: proctypeexp.ExpSpec{
								K: "down",
								Down: &proctypeexp.ShiftSpec{Cont: proctypeexp.ExpSpec{
									K:    "link",
									Link: &proctypeexp.LinkSpec{TypeQN: "t1"},
								}},
							}},
							{LabQN: "bar", Cont: proctypeexp.ExpSpec{
								K: "down",
								Down: &proctypeexp.ShiftSpec{Cont: proctypeexp.ExpSpec{
									K:    "link",
									Link: &proctypeexp.LinkSpec{TypeQN: "t1"},
								}},
							}},
						}},
					}},
				}},
			}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := MsgFromText(test.text)
			if err != nil {
				t.Fatal(err)
			}
			if !cmp.Equal(got, test.spec) {
				t.Errorf("\ntext:%s\nspec:\n%s", test.text, cmp.Diff(test.spec, got))
			}
		})
	}
}

func TestMsgFromText_Failure(t *testing.T) {
	tests := map[string]struct {
		text string
		msg  string
	}{
		"pool type with two exps (direct order)": {
			text: `
				pool type t1 {
					up {}
					down {}
				}
			`,
			msg: "unexpected token \"down\"",
		},
		"pool type with two exps (reverse order)": {
			text: `
				pool type t1 {
					down {}
					up {}
				}
			`,
			msg: "unexpected token \"up\"",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := MsgFromText(test.text)
			if err == nil || !strings.Contains(err.Error(), test.msg) {
				t.Fatalf("want: %v; got: %v", test.msg, err)
			}
		})
	}
}
