package heap_test

import (
	"testing"
	"time"

	heap "github.com/tmdgusya/go-data-structure/heap"
)

// int를 위한 wrapper 타입
type IntValue struct {
	Value int
}

func (i IntValue) GetPriority() int {
	return i.Value
}

// 실제 사용 예시: TaskRecord
type TaskRecord struct {
	Priority int
	TaskName string
	Deadline time.Time
}

func (t TaskRecord) GetPriority() int {
	return t.Priority
}

// 실제 사용 예시: Job
type Job struct {
	Priority   int
	JobID      string
	Importance float64
}

func (j Job) GetPriority() int {
	return j.Priority
}

func TestInsertSingleValue(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	h.Insert(IntValue{Value: 5})

	if h.Size() != 1 {
		t.Error("Insert 후 힙의 크기가 1이 아닙니다")
	}
}

func TestInsertMultipleValues(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 여러 값 삽입
	values := []int{10, 20, 15, 30, 40}
	for _, v := range values {
		h.Insert(IntValue{Value: v})
	}

	if h.Size() != len(values) {
		t.Errorf("Insert 후 힙의 크기가 올바르지 않습니다. 기대값: %d, 실제값: %d", len(values), h.Size())
	}
}

func TestInsertMaxHeapProperty(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 값들을 무작위 순서로 삽입
	h.Insert(IntValue{Value: 5})
	h.Insert(IntValue{Value: 15})
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 20})
	h.Insert(IntValue{Value: 8})

	// Max Heap이므로 루트는 가장 큰 값이어야 함
	if h.Peek().Value != 20 {
		t.Errorf("Max Heap 속성이 올바르지 않습니다. 루트 값이 최대값이 아닙니다. 기대값: 20, 실제값: %d", h.Peek().Value)
	}
}

func TestInsertWithResize(t *testing.T) {
	initial_size := 2
	h := heap.NewHeap[IntValue](initial_size)

	// 초기 크기보다 많은 값 삽입하여 Resize 트리거
	for i := 1; i <= 5; i++ {
		h.Insert(IntValue{Value: i * 10})
	}

	if h.Size() != 5 {
		t.Errorf("Resize 후 힙의 크기가 올바르지 않습니다. 기대값: 5, 실제값: %d", h.Size())
	}

	// Resize 후에도 Max Heap 속성 유지되는지 확인
	if h.Peek().Value != 50 {
		t.Errorf("Resize 후 Max Heap 속성이 유지되지 않았습니다. 기대값: 50, 실제값: %d", h.Peek().Value)
	}
}

func TestInsertAscendingOrder(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 오름차순으로 삽입
	for i := 1; i <= 10; i++ {
		h.Insert(IntValue{Value: i})
	}

	// Max Heap이므로 가장 큰 값이 루트에 있어야 함
	if h.Peek().Value != 10 {
		t.Errorf("오름차순 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek().Value)
	}
}

func TestInsertDescendingOrder(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 내림차순으로 삽입
	for i := 10; i >= 1; i-- {
		h.Insert(IntValue{Value: i})
	}

	// Max Heap이므로 가장 큰 값이 루트에 있어야 함
	if h.Peek().Value != 10 {
		t.Errorf("내림차순 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek().Value)
	}
}

func TestInsertDuplicateValues(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 중복된 값 삽입
	h.Insert(IntValue{Value: 5})
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 5})
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 5})

	if h.Size() != 5 {
		t.Errorf("중복 값 삽입 후 힙의 크기가 올바르지 않습니다. 기대값: 5, 실제값: %d", h.Size())
	}

	if h.Peek().Value != 10 {
		t.Errorf("중복 값 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: 10, 실제값: %d", h.Peek().Value)
	}
}

