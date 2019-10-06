package Database

type Boiler struct {
	ID          int
	DesiredTemp int
	MaxTemp     int
}

type User struct {
	UserName string
	Pass     string
	BoilerID int
}

var UserMap = make(map[string]User)
var BoilerMap = make(map[int]Boiler)

func AddNewBoiler(boilerId int, desiredTemp int, maxTemp int){
	newBoiler := Boiler{boilerId, desiredTemp, maxTemp}
	BoilerMap[boilerId] = newBoiler
}

func AddNewUser(user User, desiredTemp int, maxTemp int) {
	UserMap[user.UserName] = user
	boiler, ok := BoilerMap[user.BoilerID]
	if !ok {
		AddNewBoiler(user.BoilerID, desiredTemp, maxTemp)
	} else {
		boiler.DesiredTemp = desiredTemp
		boiler.MaxTemp = maxTemp
	}
}

func GetBoilerByUserName(userName string, pass string) *Boiler {

	user, validUser := UserMap[userName]
	if validUser {
		if user.Pass == pass {
			boiler, validBoiler := BoilerMap[user.BoilerID]

			if validBoiler {
				return &boiler
			}
		}
	}
	return nil
}

func UserExists(userName string) bool {
	_, validUser := UserMap[userName]
	if validUser {
		return true
	}
	return false
}
