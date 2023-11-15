package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"wb_go_lvl2/dev11/internal/model"
	"wb_go_lvl2/dev11/internal/service"
	"wb_go_lvl2/dev11/internal/transport"
)

type calendarHandler struct {
	serviceCalendar service.Calendar
}

func (h *calendarHandler) Init(router *http.ServeMux) {
	// TODO? --- GET запросы
	router.HandleFunc("/events_for_day", h.eventsForDay)
	router.HandleFunc("/events_for_week", h.eventsForWeek)
	router.HandleFunc("/events_for_month", h.eventsForMonth)

	// TODO? --- POST запросы
	router.HandleFunc("/create_event", h.createEvent)
	router.HandleFunc("/update_event", h.updateEvent)
	router.HandleFunc("/delete_event", h.deleteEvent)
}

// TODO? -- GET
func (h *calendarHandler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodGet {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	date, err := parseAndValidateDate(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для выдачи всех событий за определенный день
	events, err := h.serviceCalendar.EventsForDay(date)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при поиске событий за день: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: events,
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// TODO? -- GET
func (h *calendarHandler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodGet {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	date, err := parseAndValidateDate(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для выдачи всех событий за определенную неделю
	events, err := h.serviceCalendar.EventsForWeek(date)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при поиске событий за неделю: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: events,
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// TODO? -- GET
func (h *calendarHandler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodGet {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	date, err := parseAndValidateDate(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для выдачи всех событий за определенный месяц
	events, err := h.serviceCalendar.EventsForMonth(date)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при поиске событий за месяц: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: events,
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// TODO! -- POST
func (h *calendarHandler) createEvent(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodPost {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	event, err := parseAndValidateParams(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для создания события в календаре
	if err := h.serviceCalendar.Create(event); err != nil {
		responseError := ResponseError{
			Error: "Ошибка при создании события: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: "Success!",
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// TODO! -- POST
func (h *calendarHandler) updateEvent(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodPost {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	event, err := parseAndValidateParams(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для обновлении события в календаре
	if err := h.serviceCalendar.Update(event); err != nil {
		responseError := ResponseError{
			Error: "Ошибка при обновлении события: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: "Success!",
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// TODO! -- POST
func (h *calendarHandler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	// Проверка входящего запроса
	if r.Method != http.MethodPost {
		responseError := ResponseError{
			Error: "Метод не разрешен",
		}
		ResponseJson(w, http.StatusMethodNotAllowed, responseError)
		return
	}

	event, err := parseAndValidateParams(r)
	if err != nil {
		responseError := ResponseError{
			Error: "Ошибка при парсинге данных: " + err.Error(),
		}
		ResponseJson(w, http.StatusBadRequest, responseError)
		return
	}

	// бизнес-логика для удаления события в календаре
	if err := h.serviceCalendar.Delete(event); err != nil {
		responseError := ResponseError{
			Error: "Ошибка при удалении события: " + err.Error(),
		}
		ResponseJson(w, http.StatusServiceUnavailable, responseError)
		return
	}

	// Успешный ответ
	responseOk := ResponseOk{
		Result: "Success!",
	}
	ResponseJson(w, http.StatusOK, responseOk)
}

// parseAndValidateParams распарсивает и валидирует параметры запроса для /create_event и /update_event
func parseAndValidateParams(r *http.Request) (event model.Event, err error) {
	userID := r.FormValue("user_id")
	dateStr := r.FormValue("date")
	title := r.FormValue("title")

	// Получаем путь
	path := r.URL.Path

	if path == "delete_event" {
		// Пример простой валидации, здесь можно добавить дополнительные проверки
		if userID == "" || dateStr == "" {
			return event, fmt.Errorf("неверно переданы параметры запроса")
		}
	} else {
		// Пример простой валидации, здесь можно добавить дополнительные проверки
		if userID == "" || dateStr == "" || title == "" {
			return event, fmt.Errorf("неверно переданы параметры запроса")
		}
	}

	// Преобразование строки в число
	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return event, fmt.Errorf("неверный user_id")
	}

	// Преобразование строки в дату
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return event, fmt.Errorf("неверный формат даты")
	}

	return model.Event{
		UserID: userIDInt,
		Date:   date,
		Title:  title,
	}, nil
}

// parseAndValidateDate распарсивает и валидирует параметры запроса для /events_for_day и /events_for_week и /events_for_month
func parseAndValidateDate(r *http.Request) (time.Time, error) {
	dateStr := r.URL.Query().Get("date")

	// Пример простой валидации, здесь можно добавить дополнительные проверки
	if dateStr == "" {
		return time.Time{}, fmt.Errorf("неверно переданы параметры запроса")
	}

	// Преобразование строки в дату
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("неверный формат даты")
	}

	return date, nil
}

func NewHandlerCalendar(serviceCalendar service.Calendar) transport.Handler {
	return &calendarHandler{
		serviceCalendar: serviceCalendar,
	}
}
