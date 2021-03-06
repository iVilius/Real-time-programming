// Channel definitions for elevator control using LibComedi
//
// 2006, Martin Korsgaard
// 2015, VilKri

package driver

const N_elevators 		= 2
const M_floors 			= 4

var Init_port  string	= "19174"
var Alive_port string 	= "47143"
var Order_port string	= "13340"
var State_port string 	= "57473"

type Message struct {

    ID 					int // 0 for init, 1 for <I am alive>, 9000 for error, 666 to escape current go routine
    Latest_floor 		int // 1 for first, 2 for second and so on
    Direction 			int	// 0 for NO_DIRECTION, 1 for UP and 2 for DOWN
    Local_IP 			string
    Remote_IP 			string 
    Trunc_IP			int
    Order_elevator_ID   int
    Order_type          string
    State_matrix	[][]int
}

const	( 	
	INIT 				= 0
	ALIVE 				= 1
	STATE				= 2
	ORDER_NEW 			= 3
	ORDER_NEW_ACK 		= 4
	ORDER_ASSIGN		= 5
	ORDER_ASSIGN_ACK	= 6
	ORDER_DONE			= 7
	ORDER_DONE_ACK		= 8
	STATE_MATRIX_UPDATE	= 9
	)


//in port 4
const (
	PORT4 				= 3
	OBSTRUCTION 		= 0x300+23
	STOP 				= 0x300+22
	BUTTON_COMMAND1 	= 0x300+21
	BUTTON_COMMAND2 	= 0x300+20
	BUTTON_COMMAND3 	= 0x300+19
	BUTTON_COMMAND4 	= 0x300+18
	BUTTON_UP1			= 0x300+17
	BUTTON_UP2  		= 0x300+16
)
//out port 3
const (
	PORT3 			= 3
	MOTORDIR 		= 0x300+15
	LIGHT_STOP 		= 0x300+14
	ORDER1 			= 0x300+13
 	ORDER2 			= 0x300+12
 	ORDER3 			= 0x300+11
 	ORDER4 			= 0x300+10
 	UP1 			= 0x300+9
 	UP2 			= 0x300+8
)
//out port 2
const (
	PORT2 				= 3
	DOWN2 				= 0x300+7
	UP3 				= 0x300+6
	DOWN3 				= 0x300+5
	DOWN4 				= 0x300+4
	LIGHT_DOOR_OPEN 	= 0x300+3
	LIGHT_FLOOR_IND2	= 0x300+1
	LIGHT_FLOOR_IND1	= 0x300+0
)
//in port 1
const (
	PORT1 				= 2
	BUTTON_DOWN2 		= 0x200+0
	BUTTON_UP3 			= 0x200+1
	BUTTON_DOWN3 		= 0x200+2
	BUTTON_DOWN4 		= 0x200+3
	SENSOR_FLOOR1 		= 0x200+4
	SENSOR_FLOOR2 		= 0x200+5
	SENSOR_FLOOR3 		= 0x200+6
	SENSOR_FLOOR4 		= 0x200+7
)

//out port 0
const (
	PORT0 			= 1
	MOTOR 			= 0x100+0
)

//non-existing ports (for alignment)
const (
	
	BUTTON_DOWN1 		= -1
	BUTTON_UP4 		= -1
	LIGHT_DOWN1 		= -1
	LIGHT_UP4 		= -1
)

