package main

import "errors"

type Dictionary map[string]string

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search returns the definition of the given word
// It returns an error if the word does not exist in dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds a new word with its defination to the dictionary
// If the word already exists, it returns ErrWordExists and does not
// modify the existing defination.
func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = defination
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
