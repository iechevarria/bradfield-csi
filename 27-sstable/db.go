package crapdb

type DB interface {
	// Get gets the value for the given key. It returns an error if the
	// DB does not contain the key.
	Get(key []byte) (value []byte, err error)

	// Has returns true if the DB contains the given key.
	Has(key []byte) (ret bool, err error)

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a DB is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// RangeScan returns an Iterator (see below) for scanning through all
	// key-value pairs in the given range, ordered by key ascending.
	RangeScan(start, limit []byte) (Iterator, error)
}

type Iterator interface {
	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool

	// Error returns any accumulated error. Exhausting all the key/value pairs
	// is not considered to be an error.
	Error() error

	// Key returns the key of the current key/value pair, or nil if done.
	Key() []byte

	// Value returns the value of the current key/value pair, or nil if done.
	Value() []byte
}

type ImmutableDB interface {
    // Get gets the value for the given key. It returns an error if the
    // DB does not contain the key.
    Get(key []byte) (value []byte, err error)

    // Has returns true if the DB contains the given key.
    Has(key []byte) (ret bool, err error)

    // RangeScan returns an Iterator (see below) for scanning through all
    // key-value pairs in the given range, ordered by key ascending.
    RangeScan(start, limit []byte) (Iterator, error)
}

type CrapDB struct {
	Memtable *Memtable
}

func NewCrapDB(maxLevel int) *CrapDB {
	memtable := NewMemtable(maxLevel)
	return &CrapDB{Memtable: memtable}
}

func (db *CrapDB) Get(key []byte) ([]byte, error) {
	return db.Memtable.Get(key)
}

func (db *CrapDB) Has(key []byte) (bool, error) {
	return db.Memtable.Has(key)
}

func (db *CrapDB) Delete(key []byte) error {
	return db.Memtable.Delete(key)
}

func (db *CrapDB) Put(key []byte, value []byte) error {
	return db.Memtable.Put(key, value)
}

func (db *CrapDB) RangeScan(start, limit []byte) (Iterator, error) {
	return db.Memtable.RangeScan(start, limit)
}

func (db CrapDB) PrettyPrint() {
	db.Memtable.PrettyPrint()
}
