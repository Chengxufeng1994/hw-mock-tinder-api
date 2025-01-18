package ddd

type Snapshot interface {
	SnapshotName() string
}

type SnapshotApplier interface {
	ApplySnapshot(snapshot Snapshot) error
}

type Snapshotter interface {
	SnapshotApplier
	ToSnapshot()
}
