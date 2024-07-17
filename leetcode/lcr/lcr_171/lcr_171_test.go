package lcr_171

import (
	. "github.com/majunmin/leetcode-go/common"
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 8,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
							Next: &ListNode{
								Val: 8,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
									},
								},
							},
						},
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
