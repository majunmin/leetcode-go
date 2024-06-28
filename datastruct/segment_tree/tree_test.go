package segment_tree

import (
	"fmt"
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	segmentTree := NewSegmentTree(5)
	segmentTree.Insert(1)
	segmentTree.Insert(2)
	segmentTree.Insert(3)
	segmentTree.Insert(4)
	segmentTree.Insert(5)
	fmt.Println(segmentTree.count(1, 1))
	fmt.Println(segmentTree.count(2, 2))
	fmt.Println(segmentTree.count(1, 2))
	fmt.Println(segmentTree.count(1, 5))
}

func TestSegmentTree_Delete(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			st.Delete(tt.args.data)
		})
	}
}

func TestSegmentTree_Insert(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		data int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			st.Insert(tt.args.data)
		})
	}
}

func TestSegmentTree_buildSegmentTreeInternal(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		left  int
		right int
		i     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			st.buildSegmentTreeInternal(tt.args.left, tt.args.right, tt.args.i)
		})
	}
}

func TestSegmentTree_count(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			if got := st.count(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_countInternal(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		left  int
		right int
		idx   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			if got := st.countInternal(tt.args.left, tt.args.right, tt.args.idx); got != tt.want {
				t.Errorf("countInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_getKth(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		left  int
		right int
		k     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			if got := st.getKth(tt.args.left, tt.args.right, tt.args.k); got != tt.want {
				t.Errorf("getKth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegmentTree_getKthInternal(t *testing.T) {
	type fields struct {
		m        int
		segments []*Segment
	}
	type args struct {
		left  int
		right int
		idx   int
		k     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SegmentTree{
				m:        tt.fields.m,
				segments: tt.fields.segments,
			}
			if got := st.getKthInternal(tt.args.left, tt.args.right, tt.args.idx, tt.args.k); got != tt.want {
				t.Errorf("getKthInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}