func TestInsertLargeData(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 많은 양의 데이터 삽입
	max_value := 0
	for i := 1; i <= 1000; i++ {
		value := i * 7 % 997 // 무작위처럼 보이는 값
		if value > max_value {
			max_value = value
		}
		h.Insert(IntValue{Value: value})
	}

	if h.Size() != 1000 {
		t.Errorf("대량 데이터 삽입 후 힙의 크기가 올바르지 않습니다. 기대값: 1000, 실제값: %d", h.Size())
	}

	if h.Peek().Value != max_value {
		t.Errorf("대량 데이터 삽입 후 Max Heap 속성이 올바르지 않습니다. 기대값: %d, 실제값: %d", max_value, h.Peek().Value)
	}
}

func TestRemoveSingleElement(t *testing.T) {
	h := heap.NewHeap[IntValue](10)
	h.Insert(IntValue{Value: 42})

	removed := h.Remove()
	if removed.Value != 42 {
		t.Errorf("Remove 반환값이 올바르지 않습니다. 기대값: 42, 실제값: %d", removed.Value)
	}

	if h.Size() != 0 {
		t.Errorf("Remove 후 힙의 크기가 올바르지 않습니다. 기대값: 0, 실제값: %d", h.Size())
	}
}

func TestRemoveFromEmptyHeap(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	removed := h.Remove()
	if removed.Value != 0 {
		t.Errorf("빈 힙에서 Remove 시 zero value를 반환해야 합니다. 기대값: 0, 실제값: %d", removed.Value)
	}
}

func TestRemoveMaxHeapProperty(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// [50, 30, 40, 10, 20] 힙 생성
	values := []int{50, 30, 40, 10, 20}
	for _, v := range values {
		h.Insert(IntValue{Value: v})
	}

	// 첫 번째 Remove: 50이 나와야 함
	removed := h.Remove()
	if removed.Value != 50 {
		t.Errorf("첫 번째 Remove가 최대값을 반환하지 않았습니다. 기대값: 50, 실제값: %d", removed.Value)
	}

	// Remove 후에도 Max Heap 속성 유지되는지 확인
	if h.Peek().Value != 40 {
		t.Errorf("Remove 후 Max Heap 속성이 유지되지 않았습니다. 기대값: 40, 실제값: %d", h.Peek().Value)
	}
}

func TestRemoveMultipleElements(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 1부터 10까지 삽입
	for i := 1; i <= 10; i++ {
		h.Insert(IntValue{Value: i})
	}

	// 내림차순으로 제거되는지 확인
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i, exp := range expected {
		removed := h.Remove()
		if removed.Value != exp {
			t.Errorf("Remove %d번째: 기대값 %d, 실제값 %d", i+1, exp, removed.Value)
		}
	}

	// 모두 제거 후 힙이 비어있는지 확인
	if h.Size() != 0 {
		t.Errorf("모든 원소 제거 후 힙의 크기가 0이 아닙니다. 실제값: %d", h.Size())
	}
}

func TestRemoveAndInsertMixed(t *testing.T) {
	h := heap.NewHeap[IntValue](5)

	// 삽입과 삭제를 섞어서 수행
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 20})
	h.Insert(IntValue{Value: 15})

	if h.Remove().Value != 20 {
		t.Error("첫 번째 Remove가 20을 반환하지 않았습니다")
	}

	h.Insert(IntValue{Value: 25})
	h.Insert(IntValue{Value: 5})

	if h.Peek().Value != 25 {
		t.Errorf("삽입/삭제 후 최대값이 올바르지 않습니다. 기대값: 25, 실제값: %d", h.Peek().Value)
	}
}

func TestRemoveWithDuplicates(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 중복 값 삽입
	values := []int{5, 10, 5, 10, 5}
	for _, v := range values {
		h.Insert(IntValue{Value: v})
	}

	// 첫 두 번은 10이 나와야 함
	if h.Remove().Value != 10 {
		t.Error("중복 값 제거 - 첫 번째 10이 나오지 않았습니다")
	}
	if h.Remove().Value != 10 {
		t.Error("중복 값 제거 - 두 번째 10이 나오지 않았습니다")
	}

	// 나머지는 5
	for i := 0; i < 3; i++ {
		if h.Remove().Value != 5 {
			t.Errorf("중복 값 제거 - %d번째 5가 나오지 않았습니다", i+1)
		}
	}
}

