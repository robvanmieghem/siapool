package sharechain

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/NebulousLabs/Sia/persist"

	"github.com/NebulousLabs/bolt"
)

const (
	// DatabaseFilename contains the filename of the database that will be used
	DatabaseFilename = "sharechain.db"
	logFile          = "sharechain.log"
)

var (
	errRepeatInsert = errors.New("attempting to add an already existing item to the sharechain set")
	errNilItem      = errors.New("requested item does not exist")

	dbMetadata = persist.Metadata{
		Header:  "Consensus Set Database",
		Version: "0.5.0",
	}

	// ShareChainPool is a database bucket storing the current value of the
	// ShareChain pool.
	ShareChainPool = []byte("ShareChainPool")
)

// createShareChainDB initialzes the sharechain portions of the database.
func (sc *ShareChain) createShareChainDB(tx *bolt.Tx) error {
	// Enumerate and create the database buckets.
	buckets := [][]byte{
		ShareChainPool,
	}
	for _, bucket := range buckets {
		_, err := tx.CreateBucket(bucket)
		if err != nil {
			return err
		}
	}

	return nil
}

// loadDB pulls all the shares that have been saved to disk into memory, using
// them to fill out the ShareChain.
func (sc *ShareChain) loadDB() error {
	// Open the database - a new bolt database will be created if none exists.
	err := sc.openDB(filepath.Join(sc.persistDir, DatabaseFilename))
	if err != nil {
		return err
	}

	// Walk through initialization for Sia.
	return sc.db.Update(func(tx *bolt.Tx) error {
		// Check if the database has been initialized.
		if !dbInitialized(tx) {
			return sc.initDB(tx)
		}

		return nil
	})
}

// initPersist initializes the persistence structures of the sharechain.
func (sc *ShareChain) initPersist() error {
	// Create the sharechain directory.
	err := os.MkdirAll(sc.persistDir, 0700)
	if err != nil {
		return err
	}

	// Initialize the logger.
	sc.log, err = persist.NewFileLogger(filepath.Join(sc.persistDir, logFile))
	if err != nil {
		return err
	}

	// Try to load an existing database from disk - a new one will be created
	// if one does not exist.
	err = sc.loadDB()
	if err != nil {
		return err
	}
	return nil
}

// openDB loads the set database and populates it with the necessary buckets
func (sc *ShareChain) openDB(filename string) (err error) {
	sc.db, err = persist.OpenDatabase(dbMetadata, filename)
	if err == persist.ErrBadVersion {
		return sc.replaceDatabase(filename)
	}
	if err != nil {
		return errors.New("error opening sharechain database: " + err.Error())
	}
	return nil
}

// dbInitialized returns true if the database appears to be initialized, false
// if not. Checking for the existence of the ShareChain pool bucket is typically
// sufficient to determine whether the database has gone through the
// initialization process.
func dbInitialized(tx *bolt.Tx) bool {
	return tx.Bucket(ShareChainPool) != nil
}

// initDB is run if there is no existing sharechain database, creating a
// database with all the required buckets and sane initial values.
func (sc *ShareChain) initDB(tx *bolt.Tx) (err error) {
	// Create the compononents of the database.
	err = sc.createShareChainDB(tx)

	return
}

// replaceDatabase backs up the existing database and creates a new one.
func (sc *ShareChain) replaceDatabase(filename string) error {
	// Rename the existing database and create a new one.
	log.Println("Outdated sharechain database... backing up and replacing")
	err := os.Rename(filename, filename+".bck")
	if err != nil {
		return errors.New("error while backing up sharechain database: " + err.Error())
	}

	// Try again to create a new database, this time without checking for an
	// outdated database error.
	sc.db, err = persist.OpenDatabase(dbMetadata, filename)
	if err != nil {
		return errors.New("error opening sharechain database: " + err.Error())
	}
	return nil
}
