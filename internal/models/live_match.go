package models

// Импортируем пакет для работы с JSON

// LiveMatch представляет структуру данных для матча.
type LiveMatch struct {
	V              int     `json:"v"`                // Версия данных
	ID             int64   `json:"id"`               // Уникальный идентификатор матча
	Team1          string  `json:"team1"`            // Название первой команды
	Team1Rus       string  `json:"team1_rus"`        // Название первой команды на русском
	Team1ID        int64   `json:"team1_id"`         // Идентификатор первой команды
	Team2          string  `json:"team2"`            // Название второй команды
	Team2Rus       string  `json:"team2_rus"`        // Название второй команды на русском
	Team2ID        int64   `json:"team2_id"`         // Идентификатор второй команды
	Score1         int     `json:"score1"`           // Счет первой команды
	Score2         int     `json:"score2"`           // Счет второй команды
	Href           string  `json:"href"`             // Ссылка на страницу матча
	DateStart      string  `json:"date_start"`       // Дата и время начала матча
	IsLive         bool    `json:"isLive"`           // Индикатор того, что матч в режиме реального времени
	IsComposite    bool    `json:"isComposite"`      // Индикатор композитного матча
	IsSpecial      bool    `json:"isSpecial"`        // Индикатор специального матча
	Minute         int     `json:"minute"`           // Текущая минута матча
	Seconds        int     `json:"seconds"`          // Текущие секунды матча
	HalfOrderIndex int     `json:"half_order_index"` // Индекс текущей половины матча
	Title          string  `json:"title"`            // Название лиги или турнира
	Country        string  `json:"country"`          // Страна, в которой проходит матч
	League         League  `json:"league"`           // Информация о лиге
	Markets        Markets `json:"markets"`          // Рынки для ставок
	Hash           string  `json:"hash"`             // Хеш данных
	ActualAt       string  `json:"actual_at"`        // Время актуальности данных
}

// League представляет структуру данных для лиги.
type League struct {
	CountryID   int64  `json:"country_id"`  // Идентификатор страны
	LeagueID    int64  `json:"league_id"`   // Идентификатор лиги
	SportID     int64  `json:"sport_id"`    // Идентификатор вида спорта
	Name        string `json:"name"`        // Название лиги
	NameRus     string `json:"name_rus"`    // Название лиги на русском
	IsCyber     bool   `json:"isCyber"`     // Индикатор киберспорта
	IsSimulated bool   `json:"isSimulated"` // Индикатор симуляции
	IsSpecial   bool   `json:"isSpecial"`   // Индикатор специальной лиги
}

// Markets представляет структуру данных для рынков.
type Markets struct {
	Totals       []Total      `json:"totals"`       // Список тоталов
	Totals1      []Total      `json:"totals1"`      // Список тоталов для первой команды
	Totals2      []Total      `json:"totals2"`      // Список тоталов для второй команды
	TotalsAsian  []TotalAsian `json:"totalsAsian"`  // Список азиатских тоталов
	Totals1Asian []TotalAsian `json:"totals1Asian"` // Список азиатских тоталов для первой команды
	Totals2Asian []TotalAsian `json:"totals2Asian"` // Список азиатских тоталов для второй команды
	Handicaps1   []Handicap   `json:"handicaps1"`   // Список форы для первой команды
	Handicaps2   []Handicap   `json:"handicaps2"`   // Список форы для второй команды
	BothToScore  interface{}  `json:"bothToScore"`  // Информация о ставках "обе команды забьют"
	Half         interface{}  `json:"half"`         // Информация о ставках на половину матча
	Win1         Odds         `json:"win1"`         // Коэффициенты на победу первой команды
	WinX         Odds         `json:"winX"`         // Коэффициенты на ничью
	Win2         Odds         `json:"win2"`         // Коэффициенты на победу второй команды
	Win1X        Odds         `json:"win1X"`        // Коэффициенты на победу первой команды или ничью
	Win12        Odds         `json:"win12"`        // Коэффициенты на победу любой из команд
	WinX2        Odds         `json:"winX2"`        // Коэффициенты на ничью или победу второй команды
}

// Total представляет структуру данных для тоталов.
type Total struct {
	Type  float64 `json:"type"`  // Тип тотала
	Over  Odds    `json:"over"`  // Коэффициенты на тотал больше
	Under Odds    `json:"under"` // Коэффициенты на тотал меньше
}

// TotalAsian представляет структуру данных для азиатских тоталов.
type TotalAsian struct {
	Type  float64 `json:"type"`  // Тип азиатского тотала
	Over  Odds    `json:"over"`  // Коэффициенты на тотал больше
	Under Odds    `json:"under"` // Коэффициенты на тотал меньше
}

// Odds представляет структуру данных для коэффициентов.
type Odds struct {
	V float64 `json:"v"` // Значение коэффициента
}

// Handicap представляет структуру данных для форы.
type Handicap struct {
	Type float64 `json:"type"` // Значение форы
	V    float64 `json:"v"`    // Коэффициент на фору
}
