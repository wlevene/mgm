package crud

import "github.com/wlevene/mgm/v3"

func crud() error {

	book := newBook("Test", 124)
	booksColl := mgm.Coll(book)

	if err := booksColl.Create(book); err != nil {
		return err
	}

	book.Name = "Moulin Rouge!"
	if err := booksColl.Update(book); err != nil {
		return err
	}

	return nil
	return booksColl.Delete(book)
}
