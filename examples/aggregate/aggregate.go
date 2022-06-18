package aggregate

import (
	"fmt"

	"github.com/wlevene/mgm/v3"
	"github.com/wlevene/mgm/v3/builder"
	"github.com/wlevene/mgm/v3/field"
	"github.com/wlevene/mgm/v3/operator"
	"go.mongodb.org/mongo-driver/bson"
)

func seed() {

	fmt.Println("seeding data...")
	author := newAuthor("Mehran")
	_ = mgm.Coll(author).Create(author)

	book := newBook("Test", 124, author.ID)
	_ = mgm.Coll(book).Create(book)

	book1 := newBook("Test1", 124, author.ID)
	_ = mgm.Coll(book1).Create(book1)

}

func delSeededData() {
	_, _ = mgm.Coll(&book{}).DeleteMany(nil, bson.M{})
	_, _ = mgm.Coll(&author{}).DeleteMany(nil, bson.M{})
}

func lookup() error {
	seed()

	defer delSeededData()

	// Author model's collection
	authorColl := mgm.Coll(&author{})

	pipeline := bson.A{
		builder.S(builder.Lookup(authorColl.Name(), "author_id", field.ID, "author")),
		bson.M{operator.Project: bson.M{"name": "Test1"}},
	}

	cur, err := mgm.Coll(&book{}).Aggregate(mgm.Ctx(), pipeline)

	if err != nil {
		return err
	}

	defer cur.Close(nil)

	for cur.Next(nil) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return err
		}

		// do something with result....
		fmt.Printf("%+v\n", result)
	}

	return nil
}
