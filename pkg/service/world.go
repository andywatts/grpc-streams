package service

type World struct {
	domainName DomainName
	users      []*User
}

var (
	worlds map[DomainName]*World
)

func init() {
	worlds = make(map[DomainName]*World)
}

func FindWorld(domainName DomainName) *World {
	if _, ok := worlds[domainName]; !ok {
		world := World{domainName: domainName}
		worlds[domainName] = &world
	}
	return worlds[domainName]
}

func (w *World) Subscribe(user *User) {
	w.users = append(w.users, user)
}

func (w *World) Unsubscribe(user *User) {
	//w.users = remove(w.users, user) // Generics version
	w.users = removeUser(w.users, user)

	// remove world if last user
	if len(w.users) == 0 {
		delete(worlds, w.domainName)
	}
}

// TODO remove and use generics utils
func removeUser(s []*User, r *User) []*User {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
