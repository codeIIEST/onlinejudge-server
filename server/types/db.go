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
func (c *Contest) UpdateContest(db *pg.DB, cid string) error {
	_, err := db.Model(c).Table("contests").Where("contests.id = ?", cid).Update()
	return err
}

// CheckContestExists Checks if contests exists
func (c *Contest) CheckContestExists(db *pg.DB, cid string) (*Contest, bool, error) {
	contest := new(Contest)
	err := db.Model(contest).Table("contests").Where("contests.id = ?", cid).Limit(1).Select()
	if err != nil {
		return contest, false, err
	}
	return contest, true, nil

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
func (p *Problem) GetProblem(db *pg.DB) (*Problem, bool, error) {
	problem := &Problem{
		Index: p.Index,
	}

	err := db.Model(problem).Table("problems").Where("problems.index = ?", p.Index).Limit(1).Select()
	if err != nil {
		_, ok := err.(pg.Error)
		if !ok {
			return problem, true, err
		}
	}
	return problem, false, nil
}

// GetAllProblems gets all problems from db
func (p *Problem) GetAllProblems(db *pg.DB, cid string) (*[]Problem, error) {
	var problem []Problem
	err := db.Model(&problem).Table("problems").Where("problems.contest_id = ?", cid).Distinct().Select()
	return &problem, err
}
