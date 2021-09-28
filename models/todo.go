package models

import "cloud.google.com/go/firestore"

type Todo struct {
	Name   *string `firestore:"name"`
	IsDone *bool   `firestore:"isDone"`
}

func (todo *Todo) ToFirestoreUpdate() []firestore.Update {

	var updateFields []firestore.Update

	if todo.Name != nil {
		updateFields = append(updateFields, firestore.Update{
			FieldPath: firestore.FieldPath{"name"},
			Value:     todo.Name,
		})
	}

	if todo.IsDone != nil {
		updateFields = append(updateFields, firestore.Update{
			FieldPath: firestore.FieldPath{"name"},
			Value:     todo.Name,
		})
	}

	return updateFields
}