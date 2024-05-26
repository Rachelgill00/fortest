package config

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"forTest/identity"

	"github.com/gitferry/bamboo/transport"
)

var configFile = flag.String("config", "config.json", "Configuration file for bamboo replica. Defaults to config.json.")

// Config contains every system configuration
type Config struct {
	Addrs     map[identity.NodeID]string `json:"address"`      // address for node communication
	HTTPAddrs map[identity.NodeID]string `json:"http_address"` // address for client server communication

	Policy    string  `json:"policy"`    // leader change policy {consecutive, majority}
	Threshold float64 `json:"threshold"` // threshold for policy in WPaxos {n consecutive or time interval in ms}

	Thrifty        bool `json:"thrifty"`          // only send messages to a quorum
	BufferSize     int  `json:"buffer_size"`      // buffer size for maps
	ChanBufferSize int  `json:"chan_buffer_size"` // buffer size for channels
	MultiVersion   bool `json:"multiversion"`     // create multi-version database
	Timeout        int  `json:"timeout"`
	ByzNo          int  `json:"byzNo"`
	BSize          int  `json:"bsize"`
	Fixed          bool `json:"fixed"`
	//Benchmark      Bconfig         `json:"benchmark"` // benchmark configuration
	Delta       int             `json:"delta"` // timeout, seconds
	Pprof       bool            `json:"pprof"`
	MaxRound    int             `json:"maxRound"`
	Strategy    string          `json:"strategy"`
	PayloadSize int             `json:"payload_size"`
	Master      identity.NodeID `json:"master"`
	Delay       int             `json:"delay"` // transmission delay in ms
	DErr        int             `json:"derr"`  // the err taken into delays
	MemSize     int             `json:"memsize"`
	Slow        int             `json:"slow"`
	Crash       int             `json:"crash"`

	hasher string
	signer string

	// for future implementation
	// Batching bool `json:"batching"`
	// Consistency string `json:"consistency"`
	// Codec string `json:"codec"` // codec for message serialization between nodes

	n int // total number of nodes
	//z   int         // total number of zones
	//npz map[int]int // nodes per zone
}

// Bconfig holds all benchmark configuration
type Bconfig struct {
	T            int    // total number of running time in seconds
	N            int    // total number of requests
	K            int    // key sapce
	Throttle     int    // requests per second throttle, unused if 0
	Concurrency  int    // number of simulated clients
	Distribution string // distribution
	// rounds       int    // repeat in many rounds sequentially

	// conflict distribution
	Conflicts int // percentage of conflicting keys
	Min       int // min key

	// normal distribution
	Mu    float64 // mu of normal distribution
	Sigma float64 // sigma of normal distribution
	Move  bool    // moving average (mu) of normal distribution
	Speed int     // moving speed in milliseconds intervals per key
}

// Config is global configuration singleton generated by init() func below
var Configuration Config

func init() {
	Configuration = MakeDefaultConfig()
}

// MakeDefaultConfig returns Config object with few default values
// only used by init() and master
func MakeDefaultConfig() Config {
	return Config{
		Policy:         "consecutive",
		Threshold:      3,
		BufferSize:     1024,
		ChanBufferSize: 1024,
		MultiVersion:   false,
		hasher:         "sha3_256",
		signer:         "ECDSA_P256",
		//Benchmark:      DefaultBConfig(),
	}
}

func init() {
	Configuration = MakeDefaultConfig()
}

// GetConfig returns paxi package configuration
func GetConfig() Config {
	return Configuration
}

func GetTimer() time.Duration {
	return time.Duration(time.Duration(Configuration.Timeout) * time.Millisecond)
}

// Simulation enable go channel transportation to simulate distributed environment
func Simulation() {
	*transport.Scheme = "chan"
}

// Load loads configuration from Configuration file in JSON format
func (c *Config) Load() {
	file, err := os.Open(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		log.Fatal(err)
	}

	// load ips
	ip_file, err := os.Open("ips.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer ip_file.Close()

	scanner := bufio.NewScanner(ip_file)
	i := 1
	for scanner.Scan() {
		id := identity.NewNodeID(i)
		port := strconv.Itoa(3734 + i)
		addr := "tcp://" + scanner.Text() + ":" + port
		portHttp := strconv.Itoa(8069 + i)
		addrHttp := "http://" + scanner.Text() + ":" + portHttp
		c.Addrs[id] = addr
		c.HTTPAddrs[id] = addrHttp
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	c.n = len(c.Addrs)
}