func TestRemoveAllElements(t *testing.T) {
	h := heap.NewHeap[IntValue](5)

	values := []int{15, 10, 20, 5, 25}
	for _, v := range values {
		h.Insert(IntValue{Value: v})
	}

	// 모든 원소를 내림차순으로 제거
	prev := h.Remove() // 25
	for h.Size() > 0 {
		current := h.Remove()
		if current.Value > prev.Value {
			t.Errorf("Remove 순서가 올바르지 않습니다. 이전값: %d, 현재값: %d", prev.Value, current.Value)
		}
		prev = current
	}

	// 빈 힙에서 한 번 더 Remove
	if h.Remove().Value != 0 {
		t.Error("빈 힙에서 Remove 시 zero value를 반환해야 합니다")
	}
}

func TestTaskRecordHeap(t *testing.T) {
	h := heap.NewHeap[TaskRecord](10)

	// 다양한 우선순위의 Task 삽입
	h.Insert(TaskRecord{
		Priority: 3,
		TaskName: "Low priority task",
		Deadline: time.Now().Add(24 * time.Hour),
	})
	h.Insert(TaskRecord{
		Priority: 10,
		TaskName: "Critical task",
		Deadline: time.Now().Add(1 * time.Hour),
	})
	h.Insert(TaskRecord{
		Priority: 5,
		TaskName: "Medium priority task",
		Deadline: time.Now().Add(12 * time.Hour),
	})
	h.Insert(TaskRecord{
		Priority: 8,
		TaskName: "High priority task",
		Deadline: time.Now().Add(6 * time.Hour),
	})

	// 가장 높은 우선순위(10)가 먼저 나와야 함
	task := h.Remove()
	if task.Priority != 10 {
		t.Errorf("최고 우선순위 Task가 제거되지 않았습니다. 기대값: 10, 실제값: %d", task.Priority)
	}
	if task.TaskName != "Critical task" {
		t.Errorf("잘못된 Task가 제거되었습니다. 기대값: Critical task, 실제값: %s", task.TaskName)
	}

	// 다음은 우선순위 8
	task = h.Remove()
	if task.Priority != 8 {
		t.Errorf("두 번째로 높은 우선순위 Task가 제거되지 않았습니다. 기대값: 8, 실제값: %d", task.Priority)
	}

	// Peek으로 확인 (우선순위 5가 최상위에 있어야 함)
	if h.Peek().Priority != 5 {
		t.Errorf("Peek이 올바르지 않습니다. 기대값: 5, 실제값: %d", h.Peek().Priority)
	}
}

func TestJobHeap(t *testing.T) {
	h := heap.NewHeap[Job](5)

	// 여러 Job 삽입
	h.Insert(Job{Priority: 1, JobID: "JOB-001", Importance: 0.5})
	h.Insert(Job{Priority: 5, JobID: "JOB-002", Importance: 0.9})
	h.Insert(Job{Priority: 3, JobID: "JOB-003", Importance: 0.7})

	// 우선순위가 가장 높은 Job 확인
	topJob := h.Peek()
	if topJob.Priority != 5 {
		t.Errorf("최고 우선순위 Job이 올바르지 않습니다. 기대값: 5, 실제값: %d", topJob.Priority)
	}
	if topJob.JobID != "JOB-002" {
		t.Errorf("최고 우선순위 Job ID가 올바르지 않습니다. 기대값: JOB-002, 실제값: %s", topJob.JobID)
	}

	// 순서대로 제거
	job1 := h.Remove()
	job2 := h.Remove()
	job3 := h.Remove()

	if job1.Priority != 5 || job2.Priority != 3 || job3.Priority != 1 {
		t.Error("Job이 우선순위 순서대로 제거되지 않았습니다")
	}
}

