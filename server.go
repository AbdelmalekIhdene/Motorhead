package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gorilla/mux"
)

type motorcycle struct {
	Name               string
	CurbWeight         int32
	Details            string
	EngineDisplacement int32
	EngineType         string
	Manufacturer       string
	TransmissionType   string
	RetailPrice        int32
	ReviewLink         string
	HeroImagePath      string
	SideviewImagePath  string
}

type server struct {
	*http.Server
	motorcycles map[string]motorcycle
}

func NewServer() *server {
	srv := &server{
		Server: &http.Server{
			Addr:    ":8081",
			Handler: mux.NewRouter(),
		},
		motorcycles: map[string]motorcycle{
			"2019-harley-davidson-iron-883": motorcycle{
				Name:               "2019 Harley-Davidson Iron 883",
				Details:            "An original icon of the Harley-Davidson Dark Custom style. It sets the standard for the raw, stripped-down, blacked-out look with fresh new graphics",
				Manufacturer:       "Harley-Davidson",
				CurbWeight:         247,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 883,
				TransmissionType:   "5-Speed",
				RetailPrice:        11499,
				ReviewLink:         "https://www.youtube.com/embed/WpIo4NSzYVU",
			},
			"2019-harley-davidson-softail": motorcycle{
				Name:               "2019 Harley-Davidson Softail",
				Details:            "A stripped-down, souped-up, bobber with a post WWII era look. It's 16 kilograms lighter, with greater lean angle, than models from just a couple of years ago",
				Manufacturer:       "Harley-Davidson",
				CurbWeight:         291,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 1746,
				TransmissionType:   "6-Speed",
				RetailPrice:        18999,
				ReviewLink:         "https://www.youtube.com/embed/fPG_zykKvw4",
			},
			"2019-honda-rebel-300": motorcycle{
				Name:               "2019 Honda Rebel 300",
				Details:            "The combination of a fresh, new look with timeless features like a low seat height, light weight, narrow 286cc single-cylinder engine and user-friendly powerband",
				Manufacturer:       "Honda",
				CurbWeight:         168,
				EngineType:         "Single-Cylinder Engine",
				EngineDisplacement: 286,
				TransmissionType:   "6-Speed",
				RetailPrice:        4499,
				ReviewLink:         "https://www.youtube.com/embed/Z8bVMyPA3Xk",
			},
			"2019-honda-shadow-phantom": motorcycle{
				Name:               "2019 Honda Shadow Phantom",
				Details:            "The big brother of the Rebel range with a blacked-out 745CC V-Twin and a throaty exhaust that always announce your arrival",
				Manufacturer:       "Honda",
				CurbWeight:         249,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 745,
				TransmissionType:   "5-Speed",
				RetailPrice:        7899,
				ReviewLink:         "https://www.youtube.com/embed/F0ltDOqAlAM",
			},
			"2019-indian-scout-sixty": motorcycle{
				Name:               "2019 Indian Scout Sixty",
				Details:            "From its light and agile feel, to its low seat height and throwback style, the Scout Sixty is easy to ride and hard not to",
				Manufacturer:       "Indian",
				CurbWeight:         245,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 999,
				TransmissionType:   "6-Speed",
				RetailPrice:        8999,
				ReviewLink:         "https://www.youtube.com/embed/b9NvefV8yC8",
			},
			"2019-kawasaki-vulcan-900-classic": motorcycle{
				Name:               "2019 Kawasaki Vulcan 900 Classic",
				Details:            "A clean and timeless design from the 1980s in a powerful and comfortable machine that captures the spirit of the American cruiser",
				Manufacturer:       "Kawasaki",
				CurbWeight:         280,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 903,
				TransmissionType:   "5-Speed",
				RetailPrice:        7999,
				ReviewLink:         "https://www.youtube.com/embed/CpWZziNu5YU",
			},
			"2019-kawasaki-vulcan-s": motorcycle{
				Name:               "2019 Kawasaki Vulcan S",
				Details:            "A modern take on the classic formula that delivers exhilarating sport cruiser performance for maximum enjoyment on every ride",
				Manufacturer:       "Kawasaki",
				CurbWeight:         225,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 649,
				TransmissionType:   "6-Speed",
				RetailPrice:        7099,
				ReviewLink:         "https://www.youtube.com/embed/5v3UGuUPRFY",
			},
			"2019-moto-guzzi-roamer": motorcycle{
				Name:               "2019 Moto Guzzi Roamer",
				Details:            "Glamorous, easy, convenient and rewarding to ride, the Roamer is the Italian custom bike dedicated to the citizens of the world",
				Manufacturer:       "Guzzi",
				CurbWeight:         199,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 850,
				TransmissionType:   "6-Speed",
				RetailPrice:        11090,
				ReviewLink:         "https://www.youtube.com/embed/xa-3u0Qh4rA",
			},
			"2019-suzuki-boulevard-s40": motorcycle{
				Name:               "2019 Suzuki Boulevard S40",
				Details:            "A timeless single cylinder design that combines exciting performance and a bold appearance with rock-solid Suzuki reliability",
				Manufacturer:       "Suzuki",
				CurbWeight:         172,
				EngineType:         "Single-Cylinder Engine",
				EngineDisplacement: 652,
				TransmissionType:   "5-Speed",
				RetailPrice:        5799,
				ReviewLink:         "https://www.youtube.com/embed/w28iZahu9WY",
			},
			"2019-triumph-bonneville-bobber": motorcycle{
				Name:               "2019 Triumph Bonneville Bobber",
				Details:            "A design that remains true to the original’s cut-down 1950s England styling, adding other features to improve handling and overall performance",
				Manufacturer:       "Triumph",
				CurbWeight:         236,
				EngineType:         "Parallel-Twin Engine",
				EngineDisplacement: 1197,
				TransmissionType:   "6-Speed",
				RetailPrice:        13150,
				ReviewLink:         "https://www.youtube.com/embed/WdvsuNQ5924",
			},
			"2019-yamaha-bolt": motorcycle{
				Name:               "2019 Yamaha Bolt",
				Details:            "Powered by a V-twin engine with a plenty of technology on board, the Bolt is not too small that you’ll outgrow it soon and not so big that it is intimidating for new riders",
				Manufacturer:       "Yamaha",
				CurbWeight:         248,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 942,
				TransmissionType:   "5-Speed",
				RetailPrice:        9499,
				ReviewLink:         "https://www.youtube.com/embed/prrzUhe9mI8",
			},
			"2019-yamaha-v-star-250": motorcycle{
				Name:               "2019 Yamaha V-Star 250",
				Details:            "The V-Star 250 welcomes novice riders with manageable power, a smooth-shifting 5-speed transmission, light clutch pull and a low seat height",
				Manufacturer:       "Yamaha",
				CurbWeight:         147,
				EngineType:         "V-Twin Engine",
				EngineDisplacement: 249,
				TransmissionType:   "5-Speed",
				RetailPrice:        4060,
				ReviewLink:         "https://www.youtube.com/embed/JhA6fkF3xF8",
			},
		},
	}
	for key, m := range srv.motorcycles {
		m.HeroImagePath = path.Join("images/hero", key+".jpg")
		m.SideviewImagePath = path.Join("images/sideview", key+".jpg")
		srv.motorcycles[key] = m
	}
	srv.Routes()
	return srv
}

func (srv *server) HandleMotorcycleTemplate(paths ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		m, ok := srv.motorcycles[name]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 page not found")
			return
		}
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(paths...)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err := tpl.Execute(w, m)
		if err != nil {
			log.Println(err)
		}
	}
}

func (srv *server) HandleStaticTemplate(paths ...string) http.HandlerFunc {
	var (
		init   sync.Once
		tpl    *template.Template
		tplerr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplerr = template.ParseFiles(paths...)
		})
		if tplerr != nil {
			http.Error(w, tplerr.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		tpl.Execute(w, nil)
	}
}

func (srv *server) ServeDirectory(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		path := filepath.Join(directory, vars["filename"])
		log.Printf("Serving file: %s\n", path)
		http.ServeFile(w, r, path)
	}
}

func (srv *server) LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s\n", r.RemoteAddr, r.Method, r.URL)
		h(w, r)
	}
}
