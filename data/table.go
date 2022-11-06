package data

import "sync"

type Table struct {
	mutex sync.RWMutex
	rows  []Row
}

func (t *Table) BytesSize() int {
	t.mutex.RLock()

	result := Int64Size // 8 bytes for rows count value
	for _, row := range t.rows {
		result += Int64Size       // 8 bytes for row size value
		result += row.BytesSize() // Row actual size in bytes
	}

	t.mutex.RUnlock()

	return result
}

func NewTable() Table {
	return Table{
		mutex: sync.RWMutex{},
		rows:  make([]Row, 0),
	}
}