func TestMixedTypesWithPrioritized(t *testing.T) {
	// IntValue 힙
	intHeap := heap.NewHeap[IntValue](5)
	intHeap.Insert(IntValue{Value: 10})
	intHeap.Insert(IntValue{Value: 20})

	// TaskRecord 힙
	taskHeap := heap.NewHeap[TaskRecord](5)
	taskHeap.Insert(TaskRecord{Priority: 1, TaskName: "Task A"})
	taskHeap.Insert(TaskRecord{Priority: 2, TaskName: "Task B"})

	// Job 힙
	jobHeap := heap.NewHeap[Job](5)
	jobHeap.Insert(Job{Priority: 100, JobID: "JOB-X"})
	jobHeap.Insert(Job{Priority: 200, JobID: "JOB-Y"})

	// 각각 독립적으로 동작하는지 확인
	if intHeap.Peek().Value != 20 {
		t.Error("IntValue 힙이 올바르지 않습니다")
	}
	if taskHeap.Peek().Priority != 2 {
		t.Error("TaskRecord 힙이 올바르지 않습니다")
	}
	if jobHeap.Peek().Priority != 200 {
		t.Error("Job 힙이 올바르지 않습니다")
	}
}

func TestUpdateIncreasePriority(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// [50, 30, 40, 10, 20] 힙 생성
	h.Insert(IntValue{Value: 50})
	h.Insert(IntValue{Value: 30})
	h.Insert(IntValue{Value: 40})
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 20})

	// 인덱스 3 (값 10)을 60으로 업데이트 -> 새로운 최대값
	success := h.Update(3, IntValue{Value: 60})
	if !success {
		t.Error("Update가 실패했습니다")
	}

	// 60이 최상위에 있어야 함
	if h.Peek().Value != 60 {
		t.Errorf("우선순위 증가 후 Peek이 올바르지 않습니다. 기대값: 60, 실제값: %d", h.Peek().Value)
	}

	// Remove로 순서 확인
	if h.Remove().Value != 60 {
		t.Error("첫 번째 Remove가 60이 아닙니다")
	}
	if h.Remove().Value != 50 {
		t.Error("두 번째 Remove가 50이 아닙니다")
	}
}

func TestUpdateDecreasePriority(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// [50, 30, 40, 10, 20] 힙 생성
	h.Insert(IntValue{Value: 50})
	h.Insert(IntValue{Value: 30})
	h.Insert(IntValue{Value: 40})
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 20})

	// 인덱스 0 (루트, 값 50)을 5로 업데이트 -> 아래로 이동
	success := h.Update(0, IntValue{Value: 5})
	if !success {
		t.Error("Update가 실패했습니다")
	}

	// 40이 최상위에 있어야 함
	if h.Peek().Value != 40 {
		t.Errorf("우선순위 감소 후 Peek이 올바르지 않습니다. 기대값: 40, 실제값: %d", h.Peek().Value)
	}

	// 모두 Remove로 확인
	values := []int{}
	for h.Size() > 0 {
		values = append(values, h.Remove().Value)
	}

	// 내림차순이어야 함
	expected := []int{40, 30, 20, 10, 5}
	for i, exp := range expected {
		if values[i] != exp {
			t.Errorf("Remove %d번째: 기대값 %d, 실제값 %d", i+1, exp, values[i])
		}
	}
}

func TestUpdateSamePriority(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	h.Insert(IntValue{Value: 50})
	h.Insert(IntValue{Value: 30})
	h.Insert(IntValue{Value: 40})

	// 같은 우선순위로 업데이트
	success := h.Update(1, IntValue{Value: 30})
	if !success {
		t.Error("Update가 실패했습니다")
	}

	// 힙 속성이 유지되어야 함
	if h.Peek().Value != 50 {
		t.Errorf("같은 우선순위 업데이트 후 Peek이 올바르지 않습니다. 기대값: 50, 실제값: %d", h.Peek().Value)
	}

	if h.Size() != 3 {
		t.Errorf("힙 크기가 변경되었습니다. 기대값: 3, 실제값: %d", h.Size())
	}
}

