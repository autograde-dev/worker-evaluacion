package api

import (
	"net/http"
	"strconv"

	job "github.com/jhonM8a/worker-evaluacion/internal/job"
)

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan job.Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	nameFileAnswer := r.FormValue("nameFileAnswer")
	if nameFileAnswer == "" {
		http.Error(w, "Invalid nameFileAnswer", http.StatusBadRequest)
		return
	}

	nameFileEvaluation := r.FormValue("nameFileEvaluation")
	if nameFileEvaluation == "" {
		http.Error(w, "Invalid nameFileEvaluation", http.StatusBadRequest)
		return
	}

	nameBucket := r.FormValue("nameBucket")
	if nameFileEvaluation == "" {
		http.Error(w, "Invalid nameBucket", http.StatusBadRequest)
		return
	}

	idEValuation, err := strconv.Atoi(r.FormValue("idEValuation"))
	if err != nil {
		http.Error(w, "Invalid idEValuation", http.StatusBadRequest)
		return
	}

	job := job.Job{NameFileEvaluation: nameFileEvaluation, NameFileAnswer: nameFileAnswer, NameBucket: nameBucket, IDEValuation: idEValuation}

	jobQueue <- job
	w.WriteHeader(http.StatusOK)
}
