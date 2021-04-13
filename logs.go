package everquest

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

type LogAlert interface {
	Check(string) // If Regex string input matches, then act
}

const EQBaseLogLine = "\\[(\\w{3} \\w{3} \\d{2} \\d{2}:\\d{2}:\\d{2} \\d{4})] (.+)"

func BufferedLogRead(path string, fromStart bool, pollRate int, out chan EqLog) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening buffered file: %v", err)
		return
	}
	if !fromStart {
		file.Seek(0, 2) // move to end of file
	}
	bufferedReader := bufio.NewReader(file)
	r, _ := regexp.Compile(EQBaseLogLine)
	for {
		str, err := bufferedReader.ReadString('\n')
		if err == io.EOF {
			time.Sleep(time.Duration(pollRate) * time.Second) // 1 eq tick = 6 seconds
			continue
		}
		if err != nil {
			log.Printf("error opening buffered file: %v", err)
			return
		}

		results := r.FindAllStringSubmatch(str, -1) // this really needs converted to single search
		if results == nil {
			time.Sleep(3 * time.Second)
		} else {
			log := readLogLine(results)
			out <- *log
		}
	}
}

func readLogLine(results [][]string) *EqLog {
	t := eqTimeConv(results[0][1])
	msg := strings.TrimSuffix(results[0][2], "\r")
	log := &EqLog{
		T:       t,
		Msg:     msg,
		Channel: getChannel(msg),
		Source:  getSource(msg),
	}
	return log
}

func eqTimeConv(t string) time.Time {
	// Get local time zone
	localT := time.Now()
	zone, _ := localT.Zone()

	// Parse Time
	cTime, err := time.Parse("Mon Jan 02 15:04:05 2006 MST", t+" "+zone)
	if err != nil {
		fmt.Printf("Error parsing time, defaulting to now: %s\n", err.Error())
		cTime = time.Now()
	}
	return cTime
}

// EqLog represents a single line of eq logging
type EqLog struct {
	T       time.Time `json:"Time"`
	Msg     string    `json:"Message"`
	Channel string    `json:"Channel"`
	Source  string    `json:"Source"`
}

func getChannel(msg string) string {
	m := strings.Split(msg, " ")
	if len(m) > 4 {
		if m[3] == "guild," || m[4] == "guild," {
			return "guild"
		}
		if m[3] == "group," || m[4] == "group," {
			return "group"
		}
		if m[3] == "raid," || m[4] == "raid," {
			return "raid"
		}
		if m[1] == "tells" && m[2] == "you," {
			return "tell"
		}
		// fmt.Printf("Default: %s\n", m[2])
		return "system"
	}
	if len(m) > 1 && m[1] == "tells" {
		// return m[3]
		return "tell"
		// return m[0] // source should show the player not the channel
		// return strings.TrimRight(m[3], ",")
	}
	if len(m) > 1 && m[1] == "auctions," {
		return "auction"
	}
	if len(m) > 1 && m[1] == "says," {
		return "say"
	}
	return "system"
}

func getSource(msg string) string {
	m := strings.Split(msg, " ")
	return m[0]
}

func GetLogPath(player, server, basePath string) string {
	server = strings.ToLower(server) // servernames are lowercase
	player = strings.Title(player)   // first letter of player is uppercase
	return basePath + "\\Logs\\eqlog_" + player + "_" + server + ".txt"
}