func TestUpdateInvalidIndex(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	h.Insert(IntValue{Value: 50})
	h.Insert(IntValue{Value: 30})

	// 범위 밖 인덱스
	if h.Update(10, IntValue{Value: 100}) {
		t.Error("범위 밖 인덱스에 대해 Update가 성공했습니다")
	}

	// 음수 인덱스
	if h.Update(-1, IntValue{Value: 100}) {
		t.Error("음수 인덱스에 대해 Update가 성공했습니다")
	}

	// 힙이 변경되지 않았는지 확인
	if h.Size() != 2 {
		t.Error("잘못된 Update로 힙 크기가 변경되었습니다")
	}
	if h.Peek().Value != 50 {
		t.Error("잘못된 Update로 힙이 변경되었습니다")
	}
}

func TestUpdateTaskRecord(t *testing.T) {
	h := heap.NewHeap[TaskRecord](10)

	// Task 삽입
	h.Insert(TaskRecord{Priority: 5, TaskName: "Task A"})
	h.Insert(TaskRecord{Priority: 3, TaskName: "Task B"})
	h.Insert(TaskRecord{Priority: 4, TaskName: "Task C"})

	// Task B (인덱스 1 또는 2)의 우선순위를 10으로 올림
	// 힙 구조상 인덱스를 정확히 알기 어려우므로,
	// 모든 원소를 찾아서 업데이트
	for i := 0; i < h.Size(); i++ {
		// 내부 접근이 어려우므로 간접적으로 테스트
		// 인덱스 1을 업데이트
		if i == 1 {
			success := h.Update(i, TaskRecord{Priority: 10, TaskName: "Task B Updated"})
			if !success {
				t.Error("TaskRecord Update가 실패했습니다")
			}
		}
	}

	// 최고 우선순위는 10이어야 함
	top := h.Peek()
	if top.Priority != 10 {
		t.Errorf("TaskRecord Update 후 Peek이 올바르지 않습니다. 기대값: 10, 실제값: %d", top.Priority)
	}
	if top.TaskName != "Task B Updated" {
		t.Errorf("TaskRecord 내용이 올바르지 않습니다. 기대값: Task B Updated, 실제값: %s", top.TaskName)
	}
}

func TestUpdateMultipleTimes(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 초기 힙: [10, 5, 8, 3, 2]
	h.Insert(IntValue{Value: 10})
	h.Insert(IntValue{Value: 5})
	h.Insert(IntValue{Value: 8})
	h.Insert(IntValue{Value: 3})
	h.Insert(IntValue{Value: 2})

	// 첫 번째 업데이트: 인덱스 4를 15로
	h.Update(4, IntValue{Value: 15})

	// 15가 최댓값이어야 함
	if h.Peek().Value != 15 {
		t.Errorf("첫 번째 Update 후 Peek이 올바르지 않습니다. 기대값: 15, 실제값: %d", h.Peek().Value)
	}

	// 두 번째 업데이트: 루트(인덱스 0)를 20으로
	h.Update(0, IntValue{Value: 20})

	// 20이 최댓값이어야 함
	if h.Peek().Value != 20 {
		t.Errorf("두 번째 Update 후 Peek이 올바르지 않습니다. 기대값: 20, 실제값: %d", h.Peek().Value)
	}

	// 세 번째 업데이트: 루트를 1로 낮춤
	h.Update(0, IntValue{Value: 1})

	// 이제 다른 값이 최댓값이어야 함 (1은 아님)
	if h.Peek().Value == 1 {
		t.Error("루트를 낮춘 후에도 여전히 루트에 있습니다")
	}

	// 전체 순서 확인 - 힙 속성 유지되는지
	values := []int{}
	for h.Size() > 0 {
		values = append(values, h.Remove().Value)
	}

	// 내림차순 확인
	for i := 0; i < len(values)-1; i++ {
		if values[i] < values[i+1] {
			t.Errorf("힙 속성이 깨졌습니다. %d번째: %d, %d번째: %d", i, values[i], i+1, values[i+1])
		}
	}
}

func TestUpdateEmptyHeap(t *testing.T) {
	h := heap.NewHeap[IntValue](10)

	// 빈 힙에서 Update 시도
	if h.Update(0, IntValue{Value: 100}) {
		t.Error("빈 힙에서 Update가 성공했습니다")
	}
}
