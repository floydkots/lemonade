package models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
)

type Member struct {
	email string
	id int
	password string
	firstName string
}

func (this *Member) Email() string {
	return this.email
}

func (this *Member) Id() int {
	return this.id
}

func (this *Member) Password() string {
	return this.password
}

func (this *Member) FirstName() string {
	return this.firstName
}

func (this *Member) SetEmail(value string) {
	this.email = value
}

func (this *Member) SetPassword(value string) {
	this.password = value
}

func (this *Member) SetFirstName(value string) {
	this.firstName = value
}


func CreateMember(email string, password string) (Member, error) {
	result := Member{}
	hashedPassword := sha256.Sum256([]byte(password))

	result.SetEmail(email)
	result.SetPassword(hex.EncodeToString(hashedPassword[:]))

	db, err := getDBConnetion()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO member
(email, password, first_name)
VALUES ($1, $2, $3)
RETURNING id`, result.email[:], result.password[:], result.email[:]).Scan(&result.id)
		if err == nil {
			return result, nil
		} else {
			panic(err)
			return Member{}, errors.New("unable to create new Member")
		}
	} else {
		panic(err)
		return result, errors.New("unable to get database connection")
	}
}


func GetMember(email string, password string) (Member, error) {
	db, err := getDBConnetion()

	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		row := db.QueryRow(`SELECT id, email, first_name
FROM member
WHERE email = $1 AND password = $2`, email, hex.EncodeToString(pwd[:]))
		result := Member{}
		err = row.Scan(&result.id, &result.email, &result.firstName)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("unable to find Member with email: " + email)
		}
	} else {
		return Member{}, errors.New("unable to get database connection")
	}
}


type Session struct {
	id int
	memberId int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}

func (this *Session) MemberId() int {
	return this.memberId
}

func (this *Session) SessionId() string {
	return this.sessionId
}

func (this *Session) SetId(value int) {
	this.id = value
}

func (this *Session) SetMemberId(value int) {
	this.memberId = value
}

func (this *Session) SetSessionId(value string) {
	this.sessionId = value
}


func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])

	db, err := getDBConnetion()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO session
(member_id, session_id)
VALUES ($1, $2)
RETURNING id`, member.Id(), result.sessionId).Scan(&result.id)
		if err == nil {
			return result, nil
		} else {
			return Session{}, errors.New("unable to save session to database")
		}
	} else {
		return result, errors.New("unable to get database connection")
	}
}
