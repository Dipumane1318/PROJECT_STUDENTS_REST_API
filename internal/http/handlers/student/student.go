package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Dipumane1318/PROJECT_STUDENTS_REST_API/internal/types"
	"github.com/Dipumane1318/PROJECT_STUDENTS_REST_API/internal/utils/response"
	"github.com/go-playground/validator/v10"
	// "golang.org/x/mod/sumdb/storage"
	"github.com/Dipumane1318/PROJECT_STUDENTS_REST_API/internal/storage"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request)  {
		slog.Info("creating a student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation 
		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		lastId, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil{
			response.WriteJSON(w, http.StatusInternalServerError, err)
			return 
		}

		response.WriteJSON(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}  