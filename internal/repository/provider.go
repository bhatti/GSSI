package repository

// Provider represents a storage provider.
type Provider interface {
	// OpenStore opens a Store with the given name and returns it.
	// Depending on the store implementation, this may or may not create an underlying database.
	// The store implementation may defer creating the underlying database until SetStoreConfig is called or
	// data is inserted using Store.Put or Store.Batch.
	// Store names are not case-sensitive. If name is blank, then an error will be returned.
	OpenStore(name string) (Store, error)

	// SetStoreConfig sets the configuration on a Store. It's recommended calling this method at some point before
	// calling Store.Query if your store contains a large amount of data. The underlying database will use this to
	// create indexes to make querying via the Store.Query method faster. If you don't need to use Store.Query, then
	// you don't need to call this method. OpenStore must be called first before calling this method. If not, then an
	// error wrapping ErrStoreNotFound will be returned. This method will not open the store automatically.
	// If name is blank, then an error will be returned.
	SetStoreConfig(name string, config StoreConfiguration) error

	// GetStoreConfig gets the current Store configuration.
	// This method operates a bit differently in that it directly checks the underlying storage implementation to see
	// if the underlying database exists for the given name, rather than checking the currently known list of
	// open stores in memory. If no underlying database can be found, then an error wrapping ErrStoreNotFound will be
	// returned. This means that this method can be used to determine whether an underlying database for a Store
	// already exists or not. This method will not create the database automatically.
	// If name is blank, then an error will be returned.
	// As of writing, aries-framework-go code does not use this, but it may be useful for custom solutions.
	GetStoreConfig(name string) (StoreConfiguration, error)

	// GetOpenStores returns all Stores that are currently open in memory from calling OpenStore.
	// It does not check for all databases that have been created before. They have to have been opened in this Provider
	// object's lifetime from a call to OpenStore.
	// As of writing, aries-framework-go code does not use this, but it may be useful for custom solutions.
	GetOpenStores() []Store

	// Close closes all open Stores in this Provider
	// For persistent Store implementations, this does not delete any data in the underlying databases.
	Close() error
}