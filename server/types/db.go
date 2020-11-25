package types

import (
	"log"

	pg "github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

// Create create a new user and insert into db
func (u *User) Create(db *pg.DB) error {
	_, err := db.Model(u).Insert()
	return err
}

// Update update existing user
func (u *User) Update(db *pg.DB) error {
	_, err := db.Model(u).Update()
	return err
}

// CheckUserExists to check if user is present
func (u *User) CheckUserExists(db *pg.DB) (bool, error) {
	user := new(User)
	err := db.Model(user).Table("users").Where("users.email = ?", u.Email).Limit(1).Select()
	encryptionErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))

	if encryptionErr == nil {
		u.FirstName = user.FirstName
		u.LastName = user.LastName
		return true, nil
	}
	if err != nil {
		log.Println(err)
	}
	return false, err
}

//Delete delete user
func (u *User) Delete(db *pg.DB) error {
	_, err := db.Model(u).Delete()
	return err
}

// ContestCreate Create Contest
func (c *Contest) ContestCreate(db *pg.DB) error {
	_, err := db.Model(c).Insert()
	return err
}

// UpdateContest updates contest
func (c *Contest) UpdateContest(db *pg.DB) error {
	_, err := db.Model(c).Table("contests").WherePK().Update()
	return err
}

// CheckContestExists Checks if contests exists
func (c *Contest) CheckContestExists(db *pg.DB) (bool, error) {
	_, err := db.Query(c, `SELECT * FROM contests WHERE id = ?`, c.ID)
	if err != nil {
		return false, err
	}
	return true, nil

}

// CreateProblem creates problem
func (p *Problem) CreateProblem(db *pg.DB) (bool, error) {
	_, err := db.Model(p).Insert()
	if err != nil {
		pgErr, ok := err.(pg.Error)
		if ok && pgErr.IntegrityViolation() {
			return true, err
		} else if pgErr.Field('S') == "PANIC" {
			return false, err
		}
	}
	return false, nil
}

// GetProblem gets the problem
func (p *Problem) GetProblem(db *pg.DB) (bool, error) {
	_, err := db.Query(p, `SELECT * FROM problems WHERE index = ?`, p.Index)
	if err != nil {
		_, ok := err.(pg.Error)
		if !ok {
			return true, err
		}
	}
	if len(p.ID) == 0 {
		return true, nil
	}
	return false, nil
}

// GetAllProblems gets all problems from db
func (p *Problem) GetAllProblems(db *pg.DB, cid string) (*[]Problem, error) {
	var problem []Problem
	_, err := db.Query(&problem, `SELECT * FROM problems WHERE contest_id = ?`, cid)
	return &problem, err
}

// GetAllContests gets all contests from DB
func (c *Contest) GetAllContests(db *pg.DB) (*[]Contest, error) {
	var contests []Contest
	_, err := db.Query(&contests, `SELECT * FROM contests`)
	return &contests, err
}
