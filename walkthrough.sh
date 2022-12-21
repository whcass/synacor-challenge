#!/usr/bin/env expect -f
spawn "./synacor-challenge"
# cave
expect "do\?"
send "take tablet\r"
send "doorway\r"
send "north\r"
send "north\r"
send "bridge\r"
send "continue\r"
send "down\r"
send "east\r"
send "take empty lantern\r"
send "west\r"
send "west\r"
send "passage\r"
send "ladder\r"
send "west\r"
send "south\r"
send "north\r"
send "take can\r"
send "use can\r"
send "use lantern\r"
send "west\r"
send "ladder\r"
send "darkness\r"
send "continue\r"
send "west\r"
send "west\r"
send "west\r"
send "west\r"
send "north\r"
send "take red coin\r"
send "north\r"
# Ruins
# 2 7 3 9 5
# red:        2
# concave:    7
# corroded:   3
# blue:       9
# shiny:      5
# answer: 9 2 5 7 3
send "east\r"
send "take concave coin\r"
send "down\r"
send "take corroded coin\r" 
send "up\r"
send "west\r"
send "west\r"
send "take blue coin\r"
send "up\r"
send "take shiny coin\r"
send "down\r"
send "east\r"
send "use blue coin\r"
send "use red coin\r"
send "use shiny coin\r"
send "use concave coin\r"
send "use corroded coin\r"
send "north\r"
send "take teleporter\r"
# send "use teleporter\r"


interact