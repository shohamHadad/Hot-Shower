package HandleClient

import (
	"HotShower/Database"
	"fmt"
	"math/rand"
	"net"
	"strconv"
)

const (
	OFF    		int = 0
	ON    		int = 1
)

func HandleNewUser(conn net.Conn, userName string, pass string){
	if !Database.UserExists(userName){

		boilerId := ReadNumberFromClient(conn)
		desiredTemp := ReadNumberFromClient(conn)
		maxTemp := ReadNumberFromClient(conn)

		newUser := Database.User{userName, pass, boilerId}
		Database.AddNewUser(newUser, desiredTemp, maxTemp)
	}
}

func SetConfiguration (conn net.Conn, userName string, pass string){
	if !Database.UserExists(userName){
		fmt.Println("User does not exist!")
		conn.Close()
	}
	desiredTemp := ReadNumberFromClient(conn)
	maxTemp := ReadNumberFromClient(conn)

	boiler := Database.GetBoilerByUserName(string(userName), string(pass))
	boiler.MaxTemp = maxTemp
	boiler.DesiredTemp = desiredTemp

	/**
	now need to open connection with boiler and send new configurations. Didn't implement in this assignment.
	 */
}

func ControlBoilerConditions (conn net.Conn, userName string, pass string){
	if !Database.UserExists(userName){
		fmt.Println("User does not exist!")
		conn.Close()
	}
	newState := ReadNumberFromClient(conn)
	boiler := Database.GetBoilerByUserName(string(userName), string(pass))

	// only to prevent compilation errors
	fmt.Println(newState)
	fmt.Println(boiler)

	/**
	now need to open connection with boiler and send new state (on\off). Didn't implement in this assignment.
	 */
}

func ReadNumberFromClient(conn net.Conn) int{
	buff := make([]byte, 64)

	// Get User name
	_ , err := conn.Read(buff)
	if err != nil {
	CloseConnection(conn, "Error reading from user:", err)
	return -1
	}

	res, err := strconv.Atoi(string(buff))

	if err != nil{
		CloseConnection(conn, "Error reading from user:", err)
		return -1
	}
	return res
}

func GetHotWaterInMinutes(conn net.Conn, userName string, password string){
	/**
	Assume that boiler thermostat is able to send number of liters of hot water,
	and the amount of water flowing through the faucet.
	The server will open a connection with the thermostat app, receive the needed information
	and send the result after calculations.
 	 */

 	 res := rand.Int()
 	 resBuff := []byte(strconv.Itoa(res))
 	 _, err := conn.Write(resBuff)
 	 if err != nil {
		fmt.Println("Writing failed", err)
		return
	}
}