package queue_test

import (
	"testing"

	queue "github.com/tmdgusya/go-data-structure/queue"
)

func TestEnqueue(t *testing.T) {
	q := queue.NewQueue(1)

	// Enqueue 여러 값들
	for i := 2; i <= 100; i++ {
		q.Enqueue(i)
	}

	// FIFO이므로 첫 번째 값이 먼저 나와야 함
	if q.Dequeue() != 1 {
		t.Error("Enqueue가 올바르게 작동하지 않았습니다")
	}

	if q.Dequeue() != 2 {
		t.Error("Enqueue가 올바르게 작동하지 않았습니다")
	}
}

func TestDequeue(t *testing.T) {
	q := queue.NewQueue(1)

	// Enqueue 2개의 값 더 추가
	q.Enqueue(2)
	q.Enqueue(3)

	// Dequeue로 FIFO 순서 확인
	if q.Dequeue() != 1 {
		t.Error("Dequeue가 첫 번째 값을 반환하지 않았습니다")
	}

	if q.Dequeue() != 2 {
		t.Error("Dequeue가 올바른 순서로 값을 반환하지 않았습니다")
	}

	if q.Dequeue() != 3 {
		t.Error("Dequeue가 올바른 순서로 값을 반환하지 않았습니다")
	}

	// 빈 큐에서 Dequeue하면 -1 반환
	if q.Dequeue() != -1 {
		t.Error("빈 큐에서 Dequeue가 -1을 반환하지 않았습니다")
	}
}

func TestQueueFIFO(t *testing.T) {
	q := queue.NewQueue(10)

	// 20, 30, 40, 50 추가
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)

	// FIFO 순서 확인 (10, 20, 30, 40, 50)
	expected := []int{10, 20, 30, 40, 50}
	for i, exp := range expected {
		if val := q.Dequeue(); val != exp {
			t.Errorf("FIFO 순서가 맞지 않습니다. 인덱스 %d: 기대값 %d, 실제값 %d", i, exp, val)
		}
	}
}

func TestEmptyQueueDequeue(t *testing.T) {
	q := queue.NewQueue(5)

	// 모든 값을 Dequeue
	q.Dequeue()

	// 빈 큐에서 여러 번 Dequeue
	if q.Dequeue() != -1 {
		t.Error("빈 큐에서 Dequeue가 -1을 반환하지 않았습니다")
	}

	if q.Dequeue() != -1 {
		t.Error("빈 큐에서 연속 Dequeue가 -1을 반환하지 않았습니다")
	}
}

func TestQueueReuse(t *testing.T) {
	q := queue.NewQueue(1)

	// 값 추가 및 제거
	q.Enqueue(2)
	q.Dequeue() // 1
	q.Dequeue() // 2

	// 큐가 비워진 후 다시 사용
	q.Enqueue(3)
	q.Enqueue(4)

	if q.Dequeue() != 3 {
		t.Error("큐를 비운 후 재사용이 올바르게 작동하지 않았습니다")
	}

	if q.Dequeue() != 4 {
		t.Error("큐를 비운 후 재사용이 올바르게 작동하지 않았습니다")
	}
}

func TestQueueEnqueueFromEmpty(t *testing.T) {
	// 빈 큐 생성 (front와 back이 nil인 상태)
	q := &queue.Queue{}

	// 빈 큐에서 Enqueue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// FIFO 순서 확인
	if q.Dequeue() != 1 {
		t.Error("빈 큐에서 Enqueue 후 Dequeue가 올바르게 작동하지 않았습니다")
	}

	if q.Dequeue() != 2 {
		t.Error("Dequeue가 올바른 순서로 값을 반환하지 않았습니다")
	}

	if q.Dequeue() != 3 {
		t.Error("Dequeue가 올바른 순서로 값을 반환하지 않았습니다")
	}

	if q.Dequeue() != -1 {
		t.Error("모든 값을 Dequeue한 후 -1을 반환하지 않았습니다")
	}
}

func TestQueueLargeData(t *testing.T) {
	q := queue.NewQueue(1)

	// 많은 양의 데이터 추가
	for i := 2; i <= 1000; i++ {
		q.Enqueue(i)
	}

	// FIFO 순서로 확인
	for i := 1; i <= 1000; i++ {
		if val := q.Dequeue(); val != i {
			t.Errorf("큰 데이터셋에서 FIFO 순서가 맞지 않습니다. 기대값 %d, 실제값 %d", i, val)
			break
		}
	}
}
