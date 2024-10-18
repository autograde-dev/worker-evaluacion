package api

import (
	"net/http"
	"strconv"
	"time"

	job "github.com/jhonM8a/worker-evaluacion/internal/jobs"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan job.Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid Delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid Value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid NAME", http.StatusBadRequest)
		return
	}

	job := job.Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusOK)
}
