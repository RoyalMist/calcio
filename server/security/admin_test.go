package security

/*func TestCreateAdmin(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	_ = client.Schema.Create(context.Background())
	defer func(client *ent.Client) {
		_ = client.Close()
	}(client)

	err := CreateAdmin(client)
	if err != nil {
		t.Errorf("should not have raised an error but got %v", err)
	}

	admin, err := client.User.Query().Where(user.And(user.Admin(true), user.Name("Admin"))).First(context.Background())
	if err != nil {
		t.Errorf("should not have raised an error but got %v", err)
	}

	if admin == nil {
		t.Errorf("should have returned an admin user")
	}
}*/
