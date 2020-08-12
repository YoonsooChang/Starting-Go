package mydict

import "errors"

/* alias (typedef) */
// Dictionary Type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errAlreadyExists = errors.New("Already Exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a new word
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errAlreadyExists
	}
	return nil
}

// Update definition of a word
func (d Dictionary) Update(word, def string) (string, error) {
	_, err := d.Search(word)
	if err == errNotFound {
		return def, d.Add(word, def)
	}
	d[word] = def
	return def, nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		return err
	}
	delete(d, word)
	return nil
}
