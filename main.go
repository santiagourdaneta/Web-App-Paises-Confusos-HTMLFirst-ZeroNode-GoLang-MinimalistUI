package main

import (
	"html/template"
	"net/http"
	"strings"
	"time"
)

type Pais struct {
	Nombre  string
	Confuso string
	Dato    string
}

// Nuestra BD
var baseDeDatos = []Pais{
	{"Eslovenia", "Eslovaquia", "Tiene Alpes, capital Liubliana y un trozo de mar Adriático."},
	{"Eslovaquia", "Eslovenia", "Capital Bratislava, está en el centro de Europa, sin mar."},
	{"Estonia", "Letonia/Lituania", "Norte de los Bálticos, idioma parecido al finlandés."},
	{"Letonia", "Estonia/Lituania", "En el centro de los Bálticos, capital Riga."},
	{"Lituania", "Letonia", "Sur de los Bálticos, históricamente un gran ducado."},
	{"Suiza", "Suecia", "Montañas, chocolate y bancos (Centro de Europa)."},
	{"Suecia", "Suiza", "IKEA, ABBA y auroras boreales (Norte de Europa)."},
	{"Austria", "Australia", "Música clásica y Alpes. NO tienen canguros."},
	{"Australia", "Austria", "Isla continente con koalas y desiertos."},
	{"Paraguay", "Uruguay", "País mediterráneo (sin salida al mar) en Sudamérica."},
	{"Uruguay", "Paraguay", "Tiene costa al Atlántico y capital Montevideo."},
	{"Guyana", "Guayana Francesa", "Hablan inglés, fue colonia británica."},
	{"Guayana Francesa", "Surinam", "Es un departamento de Francia (usan Euros)."},
	{"Surinam", "Guyana", "Hablan neerlandés (holandés)."},
	{"Guinea", "Guinea-Bissau", "Antigua colonia francesa."},
	{"Guinea-Bissau", "Guinea", "Antigua colonia portuguesa."},
	{"Guinea Ecuatorial", "Guinea", "Único país de África que habla español."},
	{"Congo", "RD Congo", "Capital Brazzaville, es el más pequeño de los dos."},
	{"RD Congo", "Congo", "Capital Kinshasa, país gigante y muy rico en minerales."},
	{"Níger", "Nigeria", "País del desierto, al norte."},
	{"Nigeria", "Níger", "País costero, el más poblado de África."},
	{"Mali", "Malaui", "En el Sahel (África Occidental), capital Bamako."},
	{"Malaui", "Mali", "En el sureste de África, famoso por su gran lago."},
	{"Zambia", "Zimbabue", "Al norte de las Cataratas Victoria."},
	{"Zimbabue", "Zambia", "Al sur de las Cataratas Victoria."},
	{"Tailandia", "Taiwán", "Templos, elefantes y playas en el Sudeste Asiático."},
	{"Taiwán", "Tailandia", "Isla tecnológica cerca de China."},
	{"Omán", "Yemen", "País estable y seguro en la península arábiga."},
	{"Yemen", "Omán", "En el extremo sur, actualmente en conflicto."},
	{"Irak", "Irán", "Hablan árabe, capital Bagdad."},
	{"Irán", "Irak", "Hablan farsi (persa), capital Teherán."},
	{"Mauricio", "Mauritania", "Isla paradisíaca en el Índico."},
	{"Mauritania", "Mauricio", "Gran país del desierto del Sahara."},
	{"Mónaco", "Marruecos", "Microestado de lujo en la Costa Azul."},
	{"Marruecos", "Mónaco", "Reino del norte de África, famoso por sus zocos."},
	{"Dominica", "Rep. Dominicana", "Isla pequeña y volcánica (Antillas Menores)."},
	{"Rep. Dominicana", "Dominica", "Comparte isla con Haití, famosa por Punta Cana."},
	{"Senegal", "Gambia", "Rodea casi por completo a Gambia."},
	{"Gambia", "Senegal", "Un país pequeño que sigue la forma de un río."},
	{"Corea del Norte", "Corea del Sur", "Régimen cerrado al norte del paralelo 38."},
	{"Corea del Sur", "Corea del Norte", "Gigante tecnológico y K-Pop al sur."},
	{"Honduras", "El Salvador", "Tiene costa en ambos océanos."},
	{"El Salvador", "Honduras", "El más pequeño de Centroamérica, solo costa al Pacífico."},
	{"Nicaragua", "Panamá", "Tierra de lagos y volcanes."},
	{"Panamá", "Nicaragua", "Famoso por su canal que une océanos."},
	{"Jordania", "Israel", "Famoso por Petra y el Mar Muerto."},
	{"Bután", "Nepal", "El país de la felicidad, entre India y China."},
	{"Nepal", "Bután", "Donde está el Everest."},
	{"Libia", "Líbano", "En el norte de África."},
	{"Líbano", "Libia", "En el Mediterráneo oriental (Medio Oriente)."},
}

// El mapa de visitantes DEBE estar fuera de la función para que no se borre
var visitantes = make(map[string]time.Time)

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://unpkg.com; style-src 'self' 'unsafe-inline'; img-src 'self' data:;")
		w.Header().Set("X-Content-Type-Options", "nosniff")

		// Rate Limiting
		ip := r.RemoteAddr
		if lastVisit, ok := visitantes[ip]; ok {
			if time.Since(lastVisit) < 500*time.Millisecond {
				http.Error(w, "Demasiado rápido, respira...", http.StatusTooManyRequests)
				return
			}
		}
		visitantes[ip] = time.Now()

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	// 1. Servir archivos estáticos
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// 2. Cargar template
	tmpl := template.Must(template.ParseFiles("index.html"))

	// 3. Ruta principal única
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Solo procesar la raíz
		if r.URL.Path != "/" {
			return
		}

		query := strings.ToLower(r.URL.Query().Get("q"))

		// Escudo contra "Ruido" y ataques de saturación
		if len(query) > 40 {
			w.Write([]byte("<p style='color: var(--accent);'>⚠️ ENTRADA DEMASIADO LARGA</p>"))
			return
		}

		// Opcional: Si detecta caracteres sospechosos como + * / ?
		if strings.ContainsAny(query, "+*/?$%#") {
			w.Write([]byte("<p style='color: var(--accent); font-family: serif;'>🚫 BÚSQUEDA NO VÁLIDA: SOLO LETRAS</p>"))
			return
		}

		var resultados []Pais
		for _, p := range baseDeDatos {
			if strings.Contains(strings.ToLower(p.Nombre), query) {
				resultados = append(resultados, p)
			}
		}

		if r.Header.Get("HX-Request") == "true" {
			tmpl.ExecuteTemplate(w, "resultados", resultados)
			return
		}

		tmpl.Execute(w, resultados)
	})

	// 4. Configurar y encender el servidor UNA SOLA VEZ
	server := &http.Server{
		Addr:         ":8080",
		Handler:      secureHeaders(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	println("Servidor corriendo en http://localhost:8080")
	server.ListenAndServe()
}
