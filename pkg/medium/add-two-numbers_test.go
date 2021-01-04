package medium

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {

	testCases := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want []int
	}{
		{
			name: "Case 1",
			l1: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "Case 2",
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
									},
								},
							},
						},
					},
				},
			},
			l2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
						},
					},
				},
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name: "Case 3",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
				},
			},
			l2: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
					},
				},
			},
			want: []int{0, 1, 0, 1},
		},
		{
			name: "Case 4",
			l1: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			l2: &ListNode{
				Val: 9,
			},
			want: []int{9, 0, 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := addTwoNumbers(testCase.l1, testCase.l2)

			if len(got.toArray()) != len(testCase.want) {
				t.Fatalf("%v failed. Expected %v, got %v", testCase.name, testCase.want, got.toArray())
			}

			for i, v := range got.toArray() {
				if v != testCase.want[i] {
					t.Fatalf("%v failed. Expected %v, got %v", testCase.name, testCase.want, got.toArray())
					break
				}
			}
		})
	}
}
